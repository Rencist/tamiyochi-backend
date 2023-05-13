package entity

import "github.com/google/uuid"

type PenulisSeri struct {
	ID   		uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`

	PenulisID 	uuid.UUID	`gorm:"foreignKey" json:"penulis_id"`
	Penulis   	*Penulis 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"penulis,omitempty"`
	
	SeriID 		uuid.UUID	`gorm:"foreignKey" json:"seri_id"`
	Seri   		*Seri 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seri,omitempty"`

	Timestamp
}