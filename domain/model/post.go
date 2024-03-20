package model

type Post struct {
	ID     string
	Title  string
	Detail string
	// userのidをもつ
	Author string `json:"author"`
}
