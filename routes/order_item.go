package routes

import (
	"github.com/secmohammed/restaurant-management/controllers"
	"github.com/secmohammed/restaurant-management/services"
)

func (r *router) RegisterOrderItemRoutes() {
	s := services.NewOrderItemService(r.app)
	c := controllers.NewOrderItemController(s)
	r.GET("/api/order_items/:id", c.GetOrderItem)
	r.POST("/api/order_items", c.CreateOrderItem)
	r.GET("/api/order_items", c.GetOrderItems)
	r.PUT("/api/order_items/:id", c.UpdateOrderItem)
	r.DELETE("/api/order_items/:id", c.DeleteOrderItem)
}
