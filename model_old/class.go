package model_old

import "github.com/satori/go.uuid"

type Class struct {
	ID uuid.UUID `db:"id" json:"id" gorm:"primary_key"`
	Name string `db:"name" json:"name"`
}