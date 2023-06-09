package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func PeminjamanRoutes(router *gin.Engine, PeminjamanController controller.PeminjamanController, jwtService service.JWTService) {
	peminjamanRoutes := router.Group("/api/peminjaman")
	{
		peminjamanRoutes.POST("", middleware.Authenticate(jwtService), PeminjamanController.CreatePeminjaman)
		peminjamanRoutes.GET("", middleware.Authenticate(jwtService), PeminjamanController.GetAllPeminjamanUser)
	}
}