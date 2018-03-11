package main

import (
	"github.com/sanitizer/cloud_sql_dao/dao"
	"github.com/sanitizer/cloud_sql_dao/dao/model"
	"log"
	"strconv"
	"strings"
)

func main() {
	rawData, err := dao.GetDataFromFireBase("leaderboardSimpleStage2")

	if err != nil {
		log.Println(err.Error())
	}

	//log.Println(rawData)
	dt := make(map[string]bool)

	for k, v := range rawData {
		if v["leagueId"] != nil {
			lId := v["leagueId"].(string)
			p := v["points"].(int64)
			dt[lId + ":" + k + ":" + strconv.Itoa(int(p))] = true
		}
		//v["id"] = k
		//dt = append(dt, v)
	}

	query := "INSERT ignore INTO leaderboard_leagueuser_points (leaderboardId, leagueUserId, points) VALUES "
	anotherCounter := 0
	for k, _ := range dt {
		splitted := strings.Split(k, ":")
		if anotherCounter == 0 {
			query = query +"(ifnull((select id from leaderboard where type = 'SIMPLE' and stage = 2 and leagueId = (select id from league where legacyId = \"" + splitted[0] + "\")), -1), ifnull((select id from league_user where userId = (select id from app_user where legacyId = \"" + splitted[1] + "\") and leagueId = (select id from league where legacyId = \"" + splitted[0] + "\")), -1), " + splitted[2] + ")"
		} else {
			query = query + "," + "(ifnull((select id from leaderboard where type = 'SIMPLE' and stage = 2 and leagueId = (select id from league where legacyId = \"" + splitted[0] + "\")), -1), ifnull((select id from league_user where userId = (select id from app_user where legacyId = \"" + splitted[1] + "\") and leagueId = (select id from league where legacyId = \"" + splitted[0] + "\")), -1), " + splitted[2] + ")"
		}
		anotherCounter = anotherCounter + 1
	}
	//log.Println(query)

	//log.Println(dt)
	//dao.RunMigration()
	//
	//query := BuildFromRawData(dt, model.LeagueUser{})
	query = query + " on duplicate key update points=points"

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
