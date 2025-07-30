package repository

import (
	"complaint-service/config"
	"complaint-service/internal/model"
)

type NoteRepository interface {
	AddNotes(note *model.Notes) error
}

type noteRepository struct{}

func NewNoteRepository() NoteRepository {
	return &noteRepository{}
}

func (r *noteRepository) AddNotes(note *model.Notes) error {
	result := config.DB.Create(note)
	return result.Error
}
