package examples

import (
	"github.com/sayluck/mongo-example/pkgs/mongodb"
	"log"
)

func Connect() *mongodb.MongoDB {
	// new db client
	mgoer := mongodb.Create("mongodb://root:example@localhost:27017", 10, "dbTest", "userTest")
	mgo := mgoer.NewDBClient()
	db := mgo.PingTest()
	if db.Error != nil {
		log.Fatal(db.Error)
	}
	log.Println("mongo db connect succ,")
	return db
}
