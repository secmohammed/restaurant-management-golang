package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type foodController struct {
	s services.FoodService
}

type FoodController interface {
	CreateFood(c *gin.Context)
	UpdateFood(c *gin.Context)
	DeleteFood(c *gin.Context)
	GetFood(c *gin.Context)
	GetFoods(c *gin.Context)
}

func NewFoodController(o services.FoodService) FoodController {
	return &foodController{o}
}

func (o *foodController) CreateFood(c *gin.Context) {
	o.s.CreateFood()
}

func (o *foodController) UpdateFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateFood(id)
}

func (o *foodController) DeleteFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteFood(id)
}

func (o *foodController) GetFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.GetFood(id)
}

func (o *foodController) GetFoods(c *gin.Context) {
	o.s.GetFoods()
}
