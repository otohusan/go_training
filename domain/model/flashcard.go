package model

import "time"

type Flashcard struct {
	ID         string
	StudySetID string
	Question   string
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
