package httpapp

import (
	"94.Metrics/internal/http/router"
	"94.Metrics/internal/pkg/config"
	bookusecase "94.Metrics/internal/usecase/book"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type HTTPApp struct {
	Logger  *slog.Logger
	HTTPUrl string
	Server  *gin.Engine
}

func NewHTTPApp(logger *slog.Logger, cfg *config.Config, useCase *bookusecase.BookUseCase) *HTTPApp {
	server := router.NewRouter(useCase)
	return &HTTPApp{
		Logger:  logger,
		HTTPUrl: cfg.HTTPUrl,
		Server:  server,
	}
}
