package routes

import (
	"github.com/secmohammed/restaurant-management/controllers"
	"github.com/secmohammed/restaurant-management/services"
)

func (r *router) RegisterMenuRoutes() {
	s := services.NewMenuService(r.app)
	c := controllers.NewMenuController(s, r.app.Validator)
	r.GET("/api/menus/:id", c.GetMenu)
	r.POST("/api/menus", c.CreateMenu)
	r.GET("/api/menus", c.GetMenus)
	r.PUT("/api/menus/:id", c.UpdateMenu)
	r.DELETE("/api/menus/:id", c.DeleteMenu)
}
