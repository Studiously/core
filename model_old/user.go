package model_old

import "github.com/satori/go.uuid"

type User struct {
	ID uuid.UUID `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}