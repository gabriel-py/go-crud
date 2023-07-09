package sql

import (
	"github.com/amopromo/gochip/model"
	"github.com/jmoiron/sqlx"
)

type DestinationRepository struct {
	DB *sqlx.DB
}

func NewDestinationRepository(db *sqlx.DB) *DestinationRepository {
	return &DestinationRepository{DB: db}
}

func (r *DestinationRepository) Create(name string, slug string) (*model.Destination, error) {
	var Destination *model.Destination
	err := r.DB.Select(&Destination, "INSERT INTO destination (name, slug, active) VALUES ($1, $2, true)", name, slug)
	if err != nil {
		return nil, err
	}
	return Destination, nil
}

func (r *DestinationRepository) Update(destination *model.Destination) error {
	return nil
}

func (r *DestinationRepository) Delete(destination *model.Destination) error {
	return nil
}

func (r *DestinationRepository) Get(destination *model.Destination) error {
	return nil
}

func (r *DestinationRepository) List() ([]model.Destination, error) {
	var Destinations []model.Destination
	err := r.DB.Select(&Destinations, "SELECT * FROM destination")
	if err != nil {
		return nil, err
	}
	return Destinations, nil
}
