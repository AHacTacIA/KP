package service

import (
	"context"
	"github.com/AHacTacIA/KP/PhCatalog/internal/catalog"
	"github.com/AHacTacIA/KP/PhCatalog/internal/repository"
)

type Service struct {
	jwtKey []byte
	rps    repository.Repository
}

// NewService create new service connection
func NewService(pool repository.Repository, jwtKey []byte) *Service {
	return &Service{rps: pool, jwtKey: jwtKey}
}

// GetMedicine _
func (se *Service) GetMedicine(ctx context.Context, id string) (*catalog.Medicine, error) {
	return se.rps.GetMedicineByID(ctx, id)
}

// GetAllMedicine _
func (se *Service) GetAllMedicine(ctx context.Context) ([]*catalog.Medicine, error) {
	return se.rps.GetAllMedicines(ctx)
}

// DeleteMedicine _
func (se *Service) DeleteMedicine(ctx context.Context, id string) error {
	return se.rps.DeleteMedicine(ctx, id)
}

// ChangeMedicine _
func (se *Service) ChangeMedicine(ctx context.Context, id string, user *catalog.Medicine) error {
	return se.rps.ChangeMedicine(ctx, id, user)
}

// CreateMedicine create new medicine, add him to db
func (se *Service) CreateMedicine(ctx context.Context, m *catalog.Medicine) (string, error) {
	return se.rps.CreateMedicine(ctx, m)
}
