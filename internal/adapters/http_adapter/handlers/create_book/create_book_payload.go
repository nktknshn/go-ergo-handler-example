package create_book

import (
	"errors"

	builder "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
)

var (
	ErrTitleRequired       = errors.New("title is required")
	ErrAuthorRequired      = errors.New("author is required")
	ErrDescriptionRequired = errors.New("description is required")
)

type payload struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (p *payload) Validate() error {
	if p.Title == "" {
		return ErrTitleRequired
	}
	if p.Author == "" {
		return ErrAuthorRequired
	}
	if p.Description == "" {
		return ErrDescriptionRequired
	}
	return nil
}

func (p *payload) ToBook() book.Book {
	return book.Book{
		Title:       book.BookTitle(p.Title),
		Author:      book.BookAuthor(p.Author),
		Description: book.BookDescription(p.Description),
	}
}

var payloadCreateBook = builder.NewPayloadWithValidation(&payload{})
