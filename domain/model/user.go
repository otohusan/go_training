package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	Name      string
	Password  string
	Email     sql.NullString
	CreatedAt time.Time
}

type UserResponse struct {
	ID        string
	Name      string
	Email     sql.NullString
	CreatedAt time.Time
}

type PublicUser struct {
	ID   string
	Name string
}

type UserCredentials struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type CreatedUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Deadline implements context.Context.
func (User) Deadline() (deadline time.Time, ok bool) {
	panic("unimplemented")
}

// Done implements context.Context.
func (User) Done() <-chan struct{} {
	panic("unimplemented")
}

// Err implements context.Context.
func (User) Err() error {
	panic("unimplemented")
}

// Value implements context.Context.
func (User) Value(key any) any {
	panic("unimplemented")
}
