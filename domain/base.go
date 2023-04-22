package domain

import "time"

type Base struct {
	ID        string
	CreatedAt time.Timer
	UpdateAt  time.Timer
	DeletedAt time.Timer
}
