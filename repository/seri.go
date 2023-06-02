package repository

import (
	"context"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriRepository interface {
	GetTotalData(ctx context.Context) (int64, error)
	CreateSeri(ctx context.Context, seri entity.Seri) (entity.Seri, error)
	GetAllSeri(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error)
	FindSeriByID(ctx context.Context, seriID uuid.UUID) (entity.Seri, error)
	DeleteSeri(ctx context.Context, seriID uuid.UUID) (error)
	UpdateSeri(ctx context.Context, seri entity.Seri) (error)
}

type seriConnection struct {
	connection *gorm.DB
}

func NewSeriRepository(db *gorm.DB) SeriRepository {
	return &seriConnection{
		connection: db,
	}
}

func (db *seriConnection) GetTotalData(ctx context.Context) (int64, error) {
	var totalData int64
	bc := db.connection.Model(&entity.Seri{}).Count(&totalData)
	if bc.Error != nil {
		return 0, bc.Error
	}
	return totalData, nil
}

func(db *seriConnection) CreateSeri(ctx context.Context, seri entity.Seri) (entity.Seri, error) {
	uc := db.connection.Create(&seri)
	if uc.Error != nil {
		return entity.Seri{}, uc.Error
	}
	return seri, nil
}

func(db *seriConnection) GetAllSeri(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error) {
	var listSeri []entity.Seri
	totalData, _ := db.GetTotalData(ctx)
	tx := db.connection.Debug().Scopes(common.Pagination(&pagination, totalData)).Preload("Mangas").Preload("SeriGenre").Preload("PenulisSeri").Order("total_pembaca desc").Find(&listSeri)
	if tx.Error != nil {
		return dto.PaginationResponse{}, tx.Error
	}
	var listSeriDTO dto.SeriResponseDTO
	var listSeriDTOArray []dto.SeriResponseDTO
	for _, res := range listSeri {
		listSeriDTO.ID = res.ID
		listSeriDTO.Judul = res.Judul
		listSeriDTO.Sinopsis = res.Sinopsis
		listSeriDTO.TahunTerbit = res.TahunTerbit
		listSeriDTO.Skor = res.Skor
		listSeriDTO.TotalPenilai = res.TotalPenilai
		listSeriDTO.TotalPembaca = res.TotalPembaca
		listSeriDTO.PenerbitID = res.PenerbitID
		for _, res := range res.Mangas {
			listSeriDTO.Manga = append(listSeriDTO.Manga, res)
		}
		for _, res := range res.PenulisSeri {
			var penulis entity.Penulis
			db.connection.Where("id = ?", res.PenulisID).Take(&penulis)
			listSeriDTO.Penulis = append(listSeriDTO.Penulis, penulis)
		}
		for _, res := range res.SeriGenre {
			var genre entity.Genre
			db.connection.Where("id = ?", res.GenreID).Take(&genre)
			listSeriDTO.Genre = append(listSeriDTO.Genre, genre)
		}
		
		listSeriDTOArray = append(listSeriDTOArray, listSeriDTO)
	}

	meta := dto.Meta{
		Page: pagination.Page,
		PerPage: pagination.PerPage,
		MaxPage: pagination.MaxPage,
		TotalData: totalData,
	}
	paginationResponse := dto.PaginationResponse{
		DataPerPage: listSeriDTOArray,
		Meta: meta,
		
	}
	return paginationResponse, nil
}

func(db *seriConnection) FindSeriByID(ctx context.Context, seriID uuid.UUID) (entity.Seri, error) {
	var seri entity.Seri
	ux := db.connection.Where("id = ?", seriID).Take(&seri)
	if ux.Error != nil {
		return seri, ux.Error
	}
	return seri, nil
}

func(db *seriConnection) DeleteSeri(ctx context.Context, seriID uuid.UUID) (error) {
	uc := db.connection.Delete(&entity.Seri{}, &seriID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *seriConnection) UpdateSeri(ctx context.Context, seri entity.Seri) (error) {
	uc := db.connection.Updates(&seri)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}