package routes

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func SetupNonMoviesRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/non-movies", controllers.GetAllWrapper("Non-movies", &models.Other{}))
}
