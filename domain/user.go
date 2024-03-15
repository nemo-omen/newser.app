package domain

import (
	"time"
)

type User struct {
	Person
	CreatedAt time.Time
	UpdatedAt time.Time
}
