package routes

import (
	"github.com/secmohammed/restaurant-management/controllers"
	"github.com/secmohammed/restaurant-management/services"
)

func (r *router) RegisterTableRoutes() {
	s := services.NewTableService(r.app)
	c := controllers.NewTableController(s)
	r.GET("/api/tables/:id", c.GetTable)
	r.POST("/api/tables", c.CreateTable)
	r.GET("/api/tables", c.GetTables)
	r.PUT("/api/tables/:id", c.UpdateTable)
	r.DELETE("/api/tables/:id", c.DeleteTable)
}
