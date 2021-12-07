package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/models"
	"github.com/secmohammed/restaurant-management/services"
	"github.com/secmohammed/restaurant-management/utils"
)

type orderController struct {
	s services.OrderService
	v *validator.Validate
}

type OrderController interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	GetOrders(c *gin.Context)
}

func NewOrderController(o services.OrderService, v *validator.Validate) OrderController {
	return &orderController{o, v}
}

func (o *orderController) CreateOrder(c *gin.Context) {
	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err := o.v.Struct(order)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	result, err := o.s.CreateOrder(order)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (o *orderController) UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err = o.v.Struct(order)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}

	result, err := o.s.UpdateOrder(id, order)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *orderController) DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteOrder(id)
}

func (o *orderController) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	invoice, err := o.s.GetOrder(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (o *orderController) GetOrders(c *gin.Context) {
	results, err := o.s.GetOrders()
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
