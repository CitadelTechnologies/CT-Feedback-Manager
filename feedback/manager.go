package feedback

import(
	"ct-feedback-manager/author"
	"ct-feedback-manager/label"
	"ct-feedback-manager/mongo"
	"github.com/gosimple/slug"
	"errors"
	"time"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

func GetFeedbacks() []Feedback {
  	feedbacks := make([]Feedback, 0)
  	if err := mongo.Database.C("feedbacks").Find(nil).Sort("-updatedat").All(&feedbacks); err != nil {
		panic(err)
    }
	for key, feedback := range feedbacks {
		feedbacks[key].Labels = make(label.Labels, 0)
		for _, labelId := range feedback.LabelIds {
			feedbacks[key].Labels = append(feedbacks[key].Labels, label.GetLabel(labelId.Hex()))
		}
	}
    return feedbacks
}

func GetFeedback(id string) *Feedback {
  	var feedback Feedback
	var identifier bson.M

	if bson.IsObjectIdHex(id) {
		identifier = bson.M{"_id": bson.ObjectIdHex(id)}
	} else {
		identifier = bson.M{"slug": id}
	}
  	if err := mongo.Database.C("feedbacks").Find(identifier).One(&feedback); err != nil {
		return nil
    }
	for _, labelId := range feedback.LabelIds {
		feedback.Labels = append(feedback.Labels, label.GetLabel(labelId.Hex()))
	}
    return &feedback
}

func CreateFeedback(title string, feedbackType string,  description string, status string, authorData map[string]interface{}) (*Feedback, error) {
	if feedbackType != FEEDBACK_TYPE_BUG && feedbackType != FEEDBACK_TYPE_EVOLUTION {
		return nil, errors.New("Invalid feedback type. Must be 'bug' or 'evolution'")
	}
	feedbackAuthor, err := author.CreateAuthor(authorData["name"].(string), authorData["email"].(string))
	if err != nil {
		return nil, err
	}
	feedback := &Feedback{
		Id: bson.NewObjectId(),
		Type: feedbackType,
	  	Title: title,
		Slug: slug.Make(title),
	  	Description: description,
		Status: status,
		Author: feedbackAuthor,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := mongo.Database.C("feedbacks").Insert(feedback); err != nil {
		return nil, err
	}
	return feedback, nil
}

func UpdateFeedback(id string, data map[string]string) *Feedback {
	var feedback Feedback

	change := mgo.Change {
		Update: bson.M{"$set": bson.M{
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

	_, err := mongo.Database.C("feedbacks").FindId(bson.ObjectIdHex(id)).Apply(change, &feedback)
	if err != nil {
		panic(err);
	}
	return &feedback
}

func AddLabelToFeedback(feedbackId string, label *label.Label) *Feedback {
	feedback := GetFeedback(feedbackId)
	if feedback == nil {
		return nil
	}
	feedback.Labels = append(feedback.Labels, label)
	feedback.LabelIds = append(feedback.LabelIds, label.Id)
	change := bson.M{
		"$push": bson.M{"labels": label.Id},
		"$set": bson.M{"updatedat": time.Now()},
	}
	if err := mongo.Database.C("feedbacks").Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
		panic(err)
	}
	return feedback
}

func RemoveLabelFromFeedback(feedbackId string, label *label.Label) *Feedback {
	feedback := GetFeedback(feedbackId)
	if feedback == nil {
		return nil
	}
	change := bson.M{
		"$pull": bson.M{"labels": label.Id},
		"$set": bson.M{"updatedat": time.Now()},
	}
	if err := mongo.Database.C("feedbacks").Update(bson.M{"_id": bson.ObjectIdHex(feedbackId)}, change); err != nil {
		panic(err)
	}
	return feedback
}

func DeleteFeedback(id string) error {
	err := mongo.Database.C("feedbacks").RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}
