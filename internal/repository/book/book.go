package book

import (
	"context"
	"sync"

	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
)

type BookRepository struct {
	lock       *sync.RWMutex
	lastBookID book.BookID
	// books are sorted by ID ascending
	books    []book.Book
	bookByID map[book.BookID]*book.Book
}

func NewBookRepostiory() *BookRepository {
	return &BookRepository{
		lock:       &sync.RWMutex{},
		lastBookID: 0,
		books:      []book.Book{},
		bookByID:   map[book.BookID]*book.Book{},
	}
}

func (r *BookRepository) GetBookByID(_ context.Context, bookID book.BookID) (book.Book, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.getBookByID(bookID)
}

func (r *BookRepository) getBookByID(bookID book.BookID) (book.Book, error) {
	b, ok := r.bookByID[bookID]
	if !ok {
		return book.Book{}, bookRepoValObj.ErrBookNotFound
	}
	return *b, nil
}

func (r *BookRepository) GetBooksList(ctx context.Context, cursor *bookRepoValObj.BookListCursor, limit bookRepoValObj.BookListLimit) (bookRepoValObj.BookList, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.getBooksList(cursor, limit)
}

func (r *BookRepository) getBooksList(cursor *bookRepoValObj.BookListCursor, limit bookRepoValObj.BookListLimit) (bookRepoValObj.BookList, error) {
	var err error
	var result bookRepoValObj.BookList

	low := 0
	if cursor != nil {
		low, err = parseBookListCursor(*cursor)
		if err != nil {
			return result, err
		}
	}
	if low >= len(r.books) {
		return result, nil
	}
	high := low + limit.Int()
	high = min(high, len(r.books))
	page := r.books[low:high]
	cp := make([]book.Book, len(page))
	copy(cp, page)
	result.Books = cp
	result.HasMore = high < len(r.books)
	result.Cursor = cursorFromInt(high)
	return result, nil
}

// UpsertBook upserts a book into the repository.
func (r *BookRepository) UpsertBook(ctx context.Context, b book.Book) (book.Book, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if b.HasID() {
		return r.updateBook(ctx, b)
	}
	return r.createBook(ctx, b)
}

// CreateBook creates a new book in the repository.
func (r *BookRepository) CreateBook(ctx context.Context, b book.Book) (book.Book, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.createBook(ctx, b)
}

func (r *BookRepository) updateBook(_ context.Context, b book.Book) (book.Book, error) {
	existingBook, ok := r.bookByID[b.ID]
	if !ok {
		return book.Book{}, bookRepoValObj.ErrBookNotFound
	}
	existingBook.Title = b.Title
	existingBook.Author = b.Author
	existingBook.Description = b.Description
	r.bookByID[b.ID] = existingBook
	return *existingBook, nil
}

func (r *BookRepository) createBook(_ context.Context, newBook book.Book) (book.Book, error) {
	newBook.ID = r.makeNewBookID()
	r.books = append(r.books, newBook)
	r.bookByID[newBook.ID] = &newBook
	return newBook, nil
}

func (r *BookRepository) makeNewBookID() book.BookID {
	r.lastBookID++
	return r.lastBookID
}
