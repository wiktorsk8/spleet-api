package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
