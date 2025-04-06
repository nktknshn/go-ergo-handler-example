package book_favorite

import (
	"context"
	"sync"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book_favorite"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	"github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books_favorite"
	"github.com/nktknshn/go-ergo-handler-example/libs/set"
)

type BookFavoriteRepository struct {
	lock           *sync.RWMutex
	booksByUserID  map[user.UserID]*userFavoriteBooks
	lastFavoriteID book_favorite.BookFavoriteID
}

func NewBookFavoriteRepository() *BookFavoriteRepository {
	return &BookFavoriteRepository{
		lock:           &sync.RWMutex{},
		lastFavoriteID: 0,
		booksByUserID:  make(map[user.UserID]*userFavoriteBooks),
	}
}

func (r *BookFavoriteRepository) GetFavoriteBooks(ctx context.Context, userID user.UserID) ([]book.BookID, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.getFavoriteBooks(ctx, userID)
}

func (r *BookFavoriteRepository) getFavoriteBooks(_ context.Context, userID user.UserID) ([]book.BookID, error) {
	books := r.booksByUserID[userID]
	if books == nil {
		return nil, nil
	}
	return books.GetBooksIDs(), nil
}

func (r *BookFavoriteRepository) AddFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.addFavoriteBook(ctx, userID, bookID)
}

func (r *BookFavoriteRepository) addFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) (book_favorite.BookFavorite, error) {
	userBooks := r.booksByUserID[userID]
	if userBooks == nil {
		userBooks = &userFavoriteBooks{
			books:    []book_favorite.BookFavorite{},
			booksSet: set.New[book.BookID](),
		}
		r.booksByUserID[userID] = userBooks
	}
	if userBooks.HasBook(bookID) {
		return book_favorite.BookFavorite{}, books_favorite.ErrBookAlreadyInFavorites
	}
	favoriteID := r.makeNewFavoriteID()
	favorite := book_favorite.BookFavorite{
		ID:     favoriteID,
		BookID: bookID,
		UserID: userID,
	}
	userBooks.AddBook(favorite)
	return favorite, nil
}

func (r *BookFavoriteRepository) RemoveFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.removeFavoriteBook(ctx, userID, bookID)
}

func (r *BookFavoriteRepository) removeFavoriteBook(ctx context.Context, userID user.UserID, bookID book.BookID) error {
	userBooks := r.booksByUserID[userID]
	if userBooks == nil {
		return books_favorite.ErrBookNotFoundInFavorites
	}
	if !userBooks.HasBook(bookID) {
		return books_favorite.ErrBookNotFoundInFavorites
	}
	userBooks.RemoveBook(bookID)
	return nil
}

func (r *BookFavoriteRepository) makeNewFavoriteID() book_favorite.BookFavoriteID {
	r.lastFavoriteID++
	return r.lastFavoriteID
}
