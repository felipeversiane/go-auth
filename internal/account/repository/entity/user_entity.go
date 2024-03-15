package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	FirstName string             `bson:"first_name,omitempty"`
	LastName  string             `bson:"last_name,omitempty"`
	Age       int8               `bson:"age,omitempty"`
}
