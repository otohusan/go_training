package model

type Post struct {
	ID     string
	Title  string
	Detail string
	// userのidをもつ
	Author string `json:"author"`
}

type RequestedPost struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	// userのidをもつ
	Author string `json:"author"`
}
