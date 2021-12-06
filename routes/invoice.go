package routes

import (
	"github.com/secmohammed/restaurant-management/controllers"
	"github.com/secmohammed/restaurant-management/services"
)

func (r *router) RegisterInvoiceRoutes() {
	s := services.NewInvoiceService(r.app)
	c := controllers.NewInvoiceController(s)
	r.GET("/api/invoices/:id", c.GetInvoice)
	r.POST("/api/invoices", c.CreateInvoice)
	r.GET("/api/invoices", c.GetInvoices)
	r.PUT("/api/invoices/:id", c.UpdateInvoice)
	r.DELETE("/api/invoices/:id", c.DeleteInvoice)
}
