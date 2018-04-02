package comment

import(
    "ct-feedback-manager/author"
    "time"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type(
    Comment struct {
        Id bson.ObjectId `json:"id"`
        Author *author.Author `json:"author"`
        Content string `json:"content"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
    }
    Comments []Comment
)

func (c *Comment) MarshalJSON() ([]byte, error) {
    type Alias Comment
    return json.Marshal(&struct {
        CreatedAt string `json:"created_at"`
        UpdatedAt string `json:"updated_at"`
        *Alias
    } {
        CreatedAt: c.CreatedAt.Format(time.RFC3339),
        UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
        Alias: (*Alias)(c),
    })
}
