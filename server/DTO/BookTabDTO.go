package DTO

type CreateBookTabDto struct {
	BookID   int  `json:"book_id"`
	UserID   int  `json:"user_id"`
	Finished bool `json:"finished"`
}
