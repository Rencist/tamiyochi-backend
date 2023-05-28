package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriRepository interface {
	CreateSeri(ctx context.Context, seri entity.Seri) (entity.Seri, error)
	GetAllSeri(ctx context.Context) ([]entity.Seri, error)
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

func(db *seriConnection) CreateSeri(ctx context.Context, seri entity.Seri) (entity.Seri, error) {
	seri.ID = uuid.New()
	uc := db.connection.Create(&seri)
	if uc.Error != nil {
		return entity.Seri{}, uc.Error
	}
	return seri, nil
}

func(db *seriConnection) GetAllSeri(ctx context.Context) ([]entity.Seri, error) {
	var listSeri []entity.Seri
	tx := db.connection.Preload("Mangas").Find(&listSeri)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listSeri, nil
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