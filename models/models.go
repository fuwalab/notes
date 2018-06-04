package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Note struct {
	ObjectID bson.ObjectId `bson:"_id,omitempty"`
	ID string `bson:"ID,omitempty"`
	Title *string `bson:"Title"`
	Content string `bson:"Content,omitempty"`
	UA *string `bson:"UA"`
	IP *string `bson:"IP"`
	CreatedAt time.Time `bson:"CreatedAt,omitempty"`
}
