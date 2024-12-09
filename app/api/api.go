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

// REST API handler for querying all users
func QueryAllUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := controllers.QueryAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return users in the response
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

// REST API handler for querying a user by email
func QueryUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email") // Get email from URL parameter

		// Call QueryUser function
		user := controllers.QueryUser(db, email)
		if user.ID == 0 { // If the user was not found (ID is 0 for non-existent records)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Return the found user
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

// REST API handler for updating a user by email
func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name string `json:"name" binding:"required"`
			Age  int    `json:"age" binding:"required,min=1"`
		}

		email := c.Param("email") // Get email from URL parameter

		// Parse and validate the request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateUser function
		if controllers.UpdateUser(db, email, req.Name, req.Age) {
			c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		}
	}
}
