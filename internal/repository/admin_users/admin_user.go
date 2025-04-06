package admin_users

import (
	"context"
	"sync"

	adminUserModel "github.com/nktknshn/go-ergo-handler-example/internal/model/admin_user"
	adminUserValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/admin_users"
)

type AdminUserRepository struct {
	lock            *sync.RWMutex
	lastAdminUserID adminUserModel.AdminUserID
	admins          map[adminUserModel.AdminUserID]adminUserModel.AdminUser
}

func NewAdminUserRepository() *AdminUserRepository {
	return &AdminUserRepository{
		lock:            &sync.RWMutex{},
		lastAdminUserID: 0,
		admins:          make(map[adminUserModel.AdminUserID]adminUserModel.AdminUser),
	}
}

func (r *AdminUserRepository) GetAdminByID(ctx context.Context, adminID adminUserModel.AdminUserID) (adminUserModel.AdminUser, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.getAdminByID(ctx, adminID)
}

func (r *AdminUserRepository) getAdminByID(_ context.Context, adminID adminUserModel.AdminUserID) (adminUserModel.AdminUser, error) {
	admin, ok := r.admins[adminID]
	if !ok {
		return adminUserModel.AdminUser{}, adminUserValObj.ErrAdminNotFound
	}
	return admin, nil
}

func (r *AdminUserRepository) UpsertAdmin(ctx context.Context, admin adminUserModel.AdminUser) (adminUserModel.AdminUser, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if admin.HasID() {
		return r.updateAdmin(ctx, admin)
	}
	return r.createAdmin(ctx, admin)
}

// CreateAdmin
func (r *AdminUserRepository) CreateAdmin(ctx context.Context, admin adminUserModel.AdminUser) (adminUserModel.AdminUser, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.createAdmin(ctx, admin)
}

func (r *AdminUserRepository) updateAdmin(_ context.Context, admin adminUserModel.AdminUser) (adminUserModel.AdminUser, error) {
	r.admins[admin.ID] = admin
	return admin, nil
}

func (r *AdminUserRepository) createAdmin(_ context.Context, admin adminUserModel.AdminUser) (adminUserModel.AdminUser, error) {
	admin.ID = r.makeNewAdminID()
	r.admins[admin.ID] = admin
	return admin, nil
}

func (r *AdminUserRepository) makeNewAdminID() adminUserModel.AdminUserID {
	r.lastAdminUserID++
	return r.lastAdminUserID
}
