package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"2248_notebook/internal/config"
	"2248_notebook/internal/server"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetLogLoggerLevel(slog.LevelDebug)

	if err := godotenv.Load(); err != nil {
		log.Error("No .env file found or error loading")
	}

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		return
	}

	// graceful shutdown
	termCtx, termCancel := context.WithCancel(context.Background())
	go waitSigterm(termCancel, log)

	errCh := make(chan error, 1)

	srv := server.NewServer(&cfg, log)
	err = srv.Run(errCh)
	if err != nil {
		log.Error("failed to run server", slog.Any("error", err))
	}

	select {
	case err := <-errCh:
		log.Error("got error from server", slog.Any("error", err))
	case <-termCtx.Done():
		err := srv.Stop()
		if err != nil {
			log.Error("shutdown failed", slog.Any("error", err))
		}
	}

	log.Info("service terminated")

}

func waitSigterm(terminate context.CancelFunc, log *slog.Logger) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	caughtSignal := <-sigCh

	log.Warn("service starts termination", slog.String("signal", caughtSignal.String()))

	signal.Stop(sigCh)
	close(sigCh)
	terminate()
}
