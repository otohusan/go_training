package model

import "time"

type EmailVerification struct {
	Email     string
	Token     string
	Username  string
	Password  string
	CreatedAt time.Time
}
