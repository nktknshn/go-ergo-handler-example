package books

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
)

func (u *BooksUseCase) CreateBook(ctx context.Context, publisherID admin_user.AdminUserID, bookToCreate book.Book) (book.Book, error) {
	newBook, err := u.booksRepository.CreateBook(ctx, bookToCreate)
	if err != nil {
		return book.Book{}, err
	}
	return newBook, nil
}
