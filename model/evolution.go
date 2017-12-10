package model;

import(
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
  "time"
)

type(
  Evolution struct {
    Id bson.ObjectId `json:"id" bson:"_id"`
    Feedback `bson:",inline"`
  }
  Evolutions []Evolution
)

func (e *Evolution) MarshalJSON() ([]byte, error) {
  type Alias Evolution
  return json.Marshal(&struct {
      *Alias
      CreatedAt string `json:"created_at"`
      UpdatedAt string `json:"updated_at"`
  }{
      Alias: (*Alias)(e),
      CreatedAt: e.CreatedAt.Format(time.RFC3339),
      UpdatedAt: e.UpdatedAt.Format(time.RFC3339),
  })
}
