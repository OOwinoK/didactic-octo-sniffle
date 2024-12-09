package api

import (
	"didactic_octo_sniffle/app/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// REST API handler for creating a user
func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required,email"`
			Age   int    `json:"age" binding:"required,min=1"`
		}

		// Parse and validate the request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the CreateUser function
		if controllers.CreateUser(db, req.Name, req.Email, req.Age) {
			c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
	}
}
