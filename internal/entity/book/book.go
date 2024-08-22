package bookentity

import "time"

type (
	Book struct {
		Id            string    `json:"id" bson:"_id,omitempty"`
		Title         string    `json:"title" bson:"title,omitempty"`
		Author        string    `json:"author" bson:"author,omitempty"`
		PublishedDate time.Time `json:"published_date" bson:"published_date,omitempty"`
		Isbn          string    `json:"isbn" bson:"isbn,omitempty"`
	}
	CreateBookRequest struct {
		Title  string `json:"title" bson:"title,omitempty"`
		Author string `json:"author" bson:"author,omitempty"`
		Isbn   string `json:"isbn" bson:"isbn,omitempty"`
	}
	GetBookRequest struct {
		Id string `json:"-" bson:"_id,omitempty"`
	}
	GetAllBooksRequest struct {
		Field string `json:"field" bson:"field,omitempty"`
		Value string `json:"value" bson:"value,omitempty"`
	}
	UpdateBookRequest struct {
		Id     string `json:"-" bson:"_id,omitempty"`
		Title  string `json:"title" bson:"title,omitempty"`
		Author string `json:"author" bson:"author,omitempty"`
		Isbn   string `json:"isbn" bson:"isbn,omitempty"`
	}
	DeleteBookRequest struct {
		Id string `json:"-" bson:"_id,omitempty"`
	}
)
