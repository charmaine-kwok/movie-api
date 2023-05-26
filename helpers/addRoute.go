package helpers

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func AddRoute(route string, modelName string, model models.Model, router *gin.Engine) {
	handler := controllers.GetAllWrapper(modelName, model)
	router.GET(route, handler)
}
