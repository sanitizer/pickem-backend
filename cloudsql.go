package main

import (
	"github.com/sanitizer/cloud_sql_dao/dao"
	"github.com/sanitizer/cloud_sql_dao/dao/model"
	"log"
	"strconv"
	"strings"
)

func main() {
	rawData, err := dao.GetDataFromFireBase("teams")

	if err != nil {
		log.Println(err.Error())
	}

	//log.Println(rawData)
	dt := make(map[string]bool)

	for k, v := range rawData {
		if v["compId"] != nil {
			lId := v["compId"].(string)
			dt[lId + ":" + k] = true
		}
		//v["id"] = k
		//dt = append(dt, v)
	}

	query := "INSERT ignore INTO team_competition (teamId, competitionId) VALUES "
	anotherCounter := 0
	for k, _ := range dt {
		splitted := strings.Split(k, ":")
		if anotherCounter == 0 {
			query = query + "((select id from team where legacyId = \"" + splitted[1] + "\"), (select id from competition where legacyId = \"" + splitted[0] + "\"))"
		} else {
			query = query + "," + "((select id from team where legacyId = \"" + splitted[1] + "\"), (select id from competition where legacyId = \"" + splitted[0] + "\"))"
		}
		anotherCounter = anotherCounter + 1
	}
	//log.Println(query)

	//log.Println(dt)
	//dao.RunMigration()
	//
	//query := BuildFromRawData(dt, model.LeagueUser{})
	query = query + " on duplicate key update teamId=teamId"

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
