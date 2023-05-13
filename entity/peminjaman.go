package entity

import (
	"time"

	"github.com/google/uuid"
)

type Peminjaman struct {
	ID   				uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	BatasPengembalian 	time.Time 	`json:"batas_pengembalian"`
	TanggalPengembalian time.Time 	`json:"tanggal_pengembalian"`

	UserID 				uuid.UUID   `gorm:"foreignKey" json:"user_id"`
	User   				*User 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	
	MangaID 			uuid.UUID   `gorm:"foreignKey" json:"manga_id"`
	Manga   			*Manga 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manga,omitempty"`

	Timestamp
}