package entity

import "github.com/google/uuid"

type KomentarManga struct {
	ID   		uuid.UUID 		`gorm:"primary_key;not_null" json:"id"`
	
	KomentarID 	uuid.UUID		`gorm:"foreignKey" json:"komentar_id"`
	Komentar   	*Komentar 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"komentar,omitempty"`

	MangaID 	uuid.UUID		`gorm:"foreignKey" json:"manga_id"`
	Manga   	*Manga 			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manga,omitempty"`

	Timestamp
}