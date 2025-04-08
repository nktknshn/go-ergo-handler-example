package handler_admin_auth

import (
	"context"
	"net/http"
	"strings"

	geh "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
)

type adminUserKeyType string

const adminUserKey adminUserKeyType = "admin_user"

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

var AdminUserParser = geh.AuthParser[admin_user.AdminUser](adminUserKey, tokenParserFunc)

type AttachedAdminUserParser = geh.AttachedAuthParser[admin_user.AdminUser, adminUserKeyType]
