package helpers

import (
	"go-crud/controllers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func AddRoute(route string, modelName string, model models.Model, apiGroup *gin.RouterGroup) {
	handler := controllers.GetAllWrapper(modelName, model)
	apiGroup.GET(route, handler)
}
