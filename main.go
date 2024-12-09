package main

import (
	"didactic_octo_sniffle/app/api"
	"didactic_octo_sniffle/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// Initialize the database
func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// Migrate the schema
	err = db.AutoMigrate(models.User{})
	if err != nil {
		log.Println(
			"Database Table User Migration failed!",
			err)
		return nil
	}
	err = db.AutoMigrate(models.Post{})
	if err != nil {
		log.Println(
			"Database Table Posts Migration failed!",
			err)
		return nil
	}
	log.Println("Database operation completed successfully!")
	return db
}

func main() {

	db := initDB()
	router := gin.Default()
	// Define the endpoint
	router.POST("/users", api.CreateUserHandler(db))

	// Start the server
	router.Run(":9090")

}
