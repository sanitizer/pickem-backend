package main

import (
	"github.com/sanitizer/cloud_sql_dao/dao"
	"log"
	"github.com/sanitizer/cloud_sql_dao/dao/model"
	//"strconv"
)

func main() {
	rawData, err := dao.GetDataFromFireBase("leagues")

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(rawData)
	//dt := make([]map[string]interface{}, 0)

	//for k, v := range rawData {
	//	temp := v["data"].(map[string]interface{})
	//	temp["id"] = k
	//	dt = append(dt, temp)
	//}

	//log.Println(dt)
	//dao.RunMigration()

	//query := BuildFromRawData(dt, model.League{})

	//query = query + " on duplicate key update name = name, fullName = fullName"

	//inserted, err := RunInsertQuery(query)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println("inserted: " + strconv.Itoa(int(inserted)))
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
