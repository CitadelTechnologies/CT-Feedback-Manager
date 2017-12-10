package model

import(
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
  "time"
)

type(
  Bug struct {
    Id bson.ObjectId `json:"id" bson:"_id"`
    Feedback `bson:",inline"`
  }
  Bugs []Bug
)

func (b *Bug) MarshalJSON() ([]byte, error) {
  type Alias Bug
  return json.Marshal(&struct {
      *Alias
      CreatedAt string `json:"created_at"`
      UpdatedAt string `json:"updated_at"`
  }{
      Alias: (*Alias)(b),
      CreatedAt: b.CreatedAt.Format(time.RFC3339),
      UpdatedAt: b.UpdatedAt.Format(time.RFC3339),
  })
}
