package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `json:"name" validate:"required,min=3,max=50"`
	Category  string             `json:"category" validate:"required,min=3,max=50"`
	StartDate *time.Time         `json:"start_date" validate:"required,min=2,max=100"`
	EndDate   *time.Time         `json:"end_date" validate:"required,min=2,max=100"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	MenuID    string             `json:"menu_id"`
}
