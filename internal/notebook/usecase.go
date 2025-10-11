package notebook

import "2248_notebook/internal/models"

type UseCase interface {
	GetNote(id string) (string, error)
	PostNote(message models.NotePostData) (string, error)
	PutNote(noteData  models.NotePutData) (bool, error)
}
