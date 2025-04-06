package books

import "errors"

type BookListLimit int

var (
	ErrInvalidLimit = errors.New("invalid limit value")
)

func (bl BookListLimit) Int() int {
	return int(bl)
}

func (bl BookListLimit) Validate() error {
	if bl < 1 {
		return ErrInvalidLimit
	}
	return nil
}
