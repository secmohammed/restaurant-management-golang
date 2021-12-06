package services

import "github.com/secmohammed/restaurant-management/container"

type (
	NoteService interface {
		GetNote(id int)
		GetNotes()
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

func (u *noteService) GetNote(id int) {
}

func (u *noteService) GetNotes() {
}

func (u *noteService) CreateNote() {
}

func (u *noteService) UpdateNote(id int) {
}

func (u *noteService) DeleteNote(id int) {
}
