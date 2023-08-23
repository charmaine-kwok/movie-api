package main

import (
	"go-crud/controllers"
	docs "go-crud/docs"
	"go-crud/initializers"
	"go-crud/middleware"
	"go-crud/routes"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func init() {
	// initializers.LoadEnvVariables() // uncomment this line only for local use
	initializers.ConnectToDB()
	// initializers.SyncDatabase()
}

//	@title			Movie Api
//	@version		1.0.0
//	@description	This is an API server for communication between mobile application and PostgreSQL Database.
//
// go-crud.fly.dev=cloud localhost:8080=local
//
//	@host			go-crud.fly.dev
//	@BasePath		/api
//	@contact.name	charmaine.kwok
//	@license.name	Apache 2.0
//
//	@schemes		https
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

		// route to validate JWT
		apiGroup.GET("/validate", middleware.RequireAuth, controllers.ValidateJWT)

		// Register the users routes
		routes.SetupUsersRoutes(apiGroup)

		// Register the route to fetch user info calling google api
		apiGroup.GET("/userInfo", controllers.GoogleOAuth2GetUserInfo)

		// Apply RequireAuth middleware to all requests below
		apiGroup.Use(middleware.RequireAuth)

		// Register the movies routes
		routes.SetupMovieRoutes(apiGroup)

		// Register the non-movies routes
		routes.SetupNonMoviesRoutes(apiGroup)
	}

	// Start the server
	r.Run(":8080")
}
