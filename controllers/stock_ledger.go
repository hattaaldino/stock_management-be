package controllers

import (
	"github.com/hattaaldino/stock_management-be/models"
	"gorm.io/gorm"
)

func GetStockLedger(db *gorm.DB) ([]models.StockLedger, error) {
	var ledger []models.StockLedger
	if err := db.Raw("SELECT item_code, batch_id, tanggal, last_stock, qty_in, qty_out, current_stock FROM stock_ledger ").Scan(&ledger).Error; err != nil {
		return nil, err
	}
	return ledger, nil
}
