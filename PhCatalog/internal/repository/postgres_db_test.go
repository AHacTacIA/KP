package repository

import (
	"context"
	"log"
	"testing"

	"github.com/AHacTacIA/KP/PhCatalog/internal/catalog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

var (
	Pool *pgxpool.Pool
)

type Service struct { // Service new
	rps Repository
}

func NewService(newRps Repository) *Service { // create
	return &Service{rps: newRps}
}

func TestCreate(t *testing.T) {
	testValidData := []catalog.Medicine{
		{
			Name:  "Citromon",
			Count: 5,
			Price: 19,
		},
		{
			Name:  "Ibuprofen",
			Count: 15,
			Price: 25,
		},
	}
	testNoValidData := []catalog.Medicine{
		{
			Name:  "Nice",
			Count: 30,
			Price: 2,
		},
		{
			Name:  "Ibuprofen",
			Count: 5,
			Price: 15,
		},
		{
			Name:  "Ibuprofen",
			Count: 15,
			Price: 19,
		},
	}
	rps := NewService(&PRepository{Pool: Pool})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, p := range testValidData {
		_, err := rps.rps.CreateMedicine(ctx, &p)
		require.NoError(t, err, "create error")
	}
	for _, p := range testNoValidData {
		_, err := rps.rps.CreateMedicine(ctx, &p)
		require.Error(t, err, "create error")
	}
}
func TestSelectAll(t *testing.T) {
	rps := NewService(&PRepository{Pool: Pool})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := catalog.Medicine{
		Id:    "12",
		Name:  "Noshpa",
		Count: 10,
		Price: 20,
	}

	med, err := rps.rps.GetAllMedicine(ctx)
	require.NoError(t, err, "select all: problems with select all medicines")
	require.Equal(t, 2, len(med), "select all: the values are`t equals")

	_, err = Pool.Exec(ctx, "insert into medicines(id,name,count,price) values($1,$2,$3,$4)", &m.Id, &m.Name, &m.Count, &m.Price)
	require.NoError(t, err, "select all: insert error")
	med, err = rps.rps.GetAllMedicine(ctx)
	if err != nil {
		defer log.Fatalf("error with select all: %v", err)
	}
	require.NotEqual(t, 5, len(med), "select all: the values are equals")
}

func TestSelectById(t *testing.T) {
	rps := NewService(&PRepository{Pool: Pool})
	ctx, cancel := context.WithCancel(context.Background())
	_, err := rps.rps.GetMedicineByID(ctx, "12")
	require.NoError(t, err, "select medicine by id: this id dont exist")
	_, err = rps.rps.GetMedicineByID(ctx, "20")
	require.Error(t, err, "select medicine by id: this id already exist")
	cancel()
}

func TestUpdate(t *testing.T) {
	rps := NewService(&PRepository{Pool: Pool})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testValidData := []*catalog.Medicine{
		{
			Name:  "Citromon",
			Count: 20,
			Price: 19,
		},
		{
			Name:  "query21",
			Count: 8,
			Price: 19,
		},
	}
	testNoValidData := []*catalog.Medicine{
		{
			Name:  "Egor",
			Count: 7,
			Price: 19,
		},
		{
			Name:  "qwerty",
			Count: 6,
			Price: 19,
		},
		{
			Name:  "qwerty1",
			Count: 5,
			Price: 19,
		},
	}
	for _, p := range testValidData {
		err := rps.rps.ChangeMedicine(ctx, "d57d1026-c79a-443d-9d81-714381a37a80", p)
		require.NoError(t, err, "update error")
	}
	for _, p := range testNoValidData {
		err := rps.rps.ChangeMedicine(ctx, "bb839db7-4be3-41a8-a53b-403ad26593ca", p)
		require.Error(t, err, "update error")
	}
	err := rps.rps.ChangeMedicine(ctx, "bb839db7-4be3-41a8-a53b-403ad26593ca", testValidData[0])
	require.Error(t, err, "update error")
}

/*func TestPRepository_Delete(t *testing.T) {
	rps := NewService(&PRepository{Pool: Pool})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := rps.rps.SelectByIDAuth(ctx, "d57d1026-c79a-443d-9d81-714381a37a80")
	require.NoError(t, err, "there is an error")
	_, err = rps.rps.SelectByIDAuth(ctx, "3")
	require.Error(t, err, "there isn`t an error")
}*/
