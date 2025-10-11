package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"2248_notebook/internal/models"
	"2248_notebook/internal/notebook"
	"2248_notebook/internal/utils"
)

type handler struct {
	log *slog.Logger
	uc  notebook.UseCase
}

func NewHandler(uc notebook.UseCase, log *slog.Logger) notebook.Handler {
	return &handler{
		log: log,
		uc:  uc,
	}
}

// Ручка для получения записки
func (h *handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Info("Received GET request")
		id := r.URL.Query().Get("noteId")
		if id == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "Empty note ID")
			return
		}

		note, err := h.uc.GetNote(id)
		if err != nil {
			h.log.Error("failed to get note", slog.Any("error", err))
			utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to get note")
			return
		}
		utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"Your note": note})
	}
}

// Ручка для создания новой записки
func (h *handler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Info("Received POST request")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.log.Error("failed to read body", slog.Any("error", err))
			utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to read paged note body")
			return
		}
		defer r.Body.Close()

		noteMessage := models.NotePostData{}
		if err := json.Unmarshal(body, &noteMessage); err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}

		noteId, err := h.uc.PostNote(noteMessage)
		if err != nil {
			h.log.Error("failed to post note", slog.Any("error", err))
			utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to post note")
			return
		}
		utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"Note ID": noteId})
	}
}

// Ручка для изменения существующщей записки
func (h *handler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Info("Received PUT request")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.log.Error("failed to read body", slog.Any("error", err))
			utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to read paged note body")
			return
		}
		defer r.Body.Close()

		noteData := models.NotePutData{}
		if err := json.Unmarshal(body, &noteData); err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}

		success, err := h.uc.PutNote(noteData)
		if err != nil {
			h.log.Error("failed to put note", slog.Any("error", err))
			utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to put note")
			return
		}
		if !success {
			utils.WriteJSONError(w, http.StatusBadRequest, "Failed to put note")
			return
		}
		utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"Status": "OK"})
	}
}
