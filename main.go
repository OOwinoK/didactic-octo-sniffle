package main

import (
	"fmt"
	"log"

	"didactic_octo_sniffle/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(models.User{})
	if err != nil {
		return
	}

	// Create a new user
	db.Create(&models.User{Name: "Alice"})

	// Create and query records
	//models.CreateAndQueryRecords(db)
	models.CreateUser(db, "Bob", "oyugi@gmail.com", 25)
	models.QueryUser(db, "oyugi@gmail.com")
	log.Println("Database operation completed successfully!")
	users, _ := models.QueryAllUsers(db)
	for _, u := range users {
		println("User:", u.Name, u.Email, u.Age)
	}
	user := models.QueryUser(db, "oyugi@gmail.com")
	fmt.Println(user)
	models.UpdateUser(db, "kevin@gmail.com", "Mark", 35)
	models.DeleteUsers(db, "kevin@gmail.com")
	models.CreatePost(db, "Hello", "Hello World", 1)

}
