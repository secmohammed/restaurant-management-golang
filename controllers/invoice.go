package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type invoiceController struct {
	s services.InvoiceService
}

type InvoiceController interface {
	CreateInvoice(c *gin.Context)
	UpdateInvoice(c *gin.Context)
	DeleteInvoice(c *gin.Context)
	GetInvoice(c *gin.Context)
	GetInvoices(c *gin.Context)
}

func NewInvoiceController(o services.InvoiceService) InvoiceController {
	return &invoiceController{o}
}

func (o *invoiceController) CreateInvoice(c *gin.Context) {
	o.s.CreateInvoice()
}

func (o *invoiceController) UpdateInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateInvoice(id)
}

func (o *invoiceController) DeleteInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteInvoice(id)
}

func (o *invoiceController) GetInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.GetInvoice(id)
}

func (o *invoiceController) GetInvoices(c *gin.Context) {
	o.s.GetInvoices()
}
