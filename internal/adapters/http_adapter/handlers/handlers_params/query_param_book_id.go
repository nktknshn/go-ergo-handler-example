package handlers_params

import (
	geh "github.com/nktknshn/go-ergo-handler"
)

var (
	QueryParamBookIDMaybe = geh.QueryParamWithParserMaybe[paramBookIDType](paramNameBookID)
	QueryParamBookID      = geh.QueryParamWithParser[paramBookIDType](paramNameBookID, errParamBookIDMissing)
)
