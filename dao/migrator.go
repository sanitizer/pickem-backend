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

func GetDataFromFireBase(entity string) (map[string]map[string]interface{}, error) {
	ctx := context.Background()
	options := option.WithCredentialsFile("dao/config/firebase.json")
	app, err := firebase.NewApp(ctx, nil, options)

	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	iter := client.Collection(entity).Documents(ctx)
	result := make(map[string]map[string]interface{}, 0)

	for {
		document, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		result[document.Ref.ID] = document.Data()
	}

	return result, nil
}

// INSERT INTO table_tags (tag) VALUES ('tag_a'),('tab_b'),('tag_c') ON DUPLICATE KEY UPDATE tag=tag; UPSERT QUERY EXAMPLE