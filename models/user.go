package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    *string            `json:"first_name" validate:"required,min=3,max=50"`
	LastName     *string            `json:"last_name" validate:"required,min=3,max=50"`
	Password     *string            `json:"password" validate:"required,min=6,max=50"`
	Email        *string            `json:"email" validate:"required,email,min=3,max=50"`
	Avatar       *string            `json:"avatar"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}
