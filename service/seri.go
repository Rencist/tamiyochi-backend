package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type SeriService interface {
	CreateSeri(ctx context.Context, seriDTO dto.SeriCreateDTO) (entity.Seri, error)
	GetAllSeri(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error)
	DeleteSeri(ctx context.Context, seriID uuid.UUID) (error)
	UpdateSeri(ctx context.Context, seriDTO dto.SeriUpdateDTO) (error)
}

type seriService struct {
	seriRepository repository.SeriRepository
}

func NewSeriService(ur repository.SeriRepository) SeriService {
	return &seriService{
		seriRepository: ur,
	}
}

func(us *seriService) CreateSeri(ctx context.Context, seriDTO dto.SeriCreateDTO) (entity.Seri, error) {
	seri := entity.Seri{}
	err := smapping.FillStruct(&seri, smapping.MapFields(seriDTO))
	if err != nil {
		return seri, err
	}
	return us.seriRepository.CreateSeri(ctx, seri)
}

func(us *seriService) GetAllSeri(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error) {
	return us.seriRepository.GetAllSeri(ctx, pagination)
}

func(us *seriService) DeleteSeri(ctx context.Context, seriID uuid.UUID) (error) {
	return us.seriRepository.DeleteSeri(ctx, seriID)
}

func(us *seriService) UpdateSeri(ctx context.Context, seriDTO dto.SeriUpdateDTO) (error) {
	seri := entity.Seri{}
	err := smapping.FillStruct(&seri, smapping.MapFields(seriDTO))
	if err != nil {
		return err
	}
	return us.seriRepository.UpdateSeri(ctx, seri)
}