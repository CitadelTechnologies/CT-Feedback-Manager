package manager

import(
	"github.com/gosimple/slug"
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
		for key, bug := range bugs {
			bugs[key].Labels = make(model.Labels, 0)
			for _, labelId := range bug.LabelIds {
				bugs[key].Labels = append(bugs[key].Labels, GetLabel(labelId.Hex()))
			}
		}
    return bugs
}

func GetBug(id string) *model.Bug {
  	var bug model.Bug
		var identifier bson.M

		if bson.IsObjectIdHex(id) {
			identifier = bson.M{"_id": bson.ObjectIdHex(id)}
		} else {
			identifier = bson.M{"slug": id}
		}
  	if err := MongoDBConnection.DB(MongoDBName).C("bugs").Find(identifier).One(&bug); err != nil {
      if err.Error() == "not found" {
				return nil
			}
			panic(err)
    }
		for _, labelId := range bug.LabelIds {
			bug.Labels = append(bug.Labels, GetLabel(labelId.Hex()))
		}
    return &bug
}

func CreateBug(title string, description string, status string, author map[string]string) model.Bug {
	bug := model.Bug{
		Id: bson.NewObjectId(),
		Feedback: model.Feedback{
		  Title: title,
			Slug: slug.Make(title),
		  Description: description,
			Status: status,
			Author: CreateAuthor(author["name"], author["email"]),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
  if err := MongoDBConnection.DB(MongoDBName).C("bugs").Insert(bug); err != nil {
    panic(err)
  }
  return bug
}

func UpdateBug(id string, data map[string]string) *model.Bug {
	var bug model.Bug

	change := mgo.Change {
		Update:    bson.M{"$inc": bson.M{"n": 1}, "$set": bson.M{
			"title": data["title"],
			"slug": slug.Make(data["title"]),
			"description": data["description"],
			"status": data["status"],
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

func AddLabelToBug(feedbackId string, label *model.Label) *model.Bug {
	bug := GetBug(feedbackId)
	if bug == nil {
		return nil
	}
	bug.Labels = append(bug.Labels, label)
	bug.LabelIds = append(bug.LabelIds, label.Id)
  change := bson.M{
		"$push": bson.M{"labels": label.Id},
		"$set": bson.M{"updatedat": time.Now()},
	}
  if err := MongoDBConnection.DB(MongoDBName).C("bugs").Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
    panic(err)
  }
  return bug
}

func RemoveLabelFromBug(feedbackId string, label *model.Label) *model.Bug {
	bug := GetBug(feedbackId)
	if bug == nil {
		return nil
	}
  change := bson.M{
		"$pull": bson.M{"labels": label.Id},
		"$set": bson.M{"updatedat": time.Now()},
	}
  if err := MongoDBConnection.DB(MongoDBName).C("bugs").Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
    panic(err)
  }
  return bug
}

func DeleteBug(id string) bool {
		err := MongoDBConnection.DB(MongoDBName).C("bugs").RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			panic(err);
		}
		return true
}
