// Package service i
package service

import (
	"context"
	"github.com/AHacTacIA/KP/UserService/internal/repository"
	"github.com/AHacTacIA/KP/UserService/internal/user"
)

// Service s
type Service struct {
	jwtKey []byte
	rps    repository.Repository
}

// NewService create new service connection
func NewService(pool repository.Repository, jwtKey []byte) *Service {
	return &Service{rps: pool, jwtKey: jwtKey}
}

// GetUser _
func (se *Service) GetUser(ctx context.Context, id string) (*user.Person, error) {
	return se.rps.GetUserByID(ctx, id)
}

// GetAllUsers _
func (se *Service) GetAllUsers(ctx context.Context) ([]*user.Person, error) {
	return se.rps.GetAllUsers(ctx)
}

// DeleteUser _
func (se *Service) DeleteUser(ctx context.Context, id string) error {
	return se.rps.DeleteUser(ctx, id)
}

// UpdateUser _
func (se *Service) UpdateUser(ctx context.Context, id string, user *user.Person) error {
	return se.rps.UpdateUser(ctx, id, user)
}
