package dto

import (
	"time"

	"github.com/google/uuid"
)

type PeminjamanCreateDTO struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	BatasPengembalian 	time.Time 	`json:"batas_pengembalian"`
	TanggalPengembalian time.Time 	`json:"tanggal_pengembalian"`
	
	StatusPeminjaman 	string 		`json:"status_peminjaman"`
	BuktiPembayaran     string 		`json:"bukti_pembayaran" binding:"required"`
	AtasNama			string 		`json:"atas_nama" binding:"required"`

	UserID 				uuid.UUID   `gorm:"foreignKey" json:"user_id"`
}

type PeminjamanReponse struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Volume 			string    `json:"volume"`
	JumlahTersedia 	string    `json:"jumlah_tersedia"`
	HargaSewa 		string    `json:"harga_sewa"`
}