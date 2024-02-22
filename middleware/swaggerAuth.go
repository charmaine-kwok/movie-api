package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// SwaggerAuth authenticates using credentials from environment variables
func SwaggerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve username and password from environment variables
		username := os.Getenv("SWAGGER_USERNAME")
		password := os.Getenv("SWAGGER_PASSWORD")

		// Check for missing credentials and fail fast if not found
		if username == "" || password == "" {
			if gin.Mode() == gin.DebugMode {
				log.Println("WARNING: Using default credentials for Swagger UI in development mode. Do not use default credentials in production!")
				username, password = "defaultUsername", "defaultPassword"
			} else {
				log.Fatal("Swagger authentication credentials are not set. Aborting application start.")
			}
		}

		user, pass, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == username && pass == password {
			c.Next()
		} else {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization Required"})
		}
	}
}
