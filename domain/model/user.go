package model

// User はユーザー情報を表すドメインモデルです。
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
