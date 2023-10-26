package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Content   string             `json:"content,omitempty" validate:"required"`
	Image     string             `json:"image,omitempty" validate:"required"`
	UserId    primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Views     int                `json:"views,omitempty"`
	Pongs     int                `json:"pongs,omitempty" bson:"pongs,omitempty" default:"1"`
	Comments  []Comment          `json:"comments,omitempty" bson:"comments,omitempty"`
}

type Comment struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content   string             `json:"content,omitempty" validate:"required"`
	UserId    primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Pongs     int                `json:"pongs,omitempty" bson:"pongs,omitempty" default:"1"`
}

type Category struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
	Order int `json:"order,omitempty" validate:"required"`
}