package models

type Book struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

//id SERIAL PRIMARY KEY,
//title VARCHAR(255),
//author VARCHAR(255),
//price INTEGER,
//description TEXT,
//img TEXT
