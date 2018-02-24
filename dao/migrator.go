package dao

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	//"fmt"
	"github.com/sanitizer/cloud_sql_dao/dao/model"
	//"go/doc"
	"fmt"
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

func GetDataFromFireBase() string {
	ctx := context.Background()
	opt := option.WithCredentialsFile("dao/config/firebase.json")
	app, err := firebase.NewApp(ctx, nil, opt)

	users := make([]model.User, 0)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	iter := client.Collection("users").Documents(ctx)
	count := 0
	for {
		document, err := iter.Next()

		if err == iterator.Done {
			break
		}
		count = count +1

		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		user := new(model.User).Build(document.Data())
		users = append(users, user)
		//fmt.Println(user.DisplayName)
	}
	fmt.Println(count)

	var query = "INSERT INTO app_user " + model.USER_INSERT_COLS + " VALUES"

	anotherCounter := 0
	for _, user := range users {
		if anotherCounter == 0 {
			query = query + user.StringForInsert()
		} else {
			query = query + "," + user.StringForInsert()
		}
		anotherCounter = anotherCounter + 1
	}

	//fmt.Println(query)
	return query

	//fmt.Println(len(sers))
	//fmt.Println(count)

}

// INSERT INTO table_tags (tag) VALUES ('tag_a'),('tab_b'),('tag_c') ON DUPLICATE KEY UPDATE tag=tag; UPSERT QUERY EXAMPLE