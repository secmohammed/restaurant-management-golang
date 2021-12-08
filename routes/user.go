package routes

import (
	"github.com/secmohammed/restaurant-management/controllers"
	"github.com/secmohammed/restaurant-management/services"
)

func (r *router) RegisterUserRoutes() {
	s := services.NewUserService(r.app)
	c := controllers.NewUserController(s, r.app.Validator)
	r.GET("/api/users/:id", c.GetUser)
	r.POST("/api/signup", c.Signup)
	r.POST("/api/login", c.Login)
	r.GET("/api/users", c.GetUsers)
	r.PUT("/api/users/:id", c.UpdateUser)
	r.DELETE("/api/users/:id", c.DeleteUser)
}
