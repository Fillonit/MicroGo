package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Location  string             `json:"location,omitempty" validate:"required"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Company   string             `json:"company,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" validate:"required,email"`
	Password  string             `json:"password,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
