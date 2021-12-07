package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/models"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
	"github.com/secmohammed/restaurant-management/utils"
)

type menuController struct {
	s services.MenuService
	v *validator.Validate
}

type MenuController interface {
	CreateMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	DeleteMenu(c *gin.Context)
	GetMenu(c *gin.Context)
	GetMenus(c *gin.Context)
}

func NewMenuController(o services.MenuService, v *validator.Validate) MenuController {
	return &menuController{o, v}
}

func (m *menuController) CreateMenu(c *gin.Context) {
	menu := models.Menu{}
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err := m.v.Struct(menu)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	result, err := m.s.CreateMenu(menu)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (m *menuController) UpdateMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	menu := models.Menu{}
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err = m.v.Struct(menu)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}

	result, err := m.s.UpdateMenu(id, menu)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *menuController) DeleteMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	o.s.DeleteMenu(id)
}

func (m *menuController) GetMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	menu, err := m.s.GetMenu(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, menu)
}

func (o *menuController) GetMenus(c *gin.Context) {
	results, err := o.s.GetMenus()
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
