package repositories

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/repository/admin_users"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/auth_admin"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/auth_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/book_favorites"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/books"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/users"
)

type Repositories struct {
	bookRepository         *books.BookRepository
	bookFavoriteRepository *book_favorites.BookFavoriteRepository
	authUserRepository     *auth_user.AuthUserRepository
	userRepository         *users.UserRepository
	adminUserRepository    *admin_users.AdminUserRepository
	authAdminRepository    *auth_admin.AuthAdminRepository
}

func New() *Repositories {
	return &Repositories{}
}

func (r *Repositories) Init(ctx context.Context) error {
	r.bookRepository = books.NewBookRepostiory()
	r.bookFavoriteRepository = book_favorites.NewBookFavoriteRepository()
	r.authUserRepository = auth_user.NewAuthUserRepository()
	r.userRepository = users.NewUserRepository()
	r.adminUserRepository = admin_users.NewAdminUserRepository()
	r.authAdminRepository = auth_admin.NewAuthAdminRepository()
	return nil
}

func (r *Repositories) GetBookRepository() *books.BookRepository {
	return r.bookRepository
}

func (r *Repositories) GetBookFavoriteRepository() *book_favorites.BookFavoriteRepository {
	return r.bookFavoriteRepository
}

func (r *Repositories) GetAuthUserRepository() *auth_user.AuthUserRepository {
	return r.authUserRepository
}

func (r *Repositories) GetUserRepository() *users.UserRepository {
	return r.userRepository
}

func (r *Repositories) GetAdminUserRepository() *admin_users.AdminUserRepository {
	return r.adminUserRepository
}

func (r *Repositories) GetAuthAdminRepository() *auth_admin.AuthAdminRepository {
	return r.authAdminRepository
}
