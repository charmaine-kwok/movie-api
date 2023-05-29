package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupNonMoviesRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/non-movies/details/:title", controllers.GetNonMovieByTitle)
	apiGroup.GET("/non-movies", controllers.GetAllNonMovies)
	// apiGroup.GET("/non-movies", controllers.GetAllWrapper("Non-movies", &models.NonMovie{}))
}
