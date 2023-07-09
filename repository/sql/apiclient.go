package sql

import (
	"github.com/amopromo/gochip/model"
	"github.com/jmoiron/sqlx"
)

type ApiClientRepository struct {
	DB *sqlx.DB
}

func NewApiClientRepository(db *sqlx.DB) *ApiClientRepository {
	return &ApiClientRepository{DB: db}
}

func (r *ApiClientRepository) Create(apiClient *model.ApiClient) error {
	return nil
}

func (r *ApiClientRepository) Update(apiClient *model.ApiClient) error {
	return nil
}

func (r *ApiClientRepository) Delete(apiClient *model.ApiClient) error {
	return nil
}

func (r *ApiClientRepository) List() ([]model.ApiClient, error) {
	return nil, nil
}

func (r *ApiClientRepository) Get(ApiKey string) (*model.ApiClient, error) {
	var apiClient model.ApiClient

	err := r.DB.Get(&apiClient, "SELECT * FROM api_client WHERE api_key = $1 and active = True", ApiKey)

	if err != nil {
		return nil, err
	}

	return &apiClient, nil
}

func (r *ApiClientRepository) GetTenantIDByApiClientID(apiClientID int) (int, error) {
	var ApiClientTenant model.ApiClientTenant

	err := r.DB.Get(&ApiClientTenant, "SELECT * FROM api_client_tenant WHERE apiclient_id = $1", apiClientID)

	if err != nil {
		return 0, err
	}

	return ApiClientTenant.TenantID, nil
}
