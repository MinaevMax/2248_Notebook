package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"2248_notebook/internal/config"
	notebookHttp "2248_notebook/internal/notebook/delivery/http"
	notebookRepository "2248_notebook/internal/notebook/repository"
	notebookUsecase "2248_notebook/internal/notebook/usecase"

	"github.com/gorilla/mux"
)

type Server struct {
	cfg *config.Config
	log *slog.Logger
	srv *http.Server
}

func NewServer(cfg *config.Config, log *slog.Logger) *Server {
	return &Server{
		cfg: cfg,
		log: log,
	}
}

func (s *Server) Run(errCh chan error) error {
	// Init repositories
	notebookRepo := notebookRepository.NewNotebookRepo()

	// Init useCases
	notebookUC := notebookUsecase.NewNotebookUC(notebookRepo, s.log)

	// Init handlers
	notebookHandler := notebookHttp.NewHandler(notebookUC, s.log)

	r := mux.NewRouter()
	
	apiRouter := r.PathPrefix("/notebook").Subrouter()

	notebookHttp.MapRoutes(apiRouter, notebookHandler)

	s.srv = &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
		Handler:      r,
	}

	go func() {
		s.log.Info("Starting server", slog.Int("port", s.cfg.Server.Port))
		err := s.srv.ListenAndServe()
		if err != nil {
			errCh <- fmt.Errorf("listen and server error: %w", err)
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
