package book

type BookID int
type BookTitle string
type BookAuthor string
type BookDescription string

type Book struct {
	ID          BookID
	Title       BookTitle
	Author      BookAuthor
	Description BookDescription
}

func (b *Book) HasID() bool {
	return b.ID != 0
}

func (b *Book) IsZero() bool {
	return b.ID == 0 && b.Title == "" && b.Author == "" && b.Description == ""
}
