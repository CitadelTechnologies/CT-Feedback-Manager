package manager

import(
	"gopkg.in/mgo.v2"
	"os"
)

var MongoDBConnection *mgo.Session
var MongoDBName string

func InitMongo() {
	var err error

	MongoDBConnection, err = mgo.Dial("mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	MongoDBName = os.Getenv("MONGO_DBNAME")

	if err != nil {
		panic(err)
	}
}
