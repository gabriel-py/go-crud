package repository

import (
	"github.com/amopromo/gochip/model"
)

type DestinationRepository interface {
	Create(name string, slug string) (*model.Destination, error)
	Update(destination *model.Destination) error
	Delete(destination *model.Destination) error
	Get(destination *model.Destination) error
	List() ([]model.Destination, error)
}
