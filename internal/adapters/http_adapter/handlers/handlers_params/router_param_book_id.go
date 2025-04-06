package handlers_params

import (
	"context"
	"errors"
	"strconv"

	geh "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/book"
)

var (
	paramNameBookID            = "book_id"
	errParamBookIDEmpty        = geh.NewErrorStr(400, "param "+paramNameBookID+" is empty")
	errParamBookIDInvalid      = geh.NewErrorStr(400, "param "+paramNameBookID+" is invalid")
	errParamBookIDInvalidValue = geh.NewErrorStr(400, "param "+paramNameBookID+" has invalid value")
	errParamBookIDMissing      = geh.NewErrorStr(400, "param "+paramNameBookID+" is missing")
)

type paramBookIDType int

func (p paramBookIDType) Parse(ctx context.Context, v string) (paramBookIDType, error) {
	if v == "" {
		return 0, errParamBookIDEmpty
	}
	vint, err := strconv.Atoi(v)
	if err != nil {
		return 0, errParamBookIDInvalid
	}
	return paramBookIDType(vint), nil
}

func (p paramBookIDType) Validate() error {
	err := book.BookID(p).Validate()
	if err != nil {
		return errors.Join(errParamBookIDInvalidValue, err)
	}
	return nil
}

func (p paramBookIDType) ToBookID() book.BookID {
	return book.BookID(p)
}

var RouterParamBookID = geh.RouterParamWithParser[paramBookIDType](paramNameBookID, errParamBookIDMissing)
