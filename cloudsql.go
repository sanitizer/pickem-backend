package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/sanitizer/cloud_sql_dao/dao"
)

var db *sql.DB

func main() {
	dao.RunMigration()

	var err error
	con := new(dao.CloudConnection)
	db, err = con.GetNewConnection()
	if err != nil {
		log.Fatal("Error connection ", err.Error())
	}
	handler()
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


