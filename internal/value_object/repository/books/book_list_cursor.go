package books

type BookListCursor string

func (bc BookListCursor) String() string {
	return string(bc)
}

func (bc BookListCursor) IsEmpty() bool {
	return bc == ""
}
