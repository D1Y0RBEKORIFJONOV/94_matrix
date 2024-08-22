package router

import (
	"94.Metrics/internal/http/handler"
	bookusecase "94.Metrics/internal/usecase/book"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(book *bookusecase.BookUseCase) *gin.Engine {
	bookHandler := handler.NewBookHandler(book)
	router := gin.New()

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	bookGroup := router.Group("/book")
	{
		bookGroup.POST("/create", bookHandler.CreateBook)
		bookGroup.PUT("/update", bookHandler.UpdateBook)
		bookGroup.DELETE("/delete", bookHandler.DeleteBook)
		bookGroup.GET("", bookHandler.GetBook)
		bookGroup.GET("/books", bookHandler.GetBooks)
	}

	router.GET("/metrics", bookHandler.Metrics)

	return router
}
