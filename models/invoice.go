package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	PaymentMethod  *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus  *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID|eq="`
	PaymentDueDate time.Time          `json:"payment_due_date" validate:"required,min=2,max=100"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	InvoiceID      string             `json:"invoice_id"`
	OrderID        *string            `json:"order_id"`
}
