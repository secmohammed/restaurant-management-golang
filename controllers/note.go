package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/services"
	"github.com/secmohammed/restaurant-management/utils"
)

type noteController struct {
	s services.NoteService
	v *validator.Validate
}

type NoteController interface {
	CreateNote(c *gin.Context)
	UpdateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
	GetNote(c *gin.Context)
	GetNotes(c *gin.Context)
}

func NewNoteController(n services.NoteService, v *validator.Validate) NoteController {
	return &noteController{n, v}
}

func (o *noteController) CreateNote(c *gin.Context) {
	o.s.CreateNote()
}

func (o *noteController) UpdateNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.UpdateNote(id)
}

func (o *noteController) DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.DeleteNote(id)
}

func (n *noteController) GetNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	note, err := n.s.GetNote(id)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, note)
}

func (n *noteController) GetNotes(c *gin.Context) {
	results, err := n.s.GetNotes()
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, results)
}
