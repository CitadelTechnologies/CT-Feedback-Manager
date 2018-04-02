package author

import(
	"ct-feedback-manager/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func GetAuthors() Authors {
  	authors := make(Authors, 0)
  	if err := mongo.Database.C("authors").Find(nil).All(&authors); err != nil {
		panic(err)
    }
    return authors
}

func GetAuthor(id string) *Author {
  	var author Author

  	if !bson.IsObjectIdHex(id) {
  		return nil
  	}
  	if err := mongo.Database.C("authors").FindId(bson.ObjectIdHex(id)).One(&author); err != nil {
		panic(err)
    }
    return &author
}

func CreateAuthor(name string, email string) (*Author, error) {
	author := &Author{
		Id: bson.NewObjectId(),
		Username: name,
		Email: email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := mongo.Database.C("authors").Insert(author); err != nil {
		return nil, err
	}
	return author, nil
}
