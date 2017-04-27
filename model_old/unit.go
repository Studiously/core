package model_old

import "github.com/satori/go.uuid"

type Unit struct {
	ID      uuid.UUID `db:"id" json:"id"`
	ClassID uuid.UUID `db:"class_id" json:"class_id"`
	Name    string `db:"name" json:"name"`
	Weight  int `db:"weight" json:"weight"`
}
