package book_favorites

import (
	"slices"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book_favorite"
	"github.com/nktknshn/go-ergo-handler-example/libs/set"
)

type userFavoriteBooks struct {
	books    []book_favorite.BookFavorite
	booksSet *set.Set[book.BookID]
}

func (u *userFavoriteBooks) AddBook(book book_favorite.BookFavorite) {
	if u.booksSet.Has(book.BookID) {
		return
	}
	u.books = append(u.books, book)
	u.booksSet.Add(book.BookID)
}

func (u *userFavoriteBooks) RemoveBook(bookID book.BookID) {
	if !u.booksSet.Has(bookID) {
		return
	}
	u.booksSet.Remove(bookID)
	u.books = slices.DeleteFunc(u.books, func(book book_favorite.BookFavorite) bool {
		return book.BookID == bookID
	})
}

func (u *userFavoriteBooks) GetBooks() []book_favorite.BookFavorite {
	return u.books
}

func (u *userFavoriteBooks) GetBooksIDs() []book.BookID {
	favorites := u.GetBooks()
	books := make([]book.BookID, len(favorites))
	for i, favorite := range favorites {
		books[i] = favorite.BookID
	}
	return books
}

func (u *userFavoriteBooks) HasBook(bookID book.BookID) bool {
	return u.booksSet.Has(bookID)
}
