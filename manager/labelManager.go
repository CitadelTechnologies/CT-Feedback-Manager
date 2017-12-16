package manager

import(
	"ct-feedback-manager/model"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

func GetLabels() model.Labels {
  	labels := make(model.Labels, 0)
  	if err := MongoDBConnection.DB(MongoDBName).C("labels").Find(nil).All(&labels); err != nil {
      panic(err)
    }
    return labels
}

func CreateLabel(name string, color string) model.Label {
  label := model.Label{
    Id: bson.NewObjectId(),
    Name: name,
    Color: color,
  }
  if err := MongoDBConnection.DB(MongoDBName).C("labels").Insert(label); err != nil {
    panic(err)
  }
  return label
}

func UpdateLabel(id string, data map[string]string) *model.Label {
	var label model.Label

	change := mgo.Change {
		Update:    bson.M{"$inc": bson.M{"n": 1}, "$set": bson.M{
			"name": data["name"],
			"color": data["color"],
		}},
    Upsert:    false,
    Remove:    false,
    ReturnNew: true,
	}

	_, err := MongoDBConnection.DB(MongoDBName).C("labels").FindId(bson.ObjectIdHex(id)).Apply(change, &label)
	if err != nil {
		panic(err);
	}
	return &label
}


func DeleteLabel(id string) bool {
		err := MongoDBConnection.DB(MongoDBName).C("labels").RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			panic(err);
		}
		return true
}
