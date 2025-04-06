package main

import (
	"context"
	"log/slog"

	"github.com/nktknshn/go-ergo-handler-example/app"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/model/user"
)

func initState(ctx context.Context, app *app.App) error {
	slog.Info("initializing repositories state")

	repos := app.GetRepositories()
	repoBook := repos.GetBookRepository()
	repoUser := repos.GetUserRepository()
	repoAuthUser := repos.GetAuthUserRepository()
	repoAdmin := repos.GetAdminUserRepository()
	repoAuthAdmin := repos.GetAuthAdminRepository()

	for _, b := range dummyBooks {
		_, err := repoBook.CreateBook(ctx, b)
		if err != nil {
			slog.Error("failed to create book", "error", err)
			return err
		}
	}

	user, err := repoUser.CreateUser(ctx, user.User{
		Name: "John Doe",
	})
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return err
	}

	err = repoAuthUser.SetUserID(ctx, "user_token_12345", user.ID)
	if err != nil {
		slog.Error("failed to set user id", "error", err)
		return err
	}

	admin, err := repoAdmin.CreateAdmin(ctx, admin_user.AdminUser{
		Username: "admin",
		Role:     admin_user.AdminUserRoleAdmin,
	})
	if err != nil {
		slog.Error("failed to create admin user", "error", err)
		return err
	}

	err = repoAuthAdmin.SetAdminID(ctx, "admin_token_12345", admin.ID)
	if err != nil {
		slog.Error("failed to set admin user id", "error", err)
		return err
	}

	// moderator
	moderator, err := repoAdmin.CreateAdmin(ctx, admin_user.AdminUser{
		Username: "moderator",
		Role:     admin_user.AdminUserRoleModerator,
	})
	if err != nil {
		slog.Error("failed to create moderator user", "error", err)
		return err
	}

	err = repoAuthAdmin.SetAdminID(ctx, "moderator_token_12345", moderator.ID)
	if err != nil {
		slog.Error("failed to set moderator user id", "error", err)
		return err
	}

	return nil
}
