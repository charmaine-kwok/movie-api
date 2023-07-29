package routes

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func SetupMovieRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/movies/movies/details/:title", controllers.GetByTitle("Movies", &models.Movie{}))
	apiGroup.GET("/movies/movies", controllers.GetAllWrapper("Movies", &models.Movie{}))
	apiGroup.POST("/movies/movies", controllers.CreateMovie("Movies"))
}
