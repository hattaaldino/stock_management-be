package models

type Item struct {
	Code string `gorm:"primaryKey" json:"code"`
	Name string `json:"name"`
	UOM  string `json:"uom"`
}
