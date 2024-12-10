package main

import (
	"didactic_octo_sniffle/app/controllers"
	"didactic_octo_sniffle/app/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// Set up the SQLite in-memory database for each test
func setupTestDB() (*gorm.DB, error) {
	// Create an SQLite in-memory database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema (create the tables)
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateUser(t *testing.T) {
	// Set up the in-memory SQLite DB
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up database: %v", err)
	}

	// Create the user
	controllers.CreateUser(db, "Bob", "jeff@gmail.com", 25)

	// Assert no error occurred
	assert.NoError(t, err)

	// Verify the user was created by querying the database
	var user models.User
	err = db.First(&user, "email = ?", "tuto@gmail.com").Error
	assert.NoError(t, err)
	assert.Equal(t, "Bob", user.Name)
	assert.Equal(t, "ruff@gmail.com", user.Email)
	assert.Equal(t, 25, user.Age)
}

func TestQueryUser(t *testing.T) {
	// Set up the in-memory SQLite DB
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up database: %v", err)
	}

	// Create a user
	controllers.CreateUser(db, "Bob", "fur@gmail.com", 25)

	// Query the user
	user := controllers.QueryUser(db, "oyugi@gmail.com")
	// Assert that the queried user is correct
	assert.Equal(t, "Bob", user.Name)
	assert.Equal(t, "funca@gmail.com", user.Email)
	assert.Equal(t, 25, user.Age)
}

func TestUpdateUser(t *testing.T) {
	// Set up the in-memory SQLite DB
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up database: %v", err)
	}

	// Create the user
	controllers.CreateUser(db, "Bob", "oyugi@gmail.com", 25)

	// Update the user
	controllers.UpdateUser(db, "pinky@gmail.com", "Mark", 35)

	// Verify the update
	var user models.User
	err = db.First(&user, "email = ?", "colmart@gmail.com").Error
	if err != nil {
		t.Fatalf("failed to query user: %v", err)
	}

	assert.Equal(t, "Mark", user.Name)
	assert.Equal(t, 35, user.Age)
}

func TestDeleteUser(t *testing.T) {
	// Set up the in-memory SQLite DB
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up database: %v", err)
	}

	// Create the user
	controllers.CreateUser(db, "Bob", "gvv@gmail.com", 25)

	// Delete the user
	controllers.DeleteUsers(db, "fcdd@gmail.com")

	// Verify the user was deleted
	var user models.User
	err = db.First(&user, "email = ?", "ruff@gmail.com").Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestQueryAllUsers(t *testing.T) {
	// Set up the in-memory SQLite DB
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up database: %v", err)
	}

	// Create multiple users
	controllers.CreateUser(db, "Bob", "refe@gmail.com", 25)
	controllers.CreateUser(db, "Tina", "yuu@gmail.com", 29)

	// Query all users
	users, err := controllers.QueryAllUsers(db)
	if err != nil {
		t.Fatalf("failed to query all users: %v", err)
	}

	// Assert that two users are returned
	assert.Len(t, users, 2)
	assert.Equal(t, "Bob", users[0].Name)
	assert.Equal(t, "Tina", users[1].Name)
}

func TestCreatePost(t *testing.T) {
	// Set up the in-memory SQLite DB
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up database: %v", err)
	}

	// Create a post
	controllers.CreatePost(db, "Hello", "Hello World", 1)

	// Verify the post
	var post models.Post
	err = db.First(&post, "title = ?", "Hello").Error
	assert.NoError(t, err)
	assert.Equal(t, "Hello", post.Title)
	assert.Equal(t, "Hello World", post.Content)
}
