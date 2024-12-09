package controllers

import (
	"didactic_octo_sniffle/app/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func CreateUser(db *gorm.DB, name string, email string, age int) bool {
	user := models.User{Name: name, Email: email, Age: age}
	db.Create(&user)
	fmt.Println("User created successfully!")
	return true
}

func QueryAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error // Return the error if the query fails
	}
	fmt.Println("Users: %v", users)
	return users, nil // Return the slice of users and no error
}

func QueryUser(db *gorm.DB, email string) models.User {
	var user models.User
	db.Where("email = ?", email).First(&user)
	println("User Found:", user.Name, user.Email, user.Age)
	return user

}

func UpdateUser(db *gorm.DB, email string, name string, age int) bool {
	db.Model(&models.User{}).Where("email = ?", email).Update("name", name).Update("age", age)
	fmt.Println("User updated successfully!")
	return true

}

func DeleteUsers(db *gorm.DB, email string) bool {
	db.Where("email = ?", email).Delete(&models.User{})
	fmt.Println("User deleted successfully!")
	return true
}

func CreatePost(db *gorm.DB, title string, content string, userID uint) bool {
	// Create a post associated with the user
	post := models.Post{Title: title, Content: content, UserID: userID}
	db.Create(&post)
	fmt.Println("Post created successfully!")
	return true

}

func QueryAllPost(db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post
	result := db.Find(&posts) // Execute the query

	if result.Error != nil { // Check for errors in the result
		log.Printf("failed to query posts: %v", result.Error)
		return nil, result.Error // Return nil for the slice and the error
	}

	fmt.Println("Posts retrieved successfully")
	return posts, nil // Return the slice of posts and nil for the error
}

func DeletePost(db *gorm.DB, id uint) error {
	err := db.Delete(&models.Post{}, id) // Delete the post by its primary key
	if err.Error != nil {
		log.Fatalf("failed to delete the post %v", err)
		return err.Error
	}
	fmt.Println(
		"Post deleted successfully!")
	return nil
}

func UpdatePost(db *gorm.DB, id uint, title string, content string) error {
	result := db.Model(&models.Post{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":   title,
		"content": content,
	}) // Use Updates to set multiple fields at once

	if result.Error != nil { // Check for errors
		log.Printf("Failed to update the post: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 { // Optional: Check if any rows were updated
		log.Printf("No post found with ID %d to update", id)
		return fmt.Errorf("no post found with ID %d", id)
	}

	fmt.Println("Post updated successfully!")
	return nil
}
