package dto

import (
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
)

type SeriCreateDTO struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Judul 			string    `json:"judul" form:"judul" binding:"required"`
	Sinopsis 		string    `json:"sinopsis" form:"sinopsis" binding:"required"`
	TahunTerbit 	string    `json:"tahun_terbit" form:"tahun_terbit" binding:"required"`
	Skor 			string    `json:"skor" form:"skor" binding:"required"`
	TotalPenilai 	string    `json:"total_penilai" form:"total_penilai" binding:"required"`
	TotalPembaca 	string    `json:"total_pembaca" form:"total_pembaca" binding:"required"`

	PenerbitID 		uuid.UUID   `gorm:"foreignKey" json:"penerbit_id" form:"penerbit_id" binding:"required"`
}

type SeriUpdateDTO struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Judul 			string    `json:"judul" form:"judul"`
	Sinopsis 		string    `json:"sinopsis" form:"sinopsis"`
	TahunTerbit 	string    `json:"tahun_terbit" form:"tahun_terbit"`
	Skor 			string    `json:"skor" form:"skor"`
	TotalPenilai 	string    `json:"total_penilai" form:"total_penilai"`
	TotalPembaca 	string    `json:"total_pembaca" form:"total_pembaca"`

	PenerbitID 		uuid.UUID   `gorm:"foreignKey" json:"penerbit_id" form:"penerbit_id"`
}

type SeriResponseDTO struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Judul 			string    `json:"judul" form:"judul"`
	Sinopsis		string	  `json:"sinopsis" form:"sinopsis"`
	TahunTerbit 	string    `json:"tahun_terbit" form:"tahun_terbit"`
	Skor 			string    `json:"skor" form:"skor"`
	TotalPenilai 	string    `json:"total_penilai" form:"total_penilai"`
	TotalPembaca 	string    `json:"total_pembaca" form:"total_pembaca"`

	PenerbitID 		uuid.UUID   `gorm:"foreignKey" json:"penerbit_id" form:"penerbit_id"`

	Manga			[]entity.Manga `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manga,omitempty"`
}
