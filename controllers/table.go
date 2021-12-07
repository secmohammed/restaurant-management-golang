package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/services"
	"github.com/secmohammed/restaurant-management/utils"
)

type tableController struct {
	s services.TableService
	v *validator.Validate
}

type TableController interface {
	CreateTable(c *gin.Context)
	UpdateTable(c *gin.Context)
	DeleteTable(c *gin.Context)
	GetTable(c *gin.Context)
	GetTables(c *gin.Context)
}

func NewTableController(o services.TableService, v *validator.Validate) TableController {
	return &tableController{o, v}
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

func (t *tableController) GetTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	invoice, err := t.s.GetTable(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (t *tableController) GetTables(c *gin.Context) {
	results, err := t.s.GetTables()
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
