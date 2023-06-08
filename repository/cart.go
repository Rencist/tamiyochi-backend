package repository

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error)
	FindCartByUserID(ctx context.Context, userID uuid.UUID) (dto.CartResponse, error)
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

func(db *cartConnection) FindCartByUserID(ctx context.Context, userID uuid.UUID) (dto.CartResponse, error) {
	var jumlahMangaCart dto.JumlahMangaCart
	var cart []entity.Cart
	
	var cartResponseDTO dto.CartResponse
	var cartDTO dto.Cart
	totalHarga := 0
	totalManga := 0
	ux := db.connection.Select("count(1) as jumlah_penyewa, manga_id").Group("manga_id").Find(&cart)
	for _, res := range cart {
		jumlahMangaCart.JumlahPenyewa = res.JumlahPenyewa
		jumlahMangaCart.MangaID = res.MangaID

		manga := entity.Manga{}
		ux := db.connection.Where("id = ?", res.MangaID).Find(&manga)
		if ux.Error != nil {
			return dto.CartResponse{}, ux.Error
		}
		seri := entity.Seri{}
		ux = db.connection.Where("id = ?", manga.SeriID).Find(&seri)
		if ux.Error != nil {
			return dto.CartResponse{}, ux.Error
		}

		cartDTO.JumlahTersedia = manga.JumlahTersedia
		cartDTO.HargaSewa = manga.HargaSewa
		cartDTO.Volume = manga.Volume
		cartDTO.Foto = seri.Foto
		cartDTO.HargaSubTotal = manga.HargaSewa * res.JumlahPenyewa
		cartDTO.JumlahSewa = res.JumlahPenyewa
		cartDTO.JudulSeri = seri.Judul

		totalHarga += cartDTO.HargaSubTotal
		totalManga += res.JumlahPenyewa
		cartResponseDTO.Cart = append(cartResponseDTO.Cart, cartDTO)
	}
	cartResponseDTO.TotalHargaSewa = totalHarga
	cartResponseDTO.TotalPinjaman = totalManga

	if ux.Error != nil {
		return dto.CartResponse{}, ux.Error
	}
	return cartResponseDTO, nil
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
