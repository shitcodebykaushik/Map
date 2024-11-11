package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    FullName    string             `bson:"full_name" json:"full_name" binding:"required"`
    Email       string             `bson:"email" json:"email" binding:"required,email"`
    PhoneNumber string             `bson:"phone_number" json:"phone_number" binding:"required"`
    Password    string             `bson:"password" json:"password" binding:"required"`
}
