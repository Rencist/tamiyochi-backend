package entity

import "github.com/google/uuid"

type Seri struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Judul 			string    `json:"judul"`
	Sinopsis 		string    `json:"sinopsis"`
	TahunTerbit 	string    `json:"tahun_terbit"`
	Skor 			string    `json:"skor"`
	TotalPenilai 	string    `json:"total_penilai"`
	TotalPembaca 	string    `json:"total_pembaca"`

	PenerbitID 		uuid.UUID   `gorm:"foreignKey" json:"penerbit_id"`
	Penerbit   		*Penerbit 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"penerbit,omitempty"`

	Mangas 			[]Manga 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manga,omitempty"`

	Timestamp
}