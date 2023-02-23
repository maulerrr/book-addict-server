package models

type BookTab struct {
	ID       int  `gorm:"primaryKey" json:"id"`
	UserID   int  `gorm:"foreignKey:user_id" json:"user_id"`
	BookID   int  `gorm:"foreignKey:book_id;references:id" json:"book_id"`
	Finished bool `json:"finished"`
}
