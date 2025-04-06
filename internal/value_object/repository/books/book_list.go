package books

import "github.com/nktknshn/go-ergo-handler-example/internal/model/book"

type BookList struct {
	Books   []book.Book
	Cursor  BookListCursor
	HasMore bool
}
