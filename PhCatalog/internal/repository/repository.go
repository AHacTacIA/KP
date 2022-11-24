package repository

import (
	"context"
	"github.com/AHacTacIA/KP/PhCatalog/internal/catalog"
)

type Repository interface {
	CreateMedicine(ctx context.Context, p *catalog.Medicine) (string, error)
	GetMedicineByID(ctx context.Context, idPerson string) (*catalog.Medicine, error)
	GetAllMedicines(ctx context.Context) ([]*catalog.Medicine, error)
	DeleteMedicine(ctx context.Context, id string) error
	ChangeMedicine(ctx context.Context, id string, per *catalog.Medicine) error
}
