package main

import (
	docs "go-crud/docs"
	"go-crud/initializers"
	"go-crud/routes"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

//	@title				     	Movie App
//	@version					1.0.0
//	@description				This is an API server for communication between mobile application and MongoDB Database
//	@host						localhost:8080
//	@BasePath					/api
//
// @contact.name charmaine.kwok
// @license.name Apache 2.0
// schemes http
func main() {
	// Read the environment variable and set the Gin mode based on the environment variable
	if env := os.Getenv("GIN_MODE"); env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create a new Gin router
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	apiGroup := r.Group("/api")
	{
		// Add Swagger documentation endpoint
		apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Register the movies routes
		routes.SetupMovieRoutes(apiGroup)
		// Register the non-movies routes
		routes.SetupNonMoviesRoutes(apiGroup)
		// Register the others routes
		routes.SetupOthersRoutes(apiGroup)
	}

	// Start the server
	r.Run(":8080")
}
