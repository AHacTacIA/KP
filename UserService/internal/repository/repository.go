// Package repository a
package repository

import (
	"github.com/AHacTacIA/KP/UserService/internal/user"

	"context"
)

// Repository transition to mongo or postgres db
type Repository interface {
	CreateUser(ctx context.Context, p *user.Person) (string, error)
	GetUserByID(ctx context.Context, idPerson string) (*user.Person, error)
	GetAllUsers(ctx context.Context) ([]*user.Person, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string, per *user.Person) error
	SelectByIDAuth(ctx context.Context, id string) (user.Person, error)
	UpdateAuth(ctx context.Context, id string, refreshToken string) error
}
