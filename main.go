package main

import (
	"net/http"
	"os"
	"tamiyochi-backend/common"
	"tamiyochi-backend/config"
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
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
		provinsiRepository repository.ProvinsiRepository = repository.NewProvinsiRepository(db)
		kabupatenRepository repository.KabupatenRepository = repository.NewKabupatenRepository(db)
		komentarRepository repository.KomentarRepository = repository.NewKomentarRepository(db)

		userService service.UserService = service.NewUserService(userRepository)
		seriServiec service.SeriService = service.NewSeriService(seriRepository)
		provinsiService service.ProvinsiService = service.NewProvinsiService(provinsiRepository)
		kabupatenService service.KabupatenService = service.NewKabupatenService(kabupatenRepository)
		komentarService service.KomentarService = service.NewKomentarService(komentarRepository)

		userController controller.UserController = controller.NewUserController(userService, jwtService)
		seriController controller.SeriController = controller.NewSeriController(seriServiec)
		provinsiController controller.ProvinsiController = controller.NewProvinsiController(provinsiService)
		kabupatenController controller.KabupatenController = controller.NewKabupatenController(kabupatenService)
		komentarController controller.KomentarController = controller.NewKomentarController(komentarService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	
	routes.UserRoutes(server, userController, jwtService)
	routes.SeriRoutes(server, seriController, jwtService)
	routes.ProvinsiRoutes(server, provinsiController, jwtService)
	routes.KabupatenRoutes(server, kabupatenController, jwtService)
	routes.KomentarRoutes(server, komentarController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run("127.0.0.1:" + port)
}