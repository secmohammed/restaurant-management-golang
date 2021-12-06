package services

import "github.com/secmohammed/restaurant-management/container"

type (
	InvoiceService interface {
		GetInvoice(id int)
		GetInvoices()
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

func (u *invoiceService) GetInvoice(id int) {
}

func (u *invoiceService) GetInvoices() {
}

func (u *invoiceService) CreateInvoice() {
}

func (u *invoiceService) UpdateInvoice(id int) {
}

func (u *invoiceService) DeleteInvoice(id int) {
}
