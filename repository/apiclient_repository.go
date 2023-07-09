package repository

import (
	"github.com/amopromo/gochip/model"
)

type ApiClientRepository interface {
	Create(apiClient *model.ApiClient) error
	Update(apiClient *model.ApiClient) error
	Delete(apiClient *model.ApiClient) error
	Get(ApiKey string) (*model.ApiClient, error)
	List() ([]model.ApiClient, error)
	GetTenantIDByApiClientID(apiClientID int) (int, error)
}