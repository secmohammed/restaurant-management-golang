package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/secmohammed/restaurant-management/services"
)

type noteController struct {
	s services.NoteService
}

type NoteController interface {
	CreateNote(c *gin.Context)
	UpdateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
	GetNote(c *gin.Context)
	GetNotes(c *gin.Context)
}

func NewNoteController(o services.NoteService) NoteController {
	return &noteController{o}
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

func (o *noteController) GetNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	o.s.GetNote(id)
}

func (o *noteController) GetNotes(c *gin.Context) {
	o.s.GetNotes()
}
