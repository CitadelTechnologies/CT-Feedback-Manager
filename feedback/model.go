package feedback

import(
    "ct-feedback-manager/author"
    "ct-feedback-manager/comment"
    "ct-feedback-manager/label"
    "time"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

const (
    FEEDBACK_TYPE_EVOLUTION = "evolution"
    FEEDBACK_TYPE_BUG = "bug"
)

type Feedback struct {
    Id bson.ObjectId `json:"id" bson:"_id"`
    Author *author.Author `json:"author"`
    Title string `json:"title"`
    Slug string `json:"slug"`
    Type string `json:"type"`
    Description string `json:"description"`
    Labels label.Labels `json:"labels" bson:"-"`
    LabelIds label.LabelIds `json:"-" bson:"labels"`
    Comments comment.Comments `json:"comments"`
    Status string `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (f *Feedback) MarshalJSON() ([]byte, error) {
    type Alias Feedback
    return json.Marshal(&struct {
        *Alias
        CreatedAt string `json:"created_at"`
        UpdatedAt string `json:"updated_at"`
    }{
        Alias: (*Alias)(f),
        CreatedAt: f.CreatedAt.Format(time.RFC3339),
        UpdatedAt: f.UpdatedAt.Format(time.RFC3339),
    })
}
