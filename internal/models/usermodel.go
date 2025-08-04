package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Username    string             `json:"username" bson:"username"`
    Email       string             `json:"email" bson:"email"`
    Phonenumber string             `json:"phone_number" bson:"phone_number"`
    Password    string             `json:"password" bson:"password"`
}
