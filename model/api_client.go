package model

type ApiClient struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	ApiKey string `db:"api_key"`
	Active bool   `db:"active"`
}

type ApiClientTenant struct {
	ID          int `db:"id"`
	ApiClientID int `db:"apiclient_id"`
	TenantID    int `db:"tenant_id"`
}
