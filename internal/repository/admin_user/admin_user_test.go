package admin_user

import (
	"context"
	"testing"

	adminUserModel "github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	"github.com/stretchr/testify/require"
)

func TestAdminUserRepository_UpsertAdmin(t *testing.T) {
	repo := NewAdminUserRepository()
	admin, err := repo.UpsertAdmin(context.Background(), adminUserModel.AdminUser{
		Username: "admin",
		Role:     adminUserModel.AdminUserRoleAdmin,
	})

	require.NoError(t, err)
	require.Greater(t, admin.ID.Int(), 0)

	adminChanged, err := repo.UpsertAdmin(context.Background(), adminUserModel.AdminUser{
		ID:       admin.ID,
		Username: "admin2",
		Role:     adminUserModel.AdminUserRoleAdmin,
	})

	require.NoError(t, err)
	require.Equal(t, admin.ID, adminChanged.ID)
	require.Equal(t, "admin2", adminChanged.Username.String())
}
