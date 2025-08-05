package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string        `json:"username" bson:"username" binding:"required"`
	Email       string        `json:"email" bson:"email" binding:"required"`
	Phonenumber string        `json:"phone_number" bson:"phone_number" binding:"required"`
	Password    string        `json:"password" bson:"password" binding:"required"`
}
