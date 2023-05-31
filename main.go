package main

import (
	"net/http"
	"os"
	"tamiyochi-backend/common"
	"tamiyochi-backend/config"
	"tamiyochi-backend/controller"
	"tamiyochi-backend/repository"
	"tamiyochi-backend/routes"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		res := common.BuildErrorResponse("Gagal Terhubung ke Server", err.Error(), common.EmptyObj{})
		(*gin.Context).JSON((&gin.Context{}), http.StatusBadGateway, res)
		return
	}

	var (
		db *gorm.DB = config.SetupDatabaseConnection()
		
		jwtService service.JWTService = service.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		seriRepository repository.SeriRepository = repository.NewSeriRepository(db)

		userService service.UserService = service.NewUserService(userRepository)
		seriServiec service.SeriService = service.NewSeriService(seriRepository)

		userController controller.UserController = controller.NewUserController(userService, jwtService)
		seriController controller.SeriController = controller.NewSeriController(seriServiec)
	)

	server := gin.Default()
	routes.UserRoutes(server, userController, jwtService)
	routes.SeriRoutes(server, seriController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run("127.0.0.1:" + port)
}