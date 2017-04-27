package model_old

import "github.com/satori/go.uuid"

type Profile struct {
	ID uuid.UUID `db:"id" json:"id"`
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	ClassID uuid.UUID `db:"class_id" json:"class_id"`
	Role string `db:"role" json:"role"`
}
