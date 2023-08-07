package controllers

import (
	"context"
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Username string `json:"username"`
}

// @Summary Create user
// @Tags Users
// @Description Create user
// @Accept json
// @Produce json
// @Param body body models.User true "User details"
// @Success 200 {object} UserResponse "User details"
// @Failure 400  "Invalid request body"
// @Failure 500  "Internal server error"
// @Router /user [post]
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash passsword",
			})
			return
		}

		// Assign the hashed password back to the user
		user.Password = string(hash)

		// Save the user to the database
		initializers.CRED_DB.Create(&user)

		// Respond with the username
		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
		})
	}
}
