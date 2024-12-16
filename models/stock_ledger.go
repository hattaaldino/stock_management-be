package models

import "time"

type StockLedger struct {
	ItemCode     string    `gorm:"primaryKey" json:"item_code"`
	BatchID      string    `gorm:"primaryKey" json:"batch_id"`
	Tanggal      time.Time `json:"tanggal"`
	LastStock    int       `json:"last_stock"`
	QtyIn        int       `json:"qty_in"`
	QtyOut       int       `json:"qty_out"`
	CurrentStock int       `json:"current_stock"`
}
