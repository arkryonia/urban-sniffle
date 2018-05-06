package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Todo struct holds todo objects
type Todo struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Task string        `bson:"Task" json:"task"`
}
