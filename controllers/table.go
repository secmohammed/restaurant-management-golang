package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/secmohammed/restaurant-management/models"

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

func (t *tableController) CreateTable(c *gin.Context) {
	table := models.Table{}
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err := t.v.Struct(table)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	result, err := t.s.CreateTable(table)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (t *tableController) UpdateTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	table := models.Table{}
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err = t.v.Struct(table)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}

	result, err := t.s.UpdateTable(id, table)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, result)
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
