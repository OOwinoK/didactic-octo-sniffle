package main

import (
	"didactic_octo_sniffle/app/controllers"
	"didactic_octo_sniffle/app/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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
		log.Println(
			"Database Table User Migration failed!",
			err)
		return
	}
	err = db.AutoMigrate(models.Post{})
	if err != nil {
		log.Println(
			"Database Table Posts Migration failed!",
			err)
		return
	}
	log.Println("Database operation completed successfully!")

	// Create a new user
	db.Create(&models.User{Name: "Alice"})

	// Create and query records
	//models.CreateAndQueryRecords(db)
	controllers.CreateUser(db, "Bob", "oyugi@gmail.com", 25)
	controllers.CreateUser(db, "Tina", "tina@gmail.com", 29)
	controllers.QueryUser(db, "oyugi@gmail.com")
	users, _ := controllers.QueryAllUsers(db)
	for _, u := range users {
		println("User:", u.Name, u.Email, u.Age)
	}
	user := controllers.QueryUser(db, "oyugi@gmail.com")
	fmt.Println(user)
	controllers.UpdateUser(db, "kevin@gmail.com", "Mark", 35)
	controllers.DeleteUsers(db, "kevin@gmail.com")
	controllers.CreatePost(db, "Hello", "Hello World", 1)
	controllers.CreatePost(db, "Hey World", "Hello World Tommy", 2)
	posts, _ := controllers.QueryAllPost(db)
	for _, p := range posts {
		println("Post:", p.Title, p.Content)
	}
	err = controllers.UpdatePost(db, 2, "Hi tiger ", "Hello Tiger World")
	if err != nil {
		fmt.Println("Update Post Failed")
	}
	err = controllers.DeletePost(db, 1)
	if err != nil {
		fmt.Println(err)
	}

}
