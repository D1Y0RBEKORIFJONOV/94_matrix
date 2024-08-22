package handler

import (
	_ "94.Metrics/internal/app/docs"
	bookentity "94.Metrics/internal/entity/book"
	bookusecase "94.Metrics/internal/usecase/book"
	"expvar"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	book           *bookusecase.BookUseCase
	requestCount   *expvar.Int
	successCount   *expvar.Int
	errorCount     *expvar.Int
	dbRequestCount *expvar.Int
}

func NewBookHandler(book *bookusecase.BookUseCase) *Book {
	return &Book{
		book:           book,
		requestCount:   expvar.NewInt("requestCount"),
		successCount:   expvar.NewInt("successCount"),
		errorCount:     expvar.NewInt("errorCount"),
		dbRequestCount: expvar.NewInt("dbRequestCount"),
	}
}

// CreateBook godoc
// @Summary CreateBook
// @Description CreateBook a new book
// @Tags books
// @Accept json
// @Produce json
// @Param body body bookentity.CreateBookRequest true "creating new book"
// @Success 201 {object} bookentity.Book
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /book/create [post]
func (b *Book) CreateBook(c *gin.Context) {
	b.requestCount.Add(1)
	var req *bookentity.CreateBookRequest
	if err := c.ShouldBind(&req); err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b.dbRequestCount.Add(2)
	book, err := b.book.CreateBook(c.Request.Context(), req)
	if err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	b.successCount.Add(1)
	c.JSON(http.StatusCreated, book)
}

// GetBook godoc
// @Summary GetBook
// @Description GetBook
// @Tags books
// @Accept json
// @Produce json
// @Param id query string false "id book"
// @Success 200 {object} bookentity.Book
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /book/ [get]
func (b *Book) GetBook(c *gin.Context) {
	b.requestCount.Add(1)
	b.dbRequestCount.Add(1)
	id := c.DefaultQuery("id", "")
	if id == "" {
		b.errorCount.Add(1)
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &bookentity.GetBookRequest{Id: id}

	book, err := b.book.GetBook(c.Request.Context(), req)
	if err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	b.successCount.Add(1)
	c.JSON(http.StatusOK, book)
}

// GetBooks godoc
// @Summary GetBooks
// @Description GetBooks
// @Tags books
// @Accept json
// @Produce json
// @Param value query string false "value book"
// @Param field query string false "field book"
// @Success 201 {object} []bookentity.Book
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /book/books [get]
func (b *Book) GetBooks(c *gin.Context) {
	b.requestCount.Add(1)
	var req bookentity.GetAllBooksRequest

	req.Value = c.DefaultQuery("value", "")
	req.Field = c.DefaultQuery("field", "")

	if err := c.ShouldBind(&req); err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b.dbRequestCount.Add(1)
	books, err := b.book.GetAllBook(c.Request.Context(), &req)
	if err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	b.successCount.Add(1)
	c.JSON(http.StatusOK, books)
}

// DeleteBook godoc
// @Summary DeleteBook
// @Description DeleteBook
// @Tags books
// @Accept json
// @Produce json
// @Param id query string false "id book"
// @Success 201 {object} string  "message status"
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /book/delete [delete]
func (b *Book) DeleteBook(c *gin.Context) {
	b.requestCount.Add(1)
	id := c.DefaultQuery("id", "")
	if id == "" {
		b.errorCount.Add(1)
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &bookentity.DeleteBookRequest{Id: id}
	if err := c.ShouldBind(&req); err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b.dbRequestCount.Add(1)
	err := b.book.DeleteBook(c.Request.Context(), req)
	if err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	b.successCount.Add(1)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// UpdateBook godoc
// @Summary UpdateBook
// @Description UpdateBook a  book
// @Tags books
// @Accept json
// @Produce json
// @Param body body bookentity.UpdateBookRequest true "creating new book"
// @Param id query string false "id book"
// @Success 201 {object} bookentity.Book
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /book/update [put]
func (b *Book) UpdateBook(c *gin.Context) {
	b.requestCount.Add(1)
	id := c.DefaultQuery("id", "")
	if id == "" {
		b.errorCount.Add(1)
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &bookentity.UpdateBookRequest{Id: id}

	if err := c.ShouldBind(&req); err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b.dbRequestCount.Add(1)
	book, err := b.book.UpdateBook(c.Request.Context(), req)
	if err != nil {
		b.errorCount.Add(1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	b.successCount.Add(1)
	c.JSON(http.StatusOK, book)
}

// Metrics godoc
// @Summary Get application metrics
// @Description Returns application metrics in expvar format
// @Tags metrics
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /metrics [get]
func (b *Book) Metrics(c *gin.Context) {
	expvarHandler := expvar.Handler()
	expvarHandler.ServeHTTP(c.Writer, c.Request)
}
