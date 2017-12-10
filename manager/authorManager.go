package manager

import(
	"ct-feedback-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
)

/*
* @return model.Authors
*/
func GetAuthors() model.Authors {
  	authors := make(model.Authors, 0)
  	if err := MongoDBConnection.DB(MongoDBName).C("authors").Find(nil).All(&authors); err != nil {
      panic(err)
    }
    return authors
}

/*
* @return model.Author
*/
func GetAuthor(id string) *model.Author {
  	var author model.Author

  	if !bson.IsObjectIdHex(id) {
  		return nil
  	}
  	if err := MongoDBConnection.DB(MongoDBName).C("authors").FindId(bson.ObjectIdHex(id)).One(&author); err != nil {
      panic(err)
    }
    return &author
}

/*
* @param string name
* @param string description
* @return model.Bug
*/
func CreateAuthor(name string, email string) model.Author {
	var author model.Author

	author.Id = bson.NewObjectId()
  author.Username = name
  author.Email = email
	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()

  if err := MongoDBConnection.DB(MongoDBName).C("authors").Insert(author); err != nil {
    panic(err)
  }
  return author
}
