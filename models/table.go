package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Number    *int               `json:"number" validate:"required"`
	Capacity  *int               `json:"capacity" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	TableID   string             `json:"table_id"`
}
