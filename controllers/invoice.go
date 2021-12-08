package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/secmohammed/restaurant-management/models"

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

func (i *invoiceController) CreateInvoice(c *gin.Context) {
	invoice := models.Invoice{}
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err := i.v.Struct(invoice)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	result, err := i.s.CreateInvoice(invoice)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (i *invoiceController) UpdateInvoice(c *gin.Context) {
	invoice := models.Invoice{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err = i.v.Struct(invoice)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}

	result, err := i.s.UpdateInvoice(id, invoice)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, result)
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
