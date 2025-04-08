package book

import "errors"

var (
	ErrInvalidBookID          = errors.New("invalid book id")
	ErrInvalidBookTitle       = errors.New("invalid book title")
	ErrInvalidBookAuthor      = errors.New("invalid book author")
	ErrInvalidBookDescription = errors.New("invalid book description")
)

type BookID int

func NewBook(title BookTitle, author BookAuthor, description BookDescription) (Book, error) {
	if err := title.Validate(); err != nil {
		return Book{}, err
	}
	if err := author.Validate(); err != nil {
		return Book{}, err
	}
	if err := description.Validate(); err != nil {
		return Book{}, err
	}
	return Book{Title: title, Author: author, Description: description}, nil
}

func (b BookID) Validate() error {
	if b <= 0 {
		return ErrInvalidBookID
	}
	return nil
}

type BookTitle string

func (t BookTitle) Validate() error {
	if t == "" {
		return ErrInvalidBookTitle
	}
	return nil
}

type BookAuthor string

func (a BookAuthor) Validate() error {
	if a == "" {
		return ErrInvalidBookAuthor
	}
	return nil
}

type BookDescription string

func (d BookDescription) Validate() error {
	if d == "" {
		return ErrInvalidBookDescription
	}
	return nil
}

type Book struct {
	ID          BookID
	Title       BookTitle
	Author      BookAuthor
	Description BookDescription
}

func (b Book) HasID() bool {
	return b.ID != 0
}

func (b Book) IsZero() bool {
	return b == Book{}
}
