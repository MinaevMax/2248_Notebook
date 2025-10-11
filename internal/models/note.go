package models

type NotePutData struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type NotePostData struct {
	Message string `json:"message"`
}
