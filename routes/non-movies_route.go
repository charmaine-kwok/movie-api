package routes

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func SetupNonMoviesRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/non-movies/details/:title", controllers.GetByTitle("Non-movies", &models.Other{}))
	apiGroup.GET("/non-movies", controllers.GetAllWrapper("Non-movies", &models.Other{}))
}
