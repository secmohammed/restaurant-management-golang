package services

import (
	"context"
	"time"

	"github.com/secmohammed/restaurant-management/container"
	"github.com/secmohammed/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	InvoiceService interface {
		GetInvoice(id int) (*models.Invoice, error)
		GetInvoices() ([]bson.M, error)
		CreateInvoice()
		UpdateInvoice(id int)
		DeleteInvoice(id int)
	}
	invoiceService struct {
		app *container.App
	}
)

func NewInvoiceService(app *container.App) InvoiceService {
	return &invoiceService{app}
}

func (u *invoiceService) GetInvoice(id int) (*models.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	invoice := &models.Invoice{}
	err := u.app.Database.OpenCollection("invoice").FindOne(ctx, bson.M{"invoice_id": id}).Decode(invoice)
	if err != nil {
		return nil, err
	}
	return invoice, nil
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

func (u *invoiceService) CreateInvoice() {
}

func (u *invoiceService) UpdateInvoice(id int) {
}

func (u *invoiceService) DeleteInvoice(id int) {
}
