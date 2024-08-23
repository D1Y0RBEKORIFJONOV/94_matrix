package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"94.Metrics/internal/app"
	"94.Metrics/internal/pkg/config"
	"94.Metrics/logger"
)

func main() {
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	application := app.NewApp(cfg, log)

	srv := &http.Server{
		Addr:    cfg.HTTPUrl,
		Handler: application.HTTPApp.Server,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("server starting on port ", cfg.HTTPUrl)
		if err := srv.ListenAndServeTLS("localhost.pem", "localhost-key.pem"); err != nil && err != http.ErrServerClosed {
			log.Error("listen: %s\n", err)
		}
	}()

	<-quit
	log.Info("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown:", err)
	}

	log.Info("Server exiting")
}
