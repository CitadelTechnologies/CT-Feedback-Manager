package model;

import(
  "time"
)

type Feedback struct {
  Author Author `json:"author"`
  Title string `json:"title"`
  Description string `json:"description"`
  Labels Labels `json:"labels" bson:"_"`
  LabelIds LabelIds `json:"_"`
  Comments Comments `json:"comments"`
  Status string `json:"status"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
