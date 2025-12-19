package domain

import "time"

type Role struct {
	ID          uint
	Name        string
	Code        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
