package routes

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func SetupOthersRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/movies/others/details/:title", controllers.GetByTitle("Others", &models.Movie{}))
	apiGroup.GET("/movies/others", controllers.GetAllWrapper("Others", &models.Movie{}))
	apiGroup.POST("/movies/others", controllers.CreateMovie("Others"))
}
