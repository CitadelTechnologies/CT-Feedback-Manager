package manager

import(
	"ct-feedback-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

func GetFeedbackComments(feedbackId string, collectionName string) model.Comments {
  	comments := make(model.Comments, 0)

    pipe := MongoDBConnection.DB(MongoDBName).C(collectionName).Pipe(
        []bson.M{
						bson.M{
								"$match": bson.M{"_id": bson.ObjectIdHex(feedbackId)},
						},
            bson.M{
                "$unwind": "$comments",
            },
            bson.M{
                "$project": bson.M{
                    "content": "$comments.content",
                    "author": "$comments.author",
                    "createdat": "$comments.createdat",
                    "updatedat": "$comments.updatedat",
                },
            },
        },
    )
    if err := pipe.All(&comments); err != nil {
			panic(err)
		}
    return comments
}

func CreateComment(feedbackId string, content string, author map[string]interface{}, collectionName string) model.Comment {
  comment := model.Comment{
    Id: bson.NewObjectId(),
    Content: content,
    Author: CreateAuthor(author["name"].(string), author["email"].(string)),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  change := bson.M{"$push": bson.M{"comments": comment}}
  if err := MongoDBConnection.DB(MongoDBName).C(collectionName).Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
    panic(err)
  }
  return comment
}

func UpdateComment(feedbackId string, commentId string, data map[string]interface{}, collectionName string) *model.Comment {
	var comment model.Comment

	change := mgo.Change {
		Update:    bson.M{"$inc": bson.M{"n": 1}, "$set": bson.M{
			"content": data["name"].(string),
			"updatedat": time.Now(),
		}},
    Upsert:    false,
    Remove:    false,
    ReturnNew: true,
	}

  _, err := MongoDBConnection.
    DB(MongoDBName).
    C(collectionName).
    FindId(feedbackId).
    Select(bson.M{"comments": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex(commentId)}}}).
    Apply(change, &comment)

	if err != nil {
		panic(err);
	}
	return &comment
}

func DeleteComment(feedbackId string, commentId string, collectionName string) bool {
  	change := mgo.Change {
  		Update:    false,
      Upsert:    false,
      Remove:    true,
      ReturnNew: false,
  	}
    _, err := MongoDBConnection.
      DB(MongoDBName).
      C(collectionName).
      FindId(feedbackId).
      Select(bson.M{"comments": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex(commentId)}}}).
      Apply(change, nil)
		if err != nil {
			panic(err);
		}
		return true
}
