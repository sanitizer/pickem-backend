package dao

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

func RunMigration() {
	log.Println("Running db migration...")
	con := new(CloudConnection)

	db, err := con.GetNewConnection()
	if err != nil {
		log.Println("Migration stopped. Error while creating connection: ", err.Error())
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Migration stopped. Error while getting db driver: ", err.Error())
		return
	}

	migration, err := migrate.NewWithDatabaseInstance("file://dao/migrations", "mysql", driver)
	if err != nil {
		log.Println("Migration stopped. Error migrate ", err.Error())
		return
	}

	err = migration.Steps(2)
	if err != nil {
		log.Println("Migration stopped. Error while running migration steps: ", err.Error())
	}
}