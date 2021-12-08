package controllers

import (
	"net/http"
	"strconv"

	"github.com/secmohammed/restaurant-management/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/services"
	"github.com/secmohammed/restaurant-management/utils"
)

type userController struct {
	s services.UserService
	v *validator.Validate
}

type UserController interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
}

func NewUserController(o services.UserService, v *validator.Validate) UserController {
	return &userController{o, v}
}

func (u *userController) Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	err := u.v.Struct(user)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	result, err := u.s.Signup(user)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (u *userController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	err := u.v.Struct(user)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	result, err := u.s.Login(user)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *userController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateUser(id)
}

func (o *userController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteUser(id)
}

func (u *userController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	invoice, err := u.s.GetUser(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (u *userController) GetUsers(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 0 {
		limit = 10
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 0 {
		page = 1
	}
	results, err := u.s.GetUsers(limit, page)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusNotFound, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
