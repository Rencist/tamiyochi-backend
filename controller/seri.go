package controller

import (
	"net/http"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SeriController interface {
	CreateSeri(ctx *gin.Context)
	GetAllSeri(ctx *gin.Context)
	DeleteSeri(ctx *gin.Context)
	UpdateSeri(ctx *gin.Context)
}

type seriController struct {
	jwtService service.JWTService
	seriService service.SeriService
}

func NewSeriController(us service.SeriService) SeriController {
	return &seriController{
		seriService: us,
	}
}

func(uc *seriController) CreateSeri(ctx *gin.Context) {
	var seri dto.SeriCreateDTO
	err := ctx.ShouldBind(&seri)
	result, err := uc.seriService.CreateSeri(ctx.Request.Context(), seri)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Seri", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) GetAllSeri(ctx *gin.Context) {
	result, err := uc.seriService.GetAllSeri(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Seri", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) DeleteSeri(ctx *gin.Context) {
	seriID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = uc.seriService.DeleteSeri(ctx.Request.Context(), seriID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Seri", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) UpdateSeri(ctx *gin.Context) {
	seriID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var seri dto.SeriUpdateDTO
	seri.ID = seriID
	err = ctx.ShouldBind(&seri)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = uc.seriService.UpdateSeri(ctx.Request.Context(), seri)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Seri", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}