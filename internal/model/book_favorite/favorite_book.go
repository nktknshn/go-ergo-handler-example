package book_favorite

import (
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
)

type BookFavoriteID int

type BookFavorite struct {
	ID     BookFavoriteID
	UserID user.UserID
	BookID book.BookID
}
