package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      *string            `json:"name" validate:"required,min=2,max=100"`
	Price     *float64           `json:"price" validate:"required,min=0"`
	Image     *string            `json:"image" validate:"required,min=2,max=100"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	MenuID    *string            `json:"menu_id" validate:"required"`
	FoodID    string             `json:"food_id"`
}
