package controller

import (
	"net/http"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type PeminjamanController interface {
	CreatePeminjaman(ctx *gin.Context)
	GetAllPeminjamanUser(ctx *gin.Context)
}

type peminjamanController struct {
	jwtService service.JWTService
	peminjamanService service.PeminjamanService
}

func NewPeminjamanController(us service.PeminjamanService) PeminjamanController {
	return &peminjamanController{
		peminjamanService: us,
	}
}

func(uc *peminjamanController) CreatePeminjaman(ctx *gin.Context) {
	var peminjaman dto.PeminjamanCreateDTO
	err := ctx.ShouldBind(&peminjaman)
	result, err := uc.peminjamanService.CreatePeminjaman(ctx.Request.Context(), peminjaman)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Peminjaman", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Peminjaman", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *peminjamanController) GetAllPeminjamanUser(ctx *gin.Context) {
	result, err := uc.peminjamanService.GetAllPeminjamanUser(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Peminjaman", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Peminjaman", result)
	ctx.JSON(http.StatusOK, res)
}