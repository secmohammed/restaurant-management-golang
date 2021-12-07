package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/services"
	"github.com/secmohammed/restaurant-management/utils"
)

type invoiceController struct {
	s services.InvoiceService
	v *validator.Validate
}

type InvoiceController interface {
	CreateInvoice(c *gin.Context)
	UpdateInvoice(c *gin.Context)
	DeleteInvoice(c *gin.Context)
	GetInvoice(c *gin.Context)
	GetInvoices(c *gin.Context)
}

func NewInvoiceController(o services.InvoiceService, v *validator.Validate) InvoiceController {
	return &invoiceController{o, v}
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

func (i *invoiceController) GetInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	invoice, err := i.s.GetInvoice(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (i *invoiceController) GetInvoices(c *gin.Context) {
	results, err := i.s.GetInvoices()
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
