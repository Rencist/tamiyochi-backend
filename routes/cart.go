package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine, CartController controller.CartController, jwtService service.JWTService) {
	cartRoutes := router.Group("/api/cart")
	{
		cartRoutes.POST("", middleware.Authenticate(jwtService), CartController.CreateCart)
		cartRoutes.GET("", middleware.Authenticate(jwtService), CartController.FindCartByUserID)
		cartRoutes.DELETE("/:manga_id", middleware.Authenticate(jwtService), CartController.DeleteCart)
		cartRoutes.DELETE("/manga/:manga_id", middleware.Authenticate(jwtService), CartController.DeleteAllByMangaIDCart)
	}
}