package models

type Post struct {
	ID       uint `gorm:"primaryKey"`
	Content  *string
	AuthorID string
}
