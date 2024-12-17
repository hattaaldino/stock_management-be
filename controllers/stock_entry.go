package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hattaaldino/stock_management-be/models"
	"gorm.io/gorm"
)

func AddStockEntry(db *gorm.DB, entryType string, details []models.EntryDetail) (*models.Entry, error) {

	var lastEntry models.Entry
	if err := db.Raw("SELECT id, tanggal, type FROM entry ORDER BY id DESC").Scan(&lastEntry).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	lastNumber := 0
	if lastEntry.ID != "" {
		lastNumber, _ = strconv.Atoi(lastEntry.ID[2:])
	}
	entryID := fmt.Sprintf("SE%03d", lastNumber+1)

	entry := &models.Entry{
		ID:      entryID,
		Tanggal: time.Now().Format("2006-01-02"),
		Type:    entryType,
	}
	if err := db.Exec("INSERT INTO entry (id, tanggal, type) VALUES (?, ?, ?)", entry.ID, entry.Tanggal, entry.Type).Error; err != nil {
		return nil, err
	}

	// Add Details and Related Records
	for _, detail := range details {
		// Generate Batch ID for IN type
		if entryType == "IN" {
			var lastBatch models.Batch
			if err := db.Raw("SELECT id, item_code, expiry_date FROM batch ORDER BY id DESC").Scan(&lastBatch).Error; err != nil && err != gorm.ErrRecordNotFound {
				return nil, err
			}
			lastNumber := 0
			if lastBatch.ID != "" {
				lastNumber, _ = strconv.Atoi(lastEntry.ID[2:])
			}
			batchID := fmt.Sprintf("B%03d", lastNumber+1)

			batch := models.Batch{ID: batchID, ItemCode: detail.ItemCode, ExpiryDate: detail.ExpiryDate}
			if err := db.Exec("INSERT INTO batch (id, item_code, expiry_date) VALUES (?, ?, ?)", batch.ID, batch.ItemCode, batch.ExpiryDate).Error; err != nil {
				return nil, err
			}

			detail.BatchID = batchID
		} else {
			var lastMatchedBatch models.Batch
			if err := db.Raw("SELECT id, item_code, expiry_date FROM batch WHERE item_code = ? AND expiry_date = ?", detail.ItemCode, detail.ExpiryDate).Scan(&lastMatchedBatch).Error; err != nil && err != gorm.ErrRecordNotFound {
				return nil,
					err
			}

			detail.BatchID = lastMatchedBatch.ID
		}

		// Generate EntryDetail ID
		var lastEntryDetail models.EntryDetail
		lastDetailID := 0

		if err := db.Raw("SELECT entry_detail_id, entry_id, item_code, batch_id, expiry_date, qty FROM entry_detail ORDER BY entry_detail_id DESC").Scan(&lastEntryDetail).Error; err != nil && err != gorm.ErrRecordNotFound {
			return nil,
				err
		}

		lastDetailID = lastEntryDetail.EntryDetailID
		detail.EntryDetailID = lastDetailID + 1
		detail.EntryID = entry.ID

		// Insert Entry Detail
		if err := db.Exec("INSERT INTO entry_detail (entry_detail_id, entry_id, item_code, batch_id, expiry_date, qty) VALUES (?, ?, ?, ?, ?, ?)", detail.EntryDetailID, detail.EntryID, detail.ItemCode, detail.BatchID, detail.ExpiryDate, detail.Qty).Error; err != nil {
			return nil, err
		}

		// Update Stock Ledger
		var ledger models.StockLedger
		if err := db.Raw("SELECT item_code, batch_id, tanggal, last_stock, qty_in, qty_out, current_stock FROM stock_ledger WHERE item_code = ? AND batch_id = ?", detail.ItemCode, detail.BatchID).Scan(&ledger).Error; err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		lastStock := 0
		if ledger.ItemCode != "" {
			lastStock = ledger.CurrentStock
		}

		ledger = models.StockLedger{
			ItemCode:     detail.ItemCode,
			BatchID:      detail.BatchID,
			Tanggal:      time.Now(),
			LastStock:    lastStock,
			QtyIn:        0,
			QtyOut:       0,
			CurrentStock: lastStock,
		}

		if entryType == "IN" {
			ledger.QtyIn = detail.Qty
			ledger.CurrentStock += detail.Qty
		} else if entryType == "OUT" {
			ledger.QtyOut = detail.Qty
			ledger.CurrentStock -= detail.Qty
		}

		if err := db.Exec("INSERT INTO stock_ledger (item_code, batch_id, tanggal, last_stock, qty_in, qty_out, current_stock) VALUES (?, ?, ?, ?, ?, ?, ?)", ledger.ItemCode, ledger.BatchID, ledger.Tanggal, ledger.LastStock, ledger.QtyIn, ledger.QtyOut, ledger.CurrentStock).Error; err != nil {
			return nil, err
		}
	}

	return entry, nil
}

func GetStockEntry(db *gorm.DB) ([]models.Entry, error) {
	var entry []models.Entry
	if err := db.Raw("SELECT id, tanggal, type FROM entry").Scan(&entry).Error; err != nil {
		return nil, err
	}

	// var details []models.EntryDetail
	// if err := db.Find(&details, "entry_id = ?", entryID).Error; err != nil {
	// 	return nil, nil, err
	// }

	return entry, nil
}
