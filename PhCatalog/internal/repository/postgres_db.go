package repository

import (
	"context"
	"fmt"
	"github.com/AHacTacIA/KP/PhCatalog/internal/catalog"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// PRepository p
type PRepository struct {
	Pool *pgxpool.Pool
}

// CreateMedicine add Medicine to db
func (p *PRepository) CreateMedicine(ctx context.Context, med *catalog.Medicine) (string, error) {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into medicines(name, count, price, id) values($1,$2,$3,$4)",
		&med.Name, &med.Count, &med.Price, newID)
	if err != nil {
		log.Errorf("database error with create medicine: %v", err)
		return "", err
	}
	return newID, nil
}

// GetUserByID select user by id
func (p *PRepository) GetMedicineByID(ctx context.Context, idMedicine string) (*catalog.Medicine, error) {
	u := catalog.Medicine{}
	err := p.Pool.QueryRow(ctx, "select name,count,price,id from medicines where id=$4", idMedicine).Scan(
		&u.Name, &u.Count, &u.Price, &u.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &catalog.Medicine{}, fmt.Errorf("medicine with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return &catalog.Medicine{}, err
	}
	return &u, nil
}

// GetAllUsers select all users from db
func (p *PRepository) GetAllMedicine(ctx context.Context) ([]*catalog.Medicine, error) {
	var medicines []*catalog.Medicine
	rows, err := p.Pool.Query(ctx, "select name,count,price,id from medicines")
	if err != nil {
		log.Errorf("database error with select all medicines, %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		per := catalog.Medicine{}
		err = rows.Scan(&per.Name, &per.Count, &per.Price, &per.Id)
		if err != nil {
			log.Errorf("database error with select all medicines, %v", err)
			return nil, err
		}
		medicines = append(medicines, &per)
	}

	return medicines, nil
}

// DeleteUser delete user by id
func (p *PRepository) DeleteMedicine(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from medicines where id=$4", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("medicine with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("medicine with this id doesnt exist: %v", err)
		}
		log.Errorf("error with delete medicine %v", err)
		return err
	}
	return nil
}

// UpdateUser update parameters for user
func (p *PRepository) ChangeMedicine(ctx context.Context, id string, med *catalog.Medicine) error {
	a, err := p.Pool.Exec(ctx, "update medicines set name=$1,count=$2,price=$3 where id=$4", &med.Name, &med.Count, &med.Price, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("medicine with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update medicine %v", err)
		return err
	}
	return nil
}
