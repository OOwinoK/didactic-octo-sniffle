package models

import "gorm.io/gorm"

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

func CreateUser(db *gorm.DB, name string, email string, age int) bool {
	user := User{Name: name, Email: email, Age: age}
	db.Create(&user)
	return true
}

func QueryAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error // Return the error if the query fails
	}
	return users, nil // Return the slice of users and no error
}

func QueryUser(db *gorm.DB, email string) User {
	var user User
	db.Where("email = ?", email).First(&user)
	println("User:", user.Name, user.Email, user.Age)
	return user

}

func UpdateUser(db *gorm.DB, email string, name string, age int) bool {
	db.Model(&User{}).Where("email = ?", email).Update("name", name).Update("age", age)
	return true

}

func DeleteUsers(db *gorm.DB, email string) bool {
	db.Where("email = ?", email).Delete(&User{})
	return true
}

func CreatePost(db *gorm.DB, title string, content string, userID uint) bool {
	// Create a post associated with the user
	post := Post{Title: title, Content: content, UserID: userID}
	db.Create(&post)
	return true

}

func QueryAllPost(db *gorm.DB) {
	// Query posts
	var posts []Post
	db.Find(&posts)
	for _, p := range posts {
		println("Post:", p.Title, p.Content)
	}
}

func DeletePost(db *gorm.DB) {

}

func UpdatePost(db *gorm.DB) {

}
