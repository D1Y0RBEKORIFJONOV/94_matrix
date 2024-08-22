package postgresql

import (
	bookentity "94.Metrics/internal/entity/book"
	"94.Metrics/internal/pkg/postgres"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log/slog"
)

type BookRepository struct {
	db        *postgres.PostgresDB
	tableName string
	log       *slog.Logger
}

func NewBookRepository(db *postgres.PostgresDB, logger *slog.Logger) *BookRepository {
	return &BookRepository{
		db:        db,
		tableName: "books",
		log:       logger,
	}
}

func selectQuery() string {
	return `
	id,
	title,
	author,
	publisher_date,
	isb
`
}

func returning(data string) string {
	return fmt.Sprintf("RETURNING  %s", data)
}

func (b *BookRepository) SaveBook(ctx context.Context, book *bookentity.Book) error {
	const op = "Repository.BookRepository.SaveBook"
	log := b.log.With(
		slog.String("operation", op))
	log.Info("starting")
	defer log.Info("finished")
	data := map[string]interface{}{
		"id":             book.Id,
		"title":          book.Title,
		"author":         book.Author,
		"publisher_date": book.PublishedDate,
		"isb":            book.Isbn,
	}

	query, args, err := b.db.Sq.Builder.Insert(b.tableName).
		SetMap(data).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	_, err = b.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return err
	}
	return nil
}

func (b *BookRepository) UpdateBook(ctx context.Context, book *bookentity.Book) error {
	const op = "Repository.BookRepository.UpdateBook"
	log := b.log.With(
		slog.String("operation", op),
		slog.String("id", book.Id))
	log.Info("starting")
	defer log.Info("finished")
	data := map[string]interface{}{
		"title":          book.Title,
		"author":         book.Author,
		"publisher_date": book.PublishedDate,
		"isb":            book.Isbn,
	}
	query, args, err := b.db.Sq.Builder.Update(b.tableName).
		Where(sq.Eq{"id": book.Id}).
		SetMap(data).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	_, err = b.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return err
	}
	return nil
}

func (b *BookRepository) DeleteBook(ctx context.Context, id string) error {
	const op = "Repository.BookRepository.DeleteBook"
	log := b.log.With(
		slog.String("operation", op))
	log.Info("starting")
	defer log.Info("finished")
	query, args, err := b.db.Sq.Builder.Delete(b.tableName).
		Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	_, err = b.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return err
	}
	return nil
}

func (b *BookRepository) GetBook(ctx context.Context, id string) (*bookentity.Book, error) {
	const op = "Repository.BookRepository.GetBook"
	log := b.log.With(
		slog.String("operation", op))
	log.Info("starting")
	defer log.Info("finished")
	var book bookentity.Book
	query, args, err := b.db.Sq.Builder.Select(selectQuery()).From(b.tableName).
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	err = b.db.QueryRow(ctx, query, args...).Scan(&book.Id,
		&book.Title,
		&book.Author,
		&book.PublishedDate,
		&book.Isbn)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &book, nil
}

func (b *BookRepository) GetAllBooks(ctx context.Context, req *bookentity.GetAllBooksRequest) ([]*bookentity.Book, error) {
	const op = "Repository.BookRepository.GetAllBooks"
	log := b.log.With(
		slog.String("operation", op))
	log.Info("starting")
	defer log.Info("finished")
	var books []*bookentity.Book

	toSql := b.db.Sq.Builder.Select(selectQuery()).From(b.tableName)

	if req.Field != "" && req.Value != "" {
		toSql = toSql.Where(sq.Eq{req.Field: req.Value})
	}
	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	rows, err := b.db.Query(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book bookentity.Book
		err = rows.Scan(&book.Id,
			&book.Title,
			&book.Author,
			&book.PublishedDate,
			&book.Isbn)
		if err != nil {
			log.Error("err", err)
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}
