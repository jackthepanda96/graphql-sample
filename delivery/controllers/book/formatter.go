package book

type BookRequestFormat struct {
	Title  string `json:"title" form:"title"`
	Author int    `json:"author" form:"author"`
}
