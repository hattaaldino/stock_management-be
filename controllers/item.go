package controllers

import (
	"fmt"
	"strconv"

	"github.com/hattaaldino/stock_management-be/models"
	"gorm.io/gorm"
)

func AddItem(db *gorm.DB, name, uom string) (*models.Item, error) {
	var lastItem models.Item
	if err := db.Raw("SELECT code, name, uom FROM item ORDER BY code DESC").Scan(&lastItem).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	lastNumber := 0
	if lastItem.Code != "" {
		lastNumber, _ = strconv.Atoi(lastItem.Code[1:])
	}
	code := fmt.Sprintf("I%05d", lastNumber+1)

	item := &models.Item{Code: code, Name: name, UOM: uom}
	if err := db.Exec("INSERT INTO item (code, name, uom) VALUES (?, ?, ?)", item.Code, item.Name, item.UOM).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func GetItem(db *gorm.DB, code string) (*models.Item, error) {
	var item models.Item
	if err := db.Raw("SELECT code, name, uom FROM item WHERE code = ?", code).Scan(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func GetAllItem(db *gorm.DB) ([]models.Item, error) {
	var items []models.Item
	if err := db.Raw("SELECT code, name, uom FROM item").Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
