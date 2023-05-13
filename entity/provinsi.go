package entity

type Provinsi struct {
	ID   string `gorm:"primary_key;not_null" json:"id"`
	Nama string `json:"nama"`
	Timestamp
}