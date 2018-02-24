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
	"fmt"
	"github.com/sanitizer/cloud_sql_dao/dao/model"
	//"go/doc"
	"time"
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

func GetDataFromFireBase()  {
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

		data := document.Data()

		var bn = ""
		var dis = ""
		var disp = ""
		var em = ""
		var grav = ""
		var tl = ""
		var la time.Time
		var li = ""

		if data["battleNet"] != nil {
			bn = data["battleNet"].(string)
		}

		if data["discord"] != nil {
			dis = data["discord"].(string)
		}

		if  data["displayName"] != nil {
			disp = data["displayName"].(string)
		}

		if data["email"] != nil {
			em = data["email"].(string)
		}

		if  data["gravatar"] != nil {
			grav =  data["gravatar"].(string)
		}

		if data["teamLogo"] != nil {
			tl = data["teamLogo"].(string)
		}

		if data["lastActive"] != nil {
			la = time.Unix(data["lastActive"].(int64), int64(0))
		}

		if data["id"] != nil {
			li = data["id"].(string)
		}

		user:= model.User{BattleNetId: bn,
				  DiscordId: dis,
				  DisplayName: disp,
				  Email: em,
				  GrAvatar: grav,
				  TeamLogo: tl,
				  LastActive: la,
				  LegacyId: li}

		fmt.Println(user)
		users = append(users, user)
	}


	fmt.Println(len(users))
	fmt.Println(count)

	timeStamp := time.Unix(int64(1519444638), int64(0))
	fmt.Println(timeStamp.Unix())

}

// INSERT INTO table_tags (tag) VALUES ('tag_a'),('tab_b'),('tag_c') ON DUPLICATE KEY UPDATE tag=tag; UPSERT QUERY EXAMPLE