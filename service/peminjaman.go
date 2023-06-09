package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"
	"time"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type PeminjamanService interface {
	CreatePeminjaman(ctx context.Context, peminjamanDTO dto.PeminjamanCreateDTO) (entity.Peminjaman, error)
	GetAllPeminjamanUser(ctx context.Context) ([]entity.Peminjaman, error)
}

type peminjamanService struct {
	peminjamanRepository repository.PeminjamanRepository
	cartRepository repository.CartRepository
}

func NewPeminjamanService(ur repository.PeminjamanRepository, cr repository.CartRepository) PeminjamanService {
	return &peminjamanService{
		peminjamanRepository: ur,
		cartRepository: cr,
	}
}

// Menunggu Konfirmasi, Belum Diambil, Sedang Dipinjam, Sudah Dikembalikan
func(us *peminjamanService) CreatePeminjaman(ctx context.Context, peminjamanDTO dto.PeminjamanCreateDTO) (entity.Peminjaman, error) {
	peminjaman := entity.Peminjaman{}
	err := smapping.FillStruct(&peminjaman, smapping.MapFields(peminjamanDTO))
	if err != nil {
		return peminjaman, err
	}
	peminjaman.ID = uuid.New()
	peminjaman.StatusPeminjaman = "Menunggu Konfirmasi"
	peminjaman.BatasPengembalian = time.Now().AddDate(0, 0, 7)
	peminjaman.TanggalPengembalian = time.Time{}
	createPeminjaman, err := us.peminjamanRepository.CreatePeminjaman(ctx, peminjaman)
	if err != nil {
		return peminjaman, err
	}

	cart, err := us.cartRepository.FindCartByUserID(ctx, peminjamanDTO.UserID)
	if err != nil {
		return peminjaman, err
	}

	for _, res := range cart {
		peminjamanManga := entity.PeminjamanManga{}
		peminjamanManga.ID = uuid.New()
		peminjamanManga.MangaID = res.MangaID
		peminjamanManga.PeminjamanID = createPeminjaman.ID
		_, err = us.peminjamanRepository.CreatePeminjamanManga(ctx, peminjamanManga)
		if err != nil {
			return peminjaman, err
		}
	}

	err = us.cartRepository.DeleteAllUserCart(ctx, peminjamanDTO.UserID)
	if err != nil {
		return peminjaman, err
	}
	
	return createPeminjaman, err
}

func(us *peminjamanService) GetAllPeminjamanUser(ctx context.Context) ([]entity.Peminjaman, error) {
	return us.peminjamanRepository.GetAllPeminjamanUser(ctx)
}