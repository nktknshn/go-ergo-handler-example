package books

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
)

type BooksUseCase struct {
	booksRepository booksRepository
}

type booksRepository interface {
	GetBooksList(ctx context.Context, cursor *bookRepoValObj.BookListCursor, limit bookRepoValObj.BookListLimit) (bookRepoValObj.BookList, error)
	GetBookByID(ctx context.Context, bookID book.BookID) (book.Book, error)
	CreateBook(ctx context.Context, book book.Book) (book.Book, error)
}

func NewBooksUseCase(booksRepository booksRepository) *BooksUseCase {
	return &BooksUseCase{
		booksRepository,
	}
}
