package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"gorm.io/gorm"
)

type PeminjamanRepository interface {
	CreatePeminjaman(ctx context.Context, peminjaman entity.Peminjaman) (entity.Peminjaman, error)
	GetAllPeminjamanUser(ctx context.Context) ([]entity.Peminjaman, error)
	CreatePeminjamanManga(ctx context.Context, peminjamanManga entity.PeminjamanManga) (entity.PeminjamanManga, error)
}

type peminjamanConnection struct {
	connection *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) PeminjamanRepository {
	return &peminjamanConnection{
		connection: db,
	}
}

func(db *peminjamanConnection) CreatePeminjaman(ctx context.Context, peminjaman entity.Peminjaman) (entity.Peminjaman, error) {
	uc := db.connection.Create(&peminjaman)
	if uc.Error != nil {
		return entity.Peminjaman{}, uc.Error
	}
	return peminjaman, nil
}

func(db *peminjamanConnection) GetAllPeminjamanUser(ctx context.Context) ([]entity.Peminjaman, error) {
	var listPeminjaman []entity.Peminjaman
	tx := db.connection.Find(&listPeminjaman)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listPeminjaman, nil
}

func(db *peminjamanConnection) CreatePeminjamanManga(ctx context.Context, peminjamanManga entity.PeminjamanManga) (entity.PeminjamanManga, error) {
	uc := db.connection.Create(&peminjamanManga)
	if uc.Error != nil {
		return entity.PeminjamanManga{}, uc.Error
	}
	return peminjamanManga, nil
}