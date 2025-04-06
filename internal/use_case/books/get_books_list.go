package books

import (
	"context"

	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
	useCaseValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/use_case/books"
)

var (
	defaultBookListLimit = bookRepoValObj.BookListLimit(10)
)

func (u *BooksUseCase) GetBooksList(ctx context.Context, query useCaseValObj.GetBooksListQuery) (useCaseValObj.GetBooksListResponse, error) {
	response := useCaseValObj.GetBooksListResponse{}
	bookList, err := u.booksRepository.GetBooksList(ctx, query.Cursor, defaultBookListLimit)
	if err != nil {
		return useCaseValObj.GetBooksListResponse{}, err
	}
	response.Books = make([]useCaseValObj.Book, len(bookList.Books))
	for i, book := range bookList.Books {
		response.Books[i] = useCaseValObj.Book{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
		}
	}
	response.Cursor = bookList.Cursor
	response.HasMore = bookList.HasMore
	return response, nil
}
