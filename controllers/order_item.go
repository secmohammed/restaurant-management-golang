package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type orderItemController struct {
	s services.OrderItemService
}

type OrderItemController interface {
	CreateOrderItem(c *gin.Context)
	UpdateOrderItem(c *gin.Context)
	DeleteOrderItem(c *gin.Context)
	GetOrderItem(c *gin.Context)
	GetOrderItems(c *gin.Context)
}

func NewOrderItemController(o services.OrderItemService) OrderItemController {
	return &orderItemController{o}
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
	o.s.GetOrderItem(id)
}

func (o *orderItemController) GetOrderItems(c *gin.Context) {
	o.s.GetOrderItems()
}
