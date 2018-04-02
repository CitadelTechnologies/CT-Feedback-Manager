package comment

import(
	"ct-feedback-manager/author"
	"ct-feedback-manager/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

func GetFeedbackComments(feedbackId string) Comments {
  	comments := make(Comments, 0)

    pipe := mongo.Database.C("feedbacks").Pipe(
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

func CreateComment(feedbackId string, content string, authorData map[string]interface{}) (*Comment, error) {
	commentAuthor, err := author.CreateAuthor(authorData["name"].(string), authorData["email"].(string))
	if err != nil {
		return nil, err
	}
	comment := &Comment{
		Id: bson.NewObjectId(),
		Content: content,
		Author: commentAuthor,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	change := bson.M{
		"$set": bson.M{"updatedat": time.Now()},
		"$push": bson.M{"comments": comment},
	}
	if err := mongo.Database.C("feedbacks").Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
		return nil, err
	}
	return comment, nil
}

func UpdateComment(feedbackId string, commentId string, data map[string]interface{}) *Comment {
	var comment Comment

	change := mgo.Change {
		Update:    bson.M{"$set": bson.M{
			"content": data["name"].(string),
			"updatedat": time.Now(),
		}},
		Upsert:    false,
		Remove:    false,
		ReturnNew: true,
	}

	_, err := mongo.Database.
		C("feedbacks").
		FindId(feedbackId).
		Select(bson.M{"comments": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex(commentId)}}}).
		Apply(change, &comment)
	if err != nil {
		panic(err);
	}
	return &comment
}

func DeleteComment(feedbackId string, commentId string) error {
  	change := mgo.Change {
  		Update:    false,
		Upsert:    false,
		Remove:    true,
		ReturnNew: false,
  	}
    _, err := mongo.Database.
      C("feedbacks").
      FindId(feedbackId).
      Select(bson.M{"comments": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex(commentId)}}}).
      Apply(change, nil)
	if err != nil {
		return err
	}
	return nil
}
