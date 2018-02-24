package main

import (
	"database/sql"
	"fmt"
	//"log"

	"github.com/sanitizer/cloud_sql_dao/dao"
	//"log"
	//"strconv"
	"log"
	"strconv"
)

var db *sql.DB

func main() {
	query := dao.GetDataFromFireBase()

	//fmt.Println(query)
	//dao.RunMigration()

	var err error
	con := new(dao.CloudConnection)
	db, err = con.GetNewConnection()
	if err != nil {
		log.Fatal("Error connection ", err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Error starting transaction ", err.Error())
	}
	defer tx.Rollback()
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("Error preparing statement " + err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec()

	if err != nil {
		log.Fatal("Error inserting rows " + err.Error())
	}

	inserted, err := result.RowsAffected()

	if err != nil {
		log.Fatal("Error getting inserted rows count " + err.Error())
	}
	fmt.Println("inserted " + strconv.Itoa(int(inserted)))
	//handler()
}

func handler() {
	rows, err := db.Query("select version from schema_migrations")
	if err != nil {
		fmt.Println("error quering db " + err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			fmt.Println("error scanning row " + err.Error())
			return
		}
		fmt.Println(version)
	}
}


