package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error)
	FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Cart, error)
	DeleteCart(ctx context.Context, cartID uuid.UUID) (error)
	DeleteAllByMangaIDCart(ctx context.Context, mangaID int) (error)
}

type cartConnection struct {
	connection *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartConnection{
		connection: db,
	}
}

func(db *cartConnection) CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error) {
	cartID := uuid.New()
	cart.ID = cartID
	uc := db.connection.Create(&cart)
	if uc.Error != nil {
		return entity.Cart{}, uc.Error
	}
	return cart, nil
}

func(db *cartConnection) FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Cart, error) {
	var cart []entity.Cart
	ux := db.connection.Where("user_id = ?", userID).Find(&cart)
	if ux.Error != nil {
		return nil, ux.Error
	}
	return cart, nil
}

func(db *cartConnection) DeleteCart(ctx context.Context, cartID uuid.UUID) (error) {
	uc := db.connection.Delete(&entity.Cart{}, &cartID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *cartConnection) DeleteAllByMangaIDCart(ctx context.Context, mangaID int) (error) {
	uc := db.connection.Where("manga_id = ?", mangaID).Delete(&entity.Cart{})
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
