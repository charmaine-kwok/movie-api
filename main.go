package main

import (
	"go-crud/controllers"
	"go-crud/initializers"
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
	r.GET("/movies", controllers.GetAllWrapper("Movies"))
	r.GET("/non-movies", controllers.GetAllWrapper("Non-movies"))
	r.GET("/others", controllers.GetAllWrapper("Others"))
	r.Run()
}
