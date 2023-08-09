package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupNonMoviesRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/non-movies", controllers.GetAllNonMovies())
	apiGroup.GET("/non-movies/:item_id", controllers.GetNonMoviesByItemId())
	apiGroup.POST("/non-movies", controllers.CreateNonMovies())
}
