package models

type Batch struct {
	ID         string `gorm:"primaryKey" json:"id"`
	ItemCode   string `json:"item_code"`
	ExpiryDate string `json:"expiry_date"`
}
