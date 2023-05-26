package main

import (
	docs "go-crud/docs"
	"go-crud/helpers"
	"go-crud/initializers"
	"go-crud/models"
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
//	@host						localhost:3002
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/api")
	{
		helpers.AddRoute("/movies", "Movies", &models.Movie{}, apiGroup)
		helpers.AddRoute("/non-movies", "Non-movies", &models.Other{}, apiGroup)
		helpers.AddRoute("/others", "Others", &models.Other{}, apiGroup)
	}
	r.Run()
}
