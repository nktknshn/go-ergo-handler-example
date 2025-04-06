package handlers_params

import (
	geh "github.com/nktknshn/go-ergo-handler"
)

var QueryParamBookIDMaybe = geh.QueryParamWithParserMaybe[paramBookIDType](paramNameBookID, errParamBookIDMissing)

var QueryParamBookID = geh.QueryParamWithParser[paramBookIDType](paramNameBookID, errParamBookIDMissing)
