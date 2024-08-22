package bookservice

import (
	bookentity "94.Metrics/internal/entity/book"
	bookusecase "94.Metrics/internal/usecase/book"
	"context"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"log/slog"
	"time"
)

type Book struct {
	book   *bookusecase.BookRepoUseCase
	logger *slog.Logger
}

func NewBookService(book *bookusecase.BookRepoUseCase, logger *slog.Logger) *Book {
	return &Book{
		book:   book,
		logger: logger,
	}
}
func (b *Book) CreateBook(ctx context.Context, req *bookentity.CreateBookRequest) (*bookentity.Book, error) {
	const op = "Service.Book.CreateBook"
	log := b.logger.With(slog.String("method", op))
	log.Info("Starting")
	defer log.Info("Ending")
	book := bookentity.Book{
		Id:            uuid.NewString(),
		Title:         req.Title,
		Author:        req.Author,
		PublishedDate: time.Now(),
		Isbn:          req.Isbn,
	}
	err := b.book.SaveBook(ctx, &book)
	if err != nil {
		log.Error("err", err)
		return nil, xerrors.Errorf("%s: %w", op, err)
	}
	return &book, nil
}

func (b *Book) DeleteBook(ctx context.Context, req *bookentity.DeleteBookRequest) error {
	const op = "Service.Book.DeleteBook"
	log := b.logger.With(slog.String("method", op))
	log.Info("Starting")
	defer log.Info("Ending")
	err := b.book.DeleteBook(ctx, req.Id)
	if err != nil {
		log.Error("err", err)
		return xerrors.Errorf("%s: %w", op, err)
	}
	return nil
}

func (b *Book) GetBook(ctx context.Context, req *bookentity.GetBookRequest) (*bookentity.Book, error) {
	const op = "Service.Book.GetBook"
	log := b.logger.With(slog.String("method", op))
	log.Info("Starting")
	defer log.Info("Ending")
	book, err := b.book.GetBook(ctx, req.Id)
	if err != nil {
		log.Error("err", err)
		return nil, xerrors.Errorf("%s: %w", op, err)
	}
	return book, nil
}

func (b *Book) GetAllBook(ctx context.Context, req *bookentity.GetAllBooksRequest) ([]*bookentity.Book, error) {
	const op = "Service.Book.GetAllBook"
	log := b.logger.With(slog.String("method", op))
	log.Info("Starting")
	defer log.Info("Ending")

	books, err := b.book.GetAllBooks(ctx, req)
	if err != nil {
		log.Error("err", err)
		return nil, xerrors.Errorf("%s: %w", op, err)
	}
	return books, nil
}

func (b *Book) UpdateBook(ctx context.Context, req *bookentity.UpdateBookRequest) (*bookentity.Book, error) {
	const op = "Service.Book.UpdateBook"
	log := b.logger.With(slog.String("method", op))
	log.Info("Starting")
	defer log.Info("Ending")
	book, err := b.book.GetBook(ctx, req.Id)
	if err != nil {
		log.Error("err", err)
		return nil, xerrors.Errorf("%s: %w", op, err)
	}
	if req.Isbn != "" {
		book.Isbn = req.Isbn
	}
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	err = b.book.UpdateBook(ctx, book)
	if err != nil {
		log.Error("err", err)
		return nil, xerrors.Errorf("%s: %w", op, err)
	}
	return book, nil
}
