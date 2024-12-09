package models

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
	Age   int
	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
}
