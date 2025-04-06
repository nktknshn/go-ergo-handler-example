package books

import (
	"context"
	"errors"

	bookModel "github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
)

func (u *BooksUseCase) GetBookByID(ctx context.Context, bookID bookModel.BookID) (bookModel.Book, error) {
	book, err := u.booksRepository.GetBookByID(ctx, bookID)

	if errors.Is(err, bookRepoValObj.ErrBookNotFound) {
		return bookModel.Book{}, useCaseValObj.ErrBookNotFound
	}

	if err != nil {
		return bookModel.Book{}, err
	}

	return book, nil
}
