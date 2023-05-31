package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func KabupatenRoutes(router *gin.Engine, KabupatenController controller.KabupatenController, jwtService service.JWTService) {
	kabupatenRoutes := router.Group("/api/kabupaten")
	{
		kabupatenRoutes.GET("", KabupatenController.GetAllKabupaten)
		kabupatenRoutes.GET("/:id", KabupatenController.FindKabupatenByProvinsiID)
	}
}