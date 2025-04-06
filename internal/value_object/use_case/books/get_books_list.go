package books

import (
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
)

type GetBooksListQuery struct {
	Cursor *bookRepoValObj.BookListCursor `json:"cursor"`
}

type GetBooksListResponse struct {
	Books   []Book                        `json:"books"`
	Cursor  bookRepoValObj.BookListCursor `json:"cursor"`
	HasMore bool                          `json:"has_more"`
}

type Book struct {
	ID          book.BookID          `json:"id"`
	Title       book.BookTitle       `json:"title"`
	Description book.BookDescription `json:"description"`
	Author      book.BookAuthor      `json:"author"`
}
