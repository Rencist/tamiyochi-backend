package repository

import (
	"context"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"

	"gorm.io/gorm"
)

type SeriRepository interface {
	GetTotalData(ctx context.Context) (int64, error)
	CreateSeri(ctx context.Context, seri entity.Seri) (entity.Seri, error)
	GetAllSeri(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error)
	FindSeriByID(ctx context.Context, seriID int) (dto.SeriResponseDTO, error)
	DeleteSeri(ctx context.Context, seriID int) (error)
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
	var listSeriDTOArray []dto.SeriResponseDTO
	for _, res := range listSeri {
		var listSeriDTO dto.SeriResponseDTO
		listSeriDTO.ID = res.ID
		listSeriDTO.Judul = res.Judul
		listSeriDTO.Sinopsis = res.Sinopsis
		listSeriDTO.TahunTerbit = res.TahunTerbit
		listSeriDTO.Skor = res.Skor
		listSeriDTO.TotalPenilai = res.TotalPenilai
		listSeriDTO.TotalPembaca = res.TotalPembaca
		listSeriDTO.PenerbitID = res.PenerbitID
		listSeriDTO.Foto = res.Foto
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

func(db *seriConnection) FindSeriByID(ctx context.Context, seriID int) (dto.SeriResponseDTO, error) {
	var seri entity.Seri
	ux := db.connection.Where("id = ?", seriID).Preload("Mangas").Preload("SeriGenre").Preload("PenulisSeri").Take(&seri)
	var seriDTO dto.SeriResponseDTO
	seriDTO.ID = seri.ID
	seriDTO.Judul = seri.Judul
	seriDTO.Sinopsis = seri.Sinopsis
	seriDTO.TahunTerbit = seri.TahunTerbit
	seriDTO.Skor = seri.Skor
	seriDTO.TotalPenilai = seri.TotalPenilai
	seriDTO.TotalPembaca = seri.TotalPembaca
	seriDTO.PenerbitID = seri.PenerbitID
	seriDTO.Foto = seri.Foto
	for _, res := range seri.Mangas {
		seriDTO.Manga = append(seriDTO.Manga, res)
	}
	for _, res := range seri.PenulisSeri {
		var penulis entity.Penulis
		db.connection.Where("id = ?", res.PenulisID).Take(&penulis)
		seriDTO.Penulis = append(seriDTO.Penulis, penulis)
	}
	for _, res := range seri.SeriGenre {
		var genre entity.Genre
		db.connection.Where("id = ?", res.GenreID).Take(&genre)
		seriDTO.Genre = append(seriDTO.Genre, genre)
	}
	if ux.Error != nil {
		return seriDTO, ux.Error
	}
	return seriDTO, nil
}

func(db *seriConnection) DeleteSeri(ctx context.Context, seriID int) (error) {
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