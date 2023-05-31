package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func ProvinsiRoutes(router *gin.Engine, ProvinsiController controller.ProvinsiController, jwtService service.JWTService) {
	provinsiRoutes := router.Group("/api/provinsi")
	{
		provinsiRoutes.GET("", ProvinsiController.GetAllProvinsi)
	}
}