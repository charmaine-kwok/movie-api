package main

import (
	"go-crud/helpers"
	"go-crud/initializers"
	"go-crud/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	helpers.AddRoute("/movies", "Movies", &models.Movie{}, r)
	helpers.AddRoute("/non-movies", "Non-movies", &models.Other{}, r)
	helpers.AddRoute("/others", "Others", &models.Other{}, r)
	r.Run()
}
