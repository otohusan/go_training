package model

import (
	"time"
)

type Favorite struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	StudySetID string    `json:"study_set_id"`
	CreatedAt  time.Time `json:"created_at"`
}
