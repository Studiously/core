package model_old

import "github.com/satori/go.uuid"

type LocalIdentity struct {
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	Password string `db:"password" json:"password"`
}
