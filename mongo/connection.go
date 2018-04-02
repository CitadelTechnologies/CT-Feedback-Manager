package mongo

import(
	"gopkg.in/mgo.v2"
	"os"
)

var Connection *mgo.Session
var Database *mgo.Database

func init() {
	var err error

	Connection, err = mgo.Dial("mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	Database = Connection.DB(os.Getenv("MONGO_DBNAME"))

	if err != nil {
		panic(err)
	}
}
