package models

type Entry struct {
	ID      string `gorm:"primaryKey" json:"id"`
	Tanggal string `json:"tanggal"`
	Type    string `json:"type"`
}
