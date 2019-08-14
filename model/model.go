package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person represents person's struct.
type Person struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Age       int                `json:"age" bson:"age"`
}
