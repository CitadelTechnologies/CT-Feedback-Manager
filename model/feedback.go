package model;

import(
  "time"
)

type Feedback struct {
  Author Author `json:"author"`
  Title string `json:"title"`
  Description string `json:"description"`
  Comments Comments `json:"comments"`
  Status string `json:"status"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
