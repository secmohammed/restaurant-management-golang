package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type orderController struct {
	s services.OrderService
}

type OrderController interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	GetOrders(c *gin.Context)
}

func NewOrderController(o services.OrderService) OrderController {
	return &orderController{o}
}

func (o *orderController) CreateOrder(c *gin.Context) {
	o.s.CreateOrder()
}

func (o *orderController) UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateOrder(id)
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
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.GetOrder(id)
}

func (o *orderController) GetOrders(c *gin.Context) {
	o.s.GetOrders()
}
