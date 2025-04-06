package books_user_aware

import (
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
)

type GetBooksListResponse struct {
	books.GetBooksListResponse
	FavoriteBooks []book.BookID `json:"favorite_books"`
}
