package main

import (
	"github.com/sanitizer/cloud_sql_dao/dao"
	"github.com/sanitizer/cloud_sql_dao/dao/model"
	"log"
	"strconv"
	"strings"
)

func main() {
	rawData, err := dao.GetDataFromFireBase("matches")

	if err != nil {
		log.Println(err.Error())
	}

	//log.Println(rawData)
	dt := make(map[string]bool)

	for k, v := range rawData {
		var away, home string
		if v["awayScore"] == nil {
			away = "0"
		} else {
			away = v["awayScore"].(string)
		}

		if v["homeScore"] == nil {
			home = "0"
		} else {
			if _, ok := v["homeScore"].(int64); ok {
				home = strconv.Itoa(int(v["homeScore"].(int64)))
			} else {
				home = v["homeScore"].(string)
			}
		}
		dt[k + ":" + home + ":" + away] = true
		//v["id"] = k
		//dt = append(dt, v)
	}

	query := "INSERT INTO match_stat (matchId, homeScore, awayScore) VALUES "
	anotherCounter := 0
	for k, _ := range dt {
		splitted := strings.Split(k, ":")
		if anotherCounter == 0 {
			query = query + "((select id from game_match where legacyId = \"" + splitted[0] + "\"), " + splitted[1] + ",  " + splitted[2] + ")"
		} else {
			query = query + "," + "((select id from game_match where legacyId = \"" + splitted[0] + "\"), " + splitted[1] + ",  " + splitted[2] + ")"
		}
		anotherCounter = anotherCounter + 1
	}
	//log.Println(query)

	//log.Println(dt)
	//dao.RunMigration()
	//
	//query := BuildFromRawData(dt, model.LeagueUser{})
	query = query + " on duplicate key update homeScore=homeScore"

	inserted, err := RunInsertQuery(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("inserted: " + strconv.Itoa(int(inserted)))
}

func BuildFromRawData(rawData []map[string]interface{}, entity model.DaoModel) string {
	data := make([]model.DaoModel, 0)
	for _, v := range rawData {
		datum := entity.BuildFromFirestoreData(v)
		data = append(data, datum)
	}

	var query = entity.GetPartialInsertQuery()

	anotherCounter := 0
	for _, datum := range data {
		if anotherCounter == 0 {
			query = query + datum.StringForInsert()
		} else {
			query = query + "," + datum.StringForInsert()
		}
		anotherCounter = anotherCounter + 1
	}

	return query
}

func RunInsertQuery(query string) (int64, error) {
	var err error
	con := new(dao.CloudConnection)
	db, err := con.GetNewConnection()
	if err != nil {
		return int64(0), err
	}
	defer db.Close()

	transaction, err := db.Begin()
	if err != nil {
		return int64(0), err
	}
	defer transaction.Rollback()

	stmt, err := db.Prepare(query)
	if err != nil {
		return int64(0), err
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		return int64(0), err
	}

	inserted, err := result.RowsAffected()
	if err != nil {
		return int64(0), err
	}
	return inserted, nil
}
