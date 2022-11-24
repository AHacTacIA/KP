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

func NewService(pool repository.Repository, jwtKey []byte) *Service {
	return &Service{rps: pool, jwtKey: jwtKey}
}

func (se *Service) CreateMedicine(ctx context.Context, p *catalog.Medicine) (string, error) {
	return se.rps.CreateMedicine(ctx, p)
}

func (se *Service) GetMedicine(ctx context.Context, id string) (*catalog.Medicine, error) {
	return se.rps.GetMedicineByID(ctx, id)
}

func (se *Service) DeleteMedicine(ctx context.Context, id string) error {
	return se.rps.DeleteMedicine(ctx, id)
}

func (se *Service) ChangeMedicine(ctx context.Context, id string, med *catalog.Medicine) error {
	return se.rps.ChangeMedicine(ctx, id, med)
}

func (se *Service) GetAllMedicine(ctx context.Context) ([]*catalog.Medicine, error) {
	return se.rps.GetAllMedicines(ctx)
}
