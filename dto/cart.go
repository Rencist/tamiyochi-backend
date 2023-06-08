package dto

import "github.com/google/uuid"

type CartCreateDTO struct {
	ID   	uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	
	MangaID int    `gorm:"foreignKey" json:"manga_id" binding:"required"`
	
	UserID uuid.UUID `gorm:"foreignKey" json:"user_id" binding:"required"`
}

type CartReponse struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Volume 			string    `json:"volume"`
	JumlahTersedia 	string    `json:"jumlah_tersedia"`
	HargaSewa 		string    `json:"harga_sewa"`
}