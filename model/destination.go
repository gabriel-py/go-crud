package model

type Destination struct {
	ID           int `db:"id"`
	Name         string `db:"name"`
	Slug         string `db:"slug"`
	Active       bool `db:"active"`
}
