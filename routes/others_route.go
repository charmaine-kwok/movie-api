package routes

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func SetupOthersRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/others/details/:title", controllers.GetByTitle("Others", &models.Movie{}))
	apiGroup.GET("/others", controllers.GetAllWrapper("Others", &models.Movie{}))
}
