package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type CartService interface {
	CreateCart(ctx context.Context, cartDTO dto.CartCreateDTO) (entity.Cart, error)
	FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Cart, error)
	DeleteCart(ctx context.Context, cartID uuid.UUID) (error)
	DeleteAllByMangaIDCart(ctx context.Context, mangaID int) (error)
}

type cartService struct {
	cartRepository repository.CartRepository
}

func NewCartService(ur repository.CartRepository) CartService {
	return &cartService{
		cartRepository: ur,
	}
}

func(us *cartService) CreateCart(ctx context.Context, cartDTO dto.CartCreateDTO) (entity.Cart, error) {
	cart := entity.Cart{}
	err := smapping.FillStruct(&cart, smapping.MapFields(cartDTO))
	if err != nil {
		return cart, err
	}
	return us.cartRepository.CreateCart(ctx, cart)
}

func(us *cartService) FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Cart, error) {
	return us.cartRepository.FindCartByUserID(ctx, userID)
}

func(us *cartService) DeleteCart(ctx context.Context, cartID uuid.UUID) (error) {
	return us.cartRepository.DeleteCart(ctx, cartID)
}

func(us *cartService) DeleteAllByMangaIDCart(ctx context.Context, mangaID int) (error) {
	return us.cartRepository.DeleteAllByMangaIDCart(ctx, mangaID)
}