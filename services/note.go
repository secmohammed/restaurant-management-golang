package services

import (
	"context"
	"time"

	"github.com/secmohammed/restaurant-management/models"

	"github.com/secmohammed/restaurant-management/container"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	NoteService interface {
		GetNote(id int) (*models.Note, error)
		GetNotes() ([]bson.M, error)
		CreateNote()
		UpdateNote(id int)
		DeleteNote(id int)
	}
	noteService struct {
		app *container.App
	}
)

func NewNoteService(app *container.App) NoteService {
	return &noteService{app}
}

func (u *noteService) GetNote(id int) (*models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	note := &models.Note{}
	err := u.app.Database.OpenCollection("note").FindOne(ctx, bson.M{"note_id": id}).Decode(note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n *noteService) GetNotes() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := n.app.Database.OpenCollection("note").Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var notes []bson.M
	if err := result.All(ctx, &notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func (u *noteService) CreateNote() {
}

func (u *noteService) UpdateNote(id int) {
}

func (u *noteService) DeleteNote(id int) {
}
