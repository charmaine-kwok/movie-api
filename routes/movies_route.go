package routes

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func SetupMovieRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/movies/movie/:title", controllers.GetByTitle("Movies", &models.Movie{}))
	apiGroup.GET("/movies", controllers.GetAllWrapper("Movies", &models.Movie{}))
}
