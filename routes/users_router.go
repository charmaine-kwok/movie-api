package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/user", controllers.CreateUser)
	apiGroup.POST("/login", controllers.Login)
}
