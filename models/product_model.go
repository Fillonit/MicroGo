package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Price     float64            `json:"price,omitempty" validate:"required"` // Add this field
	Stock     int                `json:"stock,omitempty" validate:"required"`
	Image     string             `json:"image,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Stored    []Store            `json:"stored,omitempty" bson:"stored,omitempty"`
}

type Store struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	StoreId   primitive.ObjectID `json:"storeId,omitempty" bson:"storeId,omitempty"`
	ProductId primitive.ObjectID `json:"productId,omitempty" bson:"productId,omitempty"`
	Stock     int                `json:"stock,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
