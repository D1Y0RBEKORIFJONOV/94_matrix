package bookusecase

import (
	bookentity "94.Metrics/internal/entity/book"
	"context"
)

type bookInterface interface {
	CreateBook(ctx context.Context, req *bookentity.CreateBookRequest) (*bookentity.Book, error)
	DeleteBook(ctx context.Context, req *bookentity.DeleteBookRequest) error
	GetBook(ctx context.Context, req *bookentity.GetBookRequest) (*bookentity.Book, error)
	GetAllBook(ctx context.Context, req *bookentity.GetAllBooksRequest) ([]*bookentity.Book, error)
	UpdateBook(ctx context.Context, req *bookentity.UpdateBookRequest) (*bookentity.Book, error)
}

type BookUseCase struct {
	book bookInterface
}

func NewBookUseCase(book bookInterface) *BookUseCase {
	return &BookUseCase{
		book: book,
	}
}

func (b *BookUseCase) CreateBook(ctx context.Context, req *bookentity.CreateBookRequest) (*bookentity.Book, error) {
	return b.book.CreateBook(ctx, req)
}

func (b *BookUseCase) DeleteBook(ctx context.Context, req *bookentity.DeleteBookRequest) error {
	return b.book.DeleteBook(ctx, req)
}

func (b *BookUseCase) GetBook(ctx context.Context, req *bookentity.GetBookRequest) (*bookentity.Book, error) {
	return b.book.GetBook(ctx, req)
}
func (b *BookUseCase) GetAllBook(ctx context.Context, req *bookentity.GetAllBooksRequest) ([]*bookentity.Book, error) {
	return b.book.GetAllBook(ctx, req)
}

func (b *BookUseCase) UpdateBook(ctx context.Context, req *bookentity.UpdateBookRequest) (*bookentity.Book, error) {
	return b.book.UpdateBook(ctx, req)
}
