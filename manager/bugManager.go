package manager

import(
	"ct-feedback-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

func GetBugs() model.Bugs {
  	bugs := make(model.Bugs, 0)
  	if err := MongoDBConnection.DB(MongoDBName).C("bugs").Find(nil).All(&bugs); err != nil {
      panic(err)
    }
    return bugs
}

func GetBug(id string) *model.Bug {
  	var bug model.Bug

  	if !bson.IsObjectIdHex(id) {
  		return nil
  	}
  	if err := MongoDBConnection.DB(MongoDBName).C("bugs").FindId(bson.ObjectIdHex(id)).One(&bug); err != nil {
      if err.Error() == "not found" {
				return nil
			}
			panic(err)
    }
    return &bug
}

func CreateBug(title string, description string, status string, author map[string]interface{}) model.Bug {
	bug := model.Bug{
		Id: bson.NewObjectId(),
		Feedback: model.Feedback{
		  Title: title,
		  Description: description,
			Status: status,
			Author: CreateAuthor(author["name"].(string), author["email"].(string)),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
  if err := MongoDBConnection.DB(MongoDBName).C("bugs").Insert(bug); err != nil {
    panic(err)
  }
  return bug
}

func UpdateBug(id string, data map[string]interface{}) *model.Bug {
	var bug model.Bug

	change := mgo.Change {
		Update:    bson.M{"$inc": bson.M{"n": 1}, "$set": bson.M{
			"title": data["title"].(string),
			"description": data["description"].(string),
			"status": data["status"].(string),
			"updatedat": time.Now(),
		}},
    Upsert:    false,
    Remove:    false,
    ReturnNew: true,
	}

	_, err := MongoDBConnection.DB(MongoDBName).C("bugs").FindId(bson.ObjectIdHex(id)).Apply(change, &bug)
	if err != nil {
		panic(err);
	}
	return &bug
}

func DeleteBug(id string) bool {
		err := MongoDBConnection.DB(MongoDBName).C("bugs").RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			panic(err);
		}
		return true
}
