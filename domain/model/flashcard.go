package model

import "time"

type Flashcard struct {
	ID         string    `json:"id"`
	StudySetID string    `json:"study_set_id"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
