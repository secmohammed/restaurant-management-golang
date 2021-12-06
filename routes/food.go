package routes

import (
	"github.com/secmohammed/restaurant-management/controllers"
	"github.com/secmohammed/restaurant-management/services"
)

func (r *router) RegisterFoodRoutes() {
	s := services.NewFoodService(r.app)
	c := controllers.NewFoodController(s)
	r.GET("/api/foods/:id", c.GetFood)
	r.POST("/api/foods", c.CreateFood)
	r.GET("/api/foods", c.GetFoods)
	r.PUT("/api/foods/:id", c.UpdateFood)
	r.DELETE("/api/foods/:id", c.DeleteFood)
}
