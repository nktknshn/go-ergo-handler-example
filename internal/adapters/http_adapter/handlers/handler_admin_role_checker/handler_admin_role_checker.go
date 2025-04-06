package handler_admin_role_checker

import (
	"context"
	"errors"
	"net/http"

	"slices"

	geh "github.com/nktknshn/go-ergo-handler"
	"github.com/nktknshn/go-ergo-handler-example/internal/adapters/http_adapter/handlers/handler_admin_auth"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
)

type AdminRolerChecker struct {
	cf CheckFunc
}

var (
	ErrUserNotAllowed = errors.New("user is not allowed")
)

type CheckFunc = func(ctx context.Context, user *admin_user.AdminUser) error

func CheckSlice(roles []admin_user.AdminUserRole) CheckFunc {
	return func(ctx context.Context, user *admin_user.AdminUser) error {
		if slices.Contains(roles, user.Role) {
			return nil
		}
		return geh.NewError(http.StatusForbidden, ErrUserNotAllowed)
	}
}

func NewAdminRoleChecker(cf CheckFunc) *AdminRolerChecker {
	return &AdminRolerChecker{cf}
}

func (a *AdminRolerChecker) Attach(auth *handler_admin_auth.AttachedAdminUserParser, builder geh.HandlerBuilder) *attachedAdminRolerChecker {
	attached := &attachedAdminRolerChecker{a.cf, auth}
	builder.AddParser(attached)
	return attached
}

type attachedAdminRolerChecker struct {
	cf   CheckFunc
	auth *handler_admin_auth.AttachedAdminUserParser
}

func (at *attachedAdminRolerChecker) ParseRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, error) {
	user := at.auth.GetUser(ctx)
	err := at.cf(ctx, user)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
