package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type tableController struct {
	s services.TableService
}

type TableController interface {
	CreateTable(c *gin.Context)
	UpdateTable(c *gin.Context)
	DeleteTable(c *gin.Context)
	GetTable(c *gin.Context)
	GetTables(c *gin.Context)
}

func NewTableController(o services.TableService) TableController {
	return &tableController{o}
}

func (o *tableController) CreateTable(c *gin.Context) {
	o.s.CreateTable()
}

func (o *tableController) UpdateTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateTable(id)
}

func (o *tableController) DeleteTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteTable(id)
}

func (o *tableController) GetTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.GetTable(id)
}

func (o *tableController) GetTables(c *gin.Context) {
	o.s.GetTables()
}
