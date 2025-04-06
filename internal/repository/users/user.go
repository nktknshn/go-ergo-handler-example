package users

import (
	"context"
	"sync"

	userModel "github.com/nktknshn/go-ergo-handler-example/internal/model/user"
	userValObj "github.com/nktknshn/go-ergo-handler-example/internal/value_object/repository/users"
)

type UserRepository struct {
	lock       *sync.RWMutex
	lastUserID userModel.UserID
	users      map[userModel.UserID]userModel.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		lock:       &sync.RWMutex{},
		lastUserID: 0,
		users:      make(map[userModel.UserID]userModel.User),
	}
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID userModel.UserID) (userModel.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.getUserByID(ctx, userID)
}

func (r *UserRepository) getUserByID(ctx context.Context, userID userModel.UserID) (userModel.User, error) {
	user, ok := r.users[userID]
	if !ok {
		return userModel.User{}, userValObj.ErrUserNotFound
	}
	return user, nil
}

func (r *UserRepository) UpsertUser(ctx context.Context, user userModel.User) (userModel.User, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if user.HasID() {
		return r.updateUser(ctx, user)
	}
	return r.createUser(ctx, user)
}

// CreateUser
func (r *UserRepository) CreateUser(ctx context.Context, user userModel.User) (userModel.User, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.createUser(ctx, user)
}

func (r *UserRepository) updateUser(_ context.Context, user userModel.User) (userModel.User, error) {
	r.users[user.ID] = user
	return user, nil
}

func (r *UserRepository) createUser(_ context.Context, user userModel.User) (userModel.User, error) {
	user.ID = r.makeNewUserID()
	r.users[user.ID] = user
	return user, nil
}

func (r *UserRepository) makeNewUserID() userModel.UserID {
	r.lastUserID++
	return r.lastUserID
}
