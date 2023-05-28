package routes

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func SetupOthersRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/others", controllers.GetAllWrapper("Others", &models.Other{}))
}
