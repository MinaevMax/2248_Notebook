package usecase

import (
	"log/slog"

	"2248_notebook/internal/models"
	"2248_notebook/internal/notebook"

	"github.com/google/uuid"
)

type notebookUC struct {
	notebookRepo notebook.Repository
	log          *slog.Logger
}

func NewNotebookUC(notebookRepo notebook.Repository, log *slog.Logger) notebook.UseCase {
	return &notebookUC{
		notebookRepo: notebookRepo,
		log:          log,
	}
}

func (uc *notebookUC) GetNote(id string) (string, error) {
	note, err := uc.notebookRepo.GetNote(id)
	if err != nil {
		return "", err
	}
	return note, nil
}

func (uc *notebookUC) PostNote(message models.NotePostData) (string, error) {
	newUUID := uuid.New().String()
	err := uc.notebookRepo.PostNote(newUUID, message.Message)
	if err != nil {
		return "", err
	}
	return newUUID, nil
}

func (uc *notebookUC) PutNote(noteData models.NotePutData) (bool, error) {
	success, err := uc.notebookRepo.PutNote(noteData.ID, noteData.Message)
	if err != nil {
		return false, err
	}
	return success, nil
}
