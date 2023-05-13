package entity

import "github.com/google/uuid"

type SeriGenre struct {
	ID   	uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`

	SeriID 	uuid.UUID	`gorm:"foreignKey" json:"seri_id"`
	Seri   	*Seri 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seri,omitempty"`
	
	GenreID uuid.UUID	`gorm:"foreignKey" json:"genre_id"`
	Genre   *Genre 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"genre,omitempty"`

	Timestamp
}