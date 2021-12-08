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

type orderItemController struct {
	s services.OrderItemService
	v *validator.Validate
}

type OrderItemController interface {
	CreateOrderItem(c *gin.Context)
	UpdateOrderItem(c *gin.Context)
	DeleteOrderItem(c *gin.Context)
	GetOrderItem(c *gin.Context)
	GetOrderItems(c *gin.Context)
}

func NewOrderItemController(o services.OrderItemService, v *validator.Validate) OrderItemController {
	return &orderItemController{o, v}
}

func (o *orderItemController) CreateOrderItem(c *gin.Context) {
	order := models.OrderItemPack{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	err := o.v.Struct(order)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	id := "1" // TODO Fix
	result, err := o.s.CreateOrderItem(order, &id)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (o *orderItemController) UpdateOrderItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	order := models.OrderItem{}

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err = o.v.Struct(order)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}

	result, err := o.s.UpdateOrderItem(id, order)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *orderItemController) DeleteOrderItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteOrderItem(id)
}

func (o *orderItemController) GetOrderItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	invoice, err := o.s.GetOrderItem(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (o *orderItemController) GetOrderItems(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	results, err := o.s.ItemsByOrder(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
