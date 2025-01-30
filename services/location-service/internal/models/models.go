package models

type Location struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Coordinates string `db:"coordinates" json:"coordinates"`
}
