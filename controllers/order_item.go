package controllers

import (
	"net/http"
	"strconv"

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
	o.s.CreateOrderItem()
}

func (o *orderItemController) UpdateOrderItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateOrderItem(id)
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
	results, err := o.s.GetOrderItems()
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
