package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	AuthorID *string
}
