package app

import (
	httpapp "94.Metrics/internal/app/http"
	"94.Metrics/internal/infastructure/repository/postgresql"
	"94.Metrics/internal/pkg/config"
	"94.Metrics/internal/pkg/postgres"
	bookservice "94.Metrics/internal/services/book"
	bookusecase "94.Metrics/internal/usecase/book"
	"log"
	"log/slog"
)

type App struct {
	HTTPApp *httpapp.HTTPApp
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgresql.NewBookRepository(db, logger)
	repoUseCase := bookusecase.NewBookRepoUseCase(repo)

	bookService := bookservice.NewBookService(repoUseCase, logger)
	bookHandlerUseCase := bookusecase.NewBookUseCase(bookService)

	httpServer := httpapp.NewHTTPApp(logger, cfg, bookHandlerUseCase)
	return &App{
		HTTPApp: httpServer,
	}
}
