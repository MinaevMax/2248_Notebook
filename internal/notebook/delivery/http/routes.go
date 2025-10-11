package http

import (
	"net/http"

	"2248_notebook/internal/notebook"
	"github.com/gorilla/mux"
)

func MapRoutes(m *mux.Router, h notebook.Handler) {
	m.HandleFunc("/get", h.Get()).Methods(http.MethodGet)
	m.HandleFunc("/post", h.Post()).Methods(http.MethodPost)
	m.HandleFunc("/put", h.Put()).Methods(http.MethodPut)
}
