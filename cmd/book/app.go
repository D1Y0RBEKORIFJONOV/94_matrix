package main

import (
	"94.Metrics/internal/app"
	"94.Metrics/internal/pkg/config"
	"94.Metrics/logger"
)

func main() {
	const op = "cmd/book/app"
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	application := app.NewApp(cfg, log)
	err := application.HTTPApp.Server.RunTLS(application.HTTPApp.HTTPUrl, "./cmd/tls/localhost.pem", "./cmd/tls/localhost-key.pem")
	log.Info("server starting on port ", cfg.HTTPUrl)
	if err != nil {
		log.Error("err", err)
		return
	}
}
