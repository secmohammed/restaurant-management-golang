package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/secmohammed/restaurant-management/container"
	"github.com/secmohammed/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InvoiceService interface {
		GetInvoice(id int) (*InvoiceViewFormat, error)
		GetInvoices() ([]bson.M, error)
		CreateInvoice(invoice models.Invoice) (*mongo.InsertOneResult, error)
		UpdateInvoice(id int, invoice models.Invoice) (*mongo.UpdateResult, error)
		DeleteInvoice(id int)
	}
	invoiceService struct {
		app              *container.App
		orderItemService OrderItemService
	}
	InvoiceViewFormat struct {
		InvoiceID      string `json:"invoice_id"`
		PaymentMethod  string
		OrderID        string
		PaymentStatus  *string
		PaymentDue     interface{}
		TableNumber    interface{}
		PaymentDueDate time.Time
		OrderDetails   interface{}
	}
)

func NewInvoiceService(app *container.App) InvoiceService {
	return &invoiceService{app: app, orderItemService: NewOrderItemService(app)}
}

func (i *invoiceService) GetInvoice(id int) (*InvoiceViewFormat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	invoice := &models.Invoice{}
	err := i.app.Database.OpenCollection("invoice").FindOne(ctx, bson.M{"invoice_id": id}).Decode(invoice)
	if err != nil {
		return nil, err
	}
	invoiceView := InvoiceViewFormat{}
	orderId, err := strconv.Atoi(*invoice.OrderID)
	if err != nil {
		return nil, err
	}
	orderItems, err := i.orderItemService.ItemsByOrder(orderId)
	if err != nil {
		return nil, err
	}
	invoiceView.OrderID = *invoice.OrderID
	invoiceView.PaymentMethod = "null"
	if invoice.PaymentMethod != nil {
		invoiceView.PaymentMethod = *invoice.PaymentMethod
	}
	invoiceView.PaymentDueDate = invoice.PaymentDueDate
	invoiceView.InvoiceID = invoice.InvoiceID
	invoiceView.PaymentStatus = invoice.PaymentStatus
	invoiceView.PaymentDue = orderItems[0]["payment_due"]
	invoiceView.TableNumber = orderItems[0]["table_number"]
	invoiceView.OrderDetails = orderItems[0]["order_items"]
	return &invoiceView, nil
}

func (u *invoiceService) GetInvoices() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := u.app.Database.OpenCollection("invoice").Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var invoices []bson.M
	if err := result.All(ctx, &invoices); err != nil {
		return nil, err
	}
	return invoices, nil
}

func (i *invoiceService) CreateInvoice(invoice models.Invoice) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	order := models.Order{}

	if invoice.OrderID != nil {
		err := i.app.Database.OpenCollection("order").FindOne(ctx, bson.M{"order_id": *&order.OrderID}).Decode(&order)
		if err != nil {
			return nil, errors.New("order not found")
		}
	}

	invoice.PaymentDueDate, _ = time.Parse(time.RFC3339, time.Now().AddDate(0, 0, 1).Format(time.RFC3339))
	invoice.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	invoice.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	invoice.ID = primitive.NewObjectID()
	invoice.InvoiceID = invoice.ID.Hex()
	result, err := i.app.Database.OpenCollection("invoice").InsertOne(ctx, invoice)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *invoiceService) UpdateInvoice(id int, invoice models.Invoice) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{"order_id": id}
	var updateObj primitive.D
	if invoice.PaymentStatus != nil {
		updateObj = append(updateObj, bson.E{Key: "payment_status", Value: invoice.PaymentStatus})
	} else {
		updateObj = append(updateObj, bson.E{Key: "payment_status", Value: "PENDING"})
	}
	if invoice.PaymentMethod != nil {
		updateObj = append(updateObj, bson.E{Key: "payment_method", Value: invoice.PaymentMethod})
	}
	invoice.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: invoice.UpdatedAt})
	upsert := true
	opt := options.UpdateOptions{Upsert: &upsert}
	result, err := i.app.Database.OpenCollection("invoice").UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	if err != nil {
		return nil, errors.New("failed to update invoice")
	}
	return result, nil
}

func (u *invoiceService) DeleteInvoice(id int) {
}
