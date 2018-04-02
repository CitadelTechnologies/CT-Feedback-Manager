package label

import(
	"ct-feedback-manager/mongo"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

func GetLabels() Labels {
	labels := make(Labels, 0)
	if err := mongo.Database.C("labels").Find(nil).All(&labels); err != nil {
		panic(err)
	}
	return labels
}

func GetLabel(id string) *Label {
	var label Label

	if !bson.IsObjectIdHex(id) {
		return nil
	}
	if err := mongo.Database.C("labels").FindId(bson.ObjectIdHex(id)).One(&label); err != nil {
		return nil
	}
	return &label
}

func CreateLabel(name string, color string) Label {
	label := Label{
		Id: bson.NewObjectId(),
		Name: name,
		Color: color,
	}
	if err := mongo.Database.C("labels").Insert(label); err != nil {
		panic(err)
	}
	return label
}

func UpdateLabel(id string, data map[string]string) *Label {
	var label Label

	change := mgo.Change {
		Update:    bson.M{"$inc": bson.M{"n": 1}, "$set": bson.M{
			"name": data["name"],
			"color": data["color"],
		}},
	    Upsert:    false,
	    Remove:    false,
	    ReturnNew: true,
	}

	_, err := mongo.Database.C("labels").FindId(bson.ObjectIdHex(id)).Apply(change, &label)
	if err != nil {
		panic(err);
	}
	return &label
}


func DeleteLabel(id string) error {
	err := mongo.Database.C("labels").RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}
