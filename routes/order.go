package routes

import (
	"github.com/secmohammed/restaurant-management/controllers"
	"github.com/secmohammed/restaurant-management/services"
)

func (r *router) RegisterOrderRoutes() {
	s := services.NewOrderService(r.app)
	c := controllers.NewOrderController(s)
	r.GET("/api/orders/:id", c.GetOrder)
	r.GET("/api/orders", c.GetOrders)
	r.POST("/api/orders", c.CreateOrder)
	r.PUT("/api/orders/:id", c.UpdateOrder)
	r.DELETE("/api/orders/:id", c.DeleteOrder)
}
