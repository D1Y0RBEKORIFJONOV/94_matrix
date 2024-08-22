package bookusecase

import (
	bookentity "94.Metrics/internal/entity/book"
	"context"
)

type (
	saver interface {
		SaveBook(ctx context.Context, book *bookentity.Book) error
	}
	updater interface {
		UpdateBook(ctx context.Context, book *bookentity.Book) error
	}
	provider interface {
		GetBook(ctx context.Context, id string) (*bookentity.Book, error)
		GetAllBooks(ctx context.Context, req *bookentity.GetAllBooksRequest) ([]*bookentity.Book, error)
	}
	deleter interface {
		DeleteBook(ctx context.Context, id string) error
	}
)

type bookRepo interface {
	saver
	updater
	deleter
	provider
}

type BookRepoUseCase struct {
	bookRepo bookRepo
}

func NewBookRepoUseCase(bookRepo bookRepo) *BookRepoUseCase {
	return &BookRepoUseCase{bookRepo: bookRepo}
}
func (repo *BookRepoUseCase) SaveBook(ctx context.Context, book *bookentity.Book) error {
	return repo.bookRepo.SaveBook(ctx, book)
}
func (repo *BookRepoUseCase) UpdateBook(ctx context.Context, book *bookentity.Book) error {
	return repo.bookRepo.UpdateBook(ctx, book)
}
func (repo *BookRepoUseCase) DeleteBook(ctx context.Context, id string) error {
	return repo.bookRepo.DeleteBook(ctx, id)
}
func (repo *BookRepoUseCase) GetAllBooks(ctx context.Context, req *bookentity.GetAllBooksRequest) ([]*bookentity.Book, error) {
	return repo.bookRepo.GetAllBooks(ctx, req)
}
func (repo *BookRepoUseCase) GetBook(ctx context.Context, id string) (*bookentity.Book, error) {
	return repo.bookRepo.GetBook(ctx, id)
}
