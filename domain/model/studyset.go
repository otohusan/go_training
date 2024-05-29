package model

import "time"

type StudySet struct {
	ID          string
	UserID      string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
