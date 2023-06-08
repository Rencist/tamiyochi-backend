package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/mashingan/smapping"
)

type PeminjamanService interface {
	CreatePeminjaman(ctx context.Context, peminjamanDTO dto.PeminjamanCreateDTO) (entity.Peminjaman, error)
	GetAllPeminjamanUser(ctx context.Context) ([]entity.Peminjaman, error)
}

type peminjamanService struct {
	peminjamanRepository repository.PeminjamanRepository
}

func NewPeminjamanService(ur repository.PeminjamanRepository) PeminjamanService {
	return &peminjamanService{
		peminjamanRepository: ur,
	}
}

func(us *peminjamanService) CreatePeminjaman(ctx context.Context, peminjamanDTO dto.PeminjamanCreateDTO) (entity.Peminjaman, error) {
	peminjaman := entity.Peminjaman{}
	err := smapping.FillStruct(&peminjaman, smapping.MapFields(peminjamanDTO))
	if err != nil {
		return peminjaman, err
	}
	return us.peminjamanRepository.CreatePeminjaman(ctx, peminjaman)
}

func(us *peminjamanService) GetAllPeminjamanUser(ctx context.Context) ([]entity.Peminjaman, error) {
	return us.peminjamanRepository.GetAllPeminjamanUser(ctx)
}