package handlers_user_auth

import (
	"context"
	"net/http"
	"strings"

	geh "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
)

type userKeyType string

const userKey userKeyType = "user"

var tokenParserFunc = func(ctx context.Context, r *http.Request) (string, bool, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", false, nil
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return "", false, nil
	}
	return token, true, nil
}

var (
	UserParser      = geh.AuthParser[user.User](userKey, tokenParserFunc)
	UserParserMaybe = geh.AuthParserMaybe[user.User](userKey, tokenParserFunc)
)
