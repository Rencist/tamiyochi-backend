package entity

import "github.com/google/uuid"

type UserKomentar struct {
	ID   		uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`

	UserID		uuid.UUID	`gorm:"foreignKey" json:"user_id"`
	User   		*User 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	
	KomentarID 	uuid.UUID	`gorm:"foreignKey" json:"komentar_id"`
	Komentar   	*Komentar 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"komentar,omitempty"`

	Timestamp
}