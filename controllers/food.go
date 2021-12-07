package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/models"
	"github.com/secmohammed/restaurant-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type foodController struct {
	s services.FoodService
	v *validator.Validate
}

type FoodController interface {
	CreateFood(c *gin.Context)
	UpdateFood(c *gin.Context)
	DeleteFood(c *gin.Context)
	GetFood(c *gin.Context)
	GetFoods(c *gin.Context)
}

func NewFoodController(o services.FoodService, v *validator.Validate) FoodController {
	return &foodController{o, v}
}

func (f *foodController) CreateFood(c *gin.Context) {
	food := models.Food{}
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err := f.v.Struct(food)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	result, err := f.s.CreateFood(food)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (f *foodController) UpdateFood(c *gin.Context) {
	food := models.Food{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err = f.v.Struct(food)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}

	result, err := f.s.UpdateFood(id, food)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *foodController) DeleteFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	o.s.DeleteFood(id)
}

func (f *foodController) GetFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	food, err := f.s.GetFood(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, food)
}

func (f *foodController) GetFoods(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 0 {
		limit = 10
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		page = 1
	}
	results, err := f.s.GetFoods(limit, page)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
