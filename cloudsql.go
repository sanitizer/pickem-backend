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

	log.Println(rawData)
	dt := make(map[string]bool)

	for _, v := range rawData {
		log.Println(v["id"])
		if v["roster"] != nil {
			for _, val  := range v["roster"].([]interface{}) {
				playerName := val.(map[string]interface{})["player"].(string)
				role := val.(map[string]interface{})["role"].(string)
				dt[v["id"].(string) + ":" + strings.ToLower(playerName) + ":" + strings.ToLower(role)] = true
			}

		}
		//dt[v["id"] + ":x"] = true
		//v["id"] = k
		//dt = append(dt, v)
	}

	teams := make(map[string]string)

	for k, v := range rawData {
		teams[k] = strings.ToLower(v["name"].(string))
	}

	log.Println(teams)
	log.Println(len(teams))

	query := "INSERT INTO team_player_role (teamPlayerId, role) VALUES "
	anotherCounter := 0
	for k, _ := range dt {
		splitted := strings.Split(k, ":")
		if anotherCounter == 0 {
			query = query + "((select id from team_player where teamId = (select id from team where LOWER(name) = '" + teams[splitted[0]] + "') and playerId = (select id from player where LOWER(name) = '" + splitted[1] + "')), '" + splitted[2] + "')"
		} else {
			query = query + "," + "((select id from team_player where teamId = (select id from team where LOWER(name) = '" + teams[splitted[0]] + "') and playerId = (select id from player where LOWER(name) = '" + splitted[1] + "')), '" + splitted[2] + "')"
		}
		anotherCounter = anotherCounter + 1
	}
	//log.Println(query)

	//log.Println(dt)
	//dao.RunMigration()
	//
	//query := BuildFromRawData(dt, model.LeagueUser{})
	query = query + " on duplicate key update role=role"
	//
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
