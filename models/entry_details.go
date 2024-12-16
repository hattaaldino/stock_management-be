package models

type EntryDetail struct {
	EntryDetailID int    `gorm:"primaryKey" json:"entry_detail_id"`
	EntryID       string `json:"entry_id"`
	ItemCode      string `json:"item_code"`
	BatchID       string `json:"batch_id"`
	ExpiryDate    string `json:"expiry_date"`
	Qty           int    `json:"qty"`
}
