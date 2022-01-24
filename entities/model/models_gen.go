// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID     *int    `json:"id"`
	Title  string  `json:"title"`
	Author *Person `json:"author"`
}

type NewBook struct {
	Title  string `json:"title"`
	Author int    `json:"author"`
}

type NewPerson struct {
	Nama     string  `json:"nama"`
	Hp       *string `json:"hp"`
	Umur     int     `json:"umur"`
	Password string  `json:"password"`
}

type Person struct {
	ID       string  `json:"id"`
	Nama     string  `json:"nama"`
	Hp       *string `json:"hp"`
	Umur     int     `json:"umur"`
	Password string  `json:"password"`
}
