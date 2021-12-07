package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Size        *string            `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Price       *string            `json:"price" validate:"required,min=3,max=50"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	OrderItemID string             `json:"order_item_id"`
	FoodID      *string            `json:"food_id" validate:"required"`
	OrderID     *string            `json:"order_id" validate:"required"`
}
