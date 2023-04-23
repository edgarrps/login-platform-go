package domain

import "time"

type Base struct {
	ID        string     `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Timer `json:"created_at" gorm:"type:datetime"`
	UpdateAt  time.Timer `json:"udpated_at" gorm:"type:datetime"`
	DeletedAt time.Timer `json:"deleted_at" gorm:"type:datetime"`
}
