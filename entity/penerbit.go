package entity

import "github.com/google/uuid"

type Penerbit struct {
	ID   uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Nama string    `json:"nama"`

	Timestamp
}