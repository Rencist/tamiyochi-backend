package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func SeriRoutes(router *gin.Engine, SeriController controller.SeriController, jwtService service.JWTService) {
	seriRoutes := router.Group("/api/seri")
	{
		seriRoutes.POST("", SeriController.CreateSeri)
		seriRoutes.GET("", middleware.Authenticate(jwtService), SeriController.GetAllSeri)
		seriRoutes.DELETE("/", middleware.Authenticate(jwtService), SeriController.DeleteSeri)
		seriRoutes.PUT("/", middleware.Authenticate(jwtService), SeriController.UpdateSeri)
	}
}