package entity

import "github.com/google/uuid"

type Komentar struct {
	ID  uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Isi string    `json:"isi"`

	Timestamp
}