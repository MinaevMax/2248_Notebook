package repository

import (
	"2248_notebook/internal/notebook"
)

type notebookRepo struct {
	storage map[string]string
}

func NewNotebookRepo() notebook.Repository {
	storage := map[string]string{}
	return &notebookRepo{storage: storage}
}

// Функция получния записки из хранилища
func (r *notebookRepo) GetNote(id string) (string, error) {
	value, exists := r.storage[id]
	if !exists {
		return "", notebook.ErrNoteNotFound
	}
	return value, nil
}

// Функция создания новой записки в хранилище
func (r *notebookRepo) PostNote(id string, message string) error {
	if _, exists := r.storage[id]; exists {
		return notebook.ErrNoteExists
	}
	r.storage[id] = message
	return nil
}

// Функция изменения записки из хранилища
func (r *notebookRepo) PutNote(id string, message string) (bool, error) {
	if _, exists := r.storage[id]; !exists {
		return false, notebook.ErrNoteNotFound
	}
	r.storage[id] = message
	return true, nil
}
