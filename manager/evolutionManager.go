package manager

import(
	"ct-feedback-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

func GetEvolutions() model.Evolutions {
  	evolutions := make(model.Evolutions, 0)
  	if err := MongoDBConnection.DB(MongoDBName).C("evolutions").Find(nil).All(&evolutions); err != nil {
      panic(err)
    }
		for key, evolution := range evolutions {
			evolutions[key].Labels = make(model.Labels, 0)
			for _, labelId := range evolution.LabelIds {
				evolutions[key].Labels = append(evolutions[key].Labels, GetLabel(labelId.Hex()))
			}
		}
    return evolutions
}

func GetEvolution(id string) *model.Evolution {
  	var evolution model.Evolution

  	if !bson.IsObjectIdHex(id) {
  		return nil
  	}
  	if err := MongoDBConnection.DB(MongoDBName).C("evolutions").FindId(bson.ObjectIdHex(id)).One(&evolution); err != nil {
      panic(err)
    }
		for _, labelId := range evolution.LabelIds {
			evolution.Labels = append(evolution.Labels, GetLabel(labelId.Hex()))
		}
    return &evolution
}

func CreateEvolution(title string, description string, status string, author map[string]interface{}) model.Evolution {
	evolution := model.Evolution{
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
  if err := MongoDBConnection.DB(MongoDBName).C("evolutions").Insert(evolution); err != nil {
    panic(err)
  }
  return evolution
}

func UpdateEvolution(id string, data map[string]interface{}) *model.Evolution {
	var evolution model.Evolution

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

	_, err := MongoDBConnection.DB(MongoDBName).C("evolutions").FindId(bson.ObjectIdHex(id)).Apply(change, &evolution)
	if err != nil {
		panic(err);
	}
	return &evolution
}

func AddLabelToEvolution(feedbackId string, label *model.Label) *model.Evolution {
	evolution := GetEvolution(feedbackId)
	if evolution == nil {
		return nil
	}
	evolution.Labels = append(evolution.Labels, label)
	evolution.LabelIds = append(evolution.LabelIds, label.Id)
  change := bson.M{
		"$push": bson.M{"labels": label.Id},
		"$set": bson.M{"updatedat": time.Now()},
	}
  if err := MongoDBConnection.DB(MongoDBName).C("evolutions").Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
    panic(err)
  }
  return evolution
}

func RemoveLabelFromEvolution(feedbackId string, label *model.Label) *model.Evolution {
	evolution := GetEvolution(feedbackId)
	if evolution == nil {
		return nil
	}
  change := bson.M{
		"$pull": bson.M{"labels": label.Id},
		"$set": bson.M{"updatedat": time.Now()},
	}
  if err := MongoDBConnection.DB(MongoDBName).C("evolutions").Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
    panic(err)
  }
  return evolution
}

func DeleteEvolution(id string) bool {
		err := MongoDBConnection.DB(MongoDBName).C("evolutions").RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			panic(err);
		}
		return true
}
