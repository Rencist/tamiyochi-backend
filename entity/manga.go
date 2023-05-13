package entity

import "github.com/google/uuid"

type Manga struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Volume 			string    `json:"volume"`
	JumlahTersedia 	string    `json:"jumlah_tersedia"`
	HargaSewa 		string    `json:"harga_sewa"`

	SeriID 			uuid.UUID   `gorm:"foreignKey" json:"seri_id"`
	Seri   			*Seri 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seri,omitempty"`

	Timestamp
}