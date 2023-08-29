package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" validate:"required"`
	Content string             `json:"content,omitempty" validate:"required"`
	Image   string             `json:"image,omitempty" validate:"required"`
	UserId  primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
}
