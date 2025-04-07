package books

import (
	"context"
	"log/slog"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
)

func (u *BooksUseCase) CreateBook(ctx context.Context, publisherID admin_user.AdminUserID, bookToCreate book.Book) (book.Book, error) {
	newBook, err := u.booksRepository.CreateBook(ctx, bookToCreate)
	if err != nil {
		slog.Error("u.booksRepository.CreateBook", "error", err)
		return book.Book{}, useCaseValObj.ErrCreateBookFailed
	}
	return newBook, nil
}
