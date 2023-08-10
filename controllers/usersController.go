package controllers

import (
	"context"
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates user
//
//	@Summary		Create user
//	@Tags			Users
//	@Description	Create user
//	@Accept			json
//	@Param			body	body	models.User	true	"User details"
//	@Success		201		"User created"
//	@Failure		400		"Invalid request body"
//	@Failure		500		"Internal server error"
//	@Router			/user [post]
func CreateUser(c *gin.Context) {
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
	initializers.DB.Create(&user)

	// Respond
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, "User created")
}

// Login user
//
//	@Summary		Login user
//	@Tags			Users
//	@Description	Login user
//	@Accept			json
//	@Param			body	body		models.User	true	"User details"
//	@Success		200		{string}	string		"JWT token"
//	@Failure		400		"Invalid username or password"
//	@Failure		500		"Internal server error"
//	@Router			/login [post]
func Login(c *gin.Context) {
	// Create a context with a timeout
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	// Look up requested user
	initializers.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Compare sent in pass with saved user password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Respond with JWT
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, tokenString)
}
