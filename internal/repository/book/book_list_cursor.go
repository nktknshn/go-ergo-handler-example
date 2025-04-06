package book

import (
	"fmt"
	"strconv"

	bookRepoValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/books"
)

func parseBookListCursor(cursor bookRepoValObj.BookListCursor) (int, error) {
	if cursor.IsEmpty() {
		return 0, nil
	}

	vint, err := strconv.Atoi(cursor.String())

	if err != nil {
		return 0, bookRepoValObj.ErrBookListCursorInvalid
	}

	return vint, nil
}

func cursorFromInt(vint int) bookRepoValObj.BookListCursor {
	return bookRepoValObj.BookListCursor(fmt.Sprintf("%d", vint))
}
