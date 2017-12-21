package model

import(
  "gopkg.in/mgo.v2/bson"
)

type(
  Label struct {
    Id bson.ObjectId `json:"id" bson:"_id"`
    Name string `json:"name"`
    Color string `json:"color"`
  }
  Labels []*Label
  LabelIds []bson.ObjectId
)
