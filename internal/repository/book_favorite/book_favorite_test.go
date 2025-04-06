package book_favorite

import (
	"context"
	"testing"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	"github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books_favorite"
	"github.com/stretchr/testify/require"
)

var (
	userID1 = user.UserID(1)
	// userID2 = user.UserID(2)
	bookID1 = book.BookID(1)
	bookID2 = book.BookID(2)
)

func TestBookFavoriteRepository_AddFavoriteBook(t *testing.T) {
	repo := NewBookFavoriteRepository()
	favorite, err := repo.AddFavoriteBook(context.Background(), userID1, bookID1)
	require.NoError(t, err)
	require.Equal(t, favorite.UserID, userID1)
	require.Equal(t, favorite.BookID, bookID1)
}

func TestBookFavoriteRepository_AddFavoriteBook_AlreadyExists(t *testing.T) {
	repo := NewBookFavoriteRepository()
	favorite, err := repo.AddFavoriteBook(context.Background(), userID1, bookID1)
	require.NoError(t, err)
	require.Equal(t, favorite.UserID, userID1)
	require.Equal(t, favorite.BookID, bookID1)

	favorite, err = repo.AddFavoriteBook(context.Background(), userID1, bookID1)
	require.Error(t, err)
	require.ErrorIs(t, err, books_favorite.ErrBookAlreadyInFavorites)
}

func TestBookFavoriteRepository_GetFavoriteBooks(t *testing.T) {
	repo := NewBookFavoriteRepository()
	_, err := repo.AddFavoriteBook(context.Background(), userID1, bookID1)
	require.NoError(t, err)
	_, err = repo.AddFavoriteBook(context.Background(), userID1, bookID2)
	require.NoError(t, err)

	favorites, err := repo.GetFavoriteBooks(context.Background(), userID1)
	require.NoError(t, err)
	require.Equal(t, len(favorites), 2)
	require.Contains(t, favorites, bookID1)
	require.Contains(t, favorites, bookID2)
}

func TestBookFavoriteRepository_RemoveFavoriteBook(t *testing.T) {
	repo := NewBookFavoriteRepository()
	_, err := repo.AddFavoriteBook(context.Background(), userID1, bookID1)
	require.NoError(t, err)

	err = repo.RemoveFavoriteBook(context.Background(), userID1, bookID1)
	require.NoError(t, err)

	favorites, err := repo.GetFavoriteBooks(context.Background(), userID1)
	require.NoError(t, err)
	require.Equal(t, len(favorites), 0)
}
