package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupMovieRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/movies", controllers.GetAllMovies())
	apiGroup.GET("/movies/:item_id", controllers.GetMoviesByItemId())
	apiGroup.POST("/movies", controllers.CreateMovies())
}
