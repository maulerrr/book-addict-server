package DTO

type CreateBookDto struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Img         string `json:"img"`
}
