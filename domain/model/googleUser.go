package model

import "time"

type GoogleUser struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	GoogleID  string    `json:"google_id"`
	CreatedAt time.Time `json:"created_at"`
}
