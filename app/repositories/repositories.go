package repositories

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/internal/repository/admin_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/auth_admin"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/auth_user"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/book"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/book_favorite"
	"github.com/nktknshn/go-ergo-handler-example/internal/repository/user"
)

type Repositories struct {
	bookRepository         *book.BookRepository
	bookFavoriteRepository *book_favorite.BookFavoriteRepository
	authUserRepository     *auth_user.AuthUserRepository
	userRepository         *user.UserRepository
	adminUserRepository    *admin_user.AdminUserRepository
	authAdminRepository    *auth_admin.AuthAdminRepository
}

func New() *Repositories {
	return &Repositories{}
}

func (r *Repositories) Init(ctx context.Context) error {
	r.bookRepository = book.NewBookRepostiory()
	r.bookFavoriteRepository = book_favorite.NewBookFavoriteRepository()
	r.authUserRepository = auth_user.NewAuthUserRepository()
	r.userRepository = user.NewUserRepository()
	r.adminUserRepository = admin_user.NewAdminUserRepository()
	r.authAdminRepository = auth_admin.NewAuthAdminRepository()
	return nil
}

func (r *Repositories) GetBookRepository() *book.BookRepository {
	return r.bookRepository
}

func (r *Repositories) GetBookFavoriteRepository() *book_favorite.BookFavoriteRepository {
	return r.bookFavoriteRepository
}

func (r *Repositories) GetAuthUserRepository() *auth_user.AuthUserRepository {
	return r.authUserRepository
}

func (r *Repositories) GetUserRepository() *user.UserRepository {
	return r.userRepository
}

func (r *Repositories) GetAdminUserRepository() *admin_user.AdminUserRepository {
	return r.adminUserRepository
}

func (r *Repositories) GetAuthAdminRepository() *auth_admin.AuthAdminRepository {
	return r.authAdminRepository
}
