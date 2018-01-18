package models

import "gopkg.in/mgo.v2/bson"

// Represents a check-in, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type CheckIn struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Location string        `bson:"location" json:"location"`
	Stamp    string	       `bson:"timestamp" json:"timestamp"`	
}

