package main

import (
	docs "go-crud/docs"
	"go-crud/initializers"
	"go-crud/routes"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		routes.SetupMovieRoutes(apiGroup)
		routes.SetupNonMoviesRoutes(apiGroup)
		routes.SetupOthersRoutes(apiGroup)
	}
	r.Run()
}
