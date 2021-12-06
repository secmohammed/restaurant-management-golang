package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type menuController struct {
	s services.MenuService
}

type MenuController interface {
	CreateMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	DeleteMenu(c *gin.Context)
	GetMenu(c *gin.Context)
	GetMenus(c *gin.Context)
}

func NewMenuController(o services.MenuService) MenuController {
	return &menuController{o}
}

func (o *menuController) CreateMenu(c *gin.Context) {
	o.s.CreateMenu()
}

func (o *menuController) UpdateMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateMenu(id)
}

func (o *menuController) DeleteMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteMenu(id)
}

func (o *menuController) GetMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.GetMenu(id)
}

func (o *menuController) GetMenus(c *gin.Context) {
	o.s.GetMenus()
}
