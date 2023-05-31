package entity

type Seri struct {
	ID           int       `gorm:"primary_key;not_null" json:"id"`
	Judul        string    `json:"judul"`
	Sinopsis     string    `json:"sinopsis"`
	TahunTerbit  string    `json:"tahun_terbit"`
	Skor         string    `json:"skor"`
	TotalPenilai string    `json:"total_penilai"`
	TotalPembaca string    `json:"total_pembaca"`
	Foto         string    `json:"foto"`
	PenerbitID   int       `gorm:"foreignKey" json:"penerbit_id"`
	Penerbit     *Penerbit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"penerbit,omitempty"`

	Mangas []Manga `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manga,omitempty"`

	Timestamp
}