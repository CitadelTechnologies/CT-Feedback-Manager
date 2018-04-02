package author

import(
    "gopkg.in/mgo.v2/bson"
    "time"
)

type(
    Author struct{
        Id bson.ObjectId `json:"id"`
        Username string `json:"username"`
        Email string `json:"email"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
    }
    Authors []Author
)
