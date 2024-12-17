package routes

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/hattaaldino/stock_management-be/config"
	"github.com/hattaaldino/stock_management-be/controllers"
	"github.com/hattaaldino/stock_management-be/models"
)

type Entry struct {
	Type         string               `json:"type"`
	EntryDetails []models.EntryDetail `json:"details"`
}

func AddStockEntry(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, gin.H{
				"result":  "FAILED",
				"message": "Request failed: " + message,
			})
		}
	}()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic("read body failed")
	}

	var newEntry Entry
	err = json.Unmarshal(body, &newEntry)
	if err != nil {
		panic("parsing body failed")
	}

	if newEntry.Type != "IN" && newEntry.Type != "OUT" {
		panic("wrong request data")
	}

	_, err = controllers.AddStockEntry(config.DB, newEntry.Type, newEntry.EntryDetails)
	if err != nil {
		panic("add stock entry failed")
	}

	c.JSON(200, gin.H{
		"result":  "SUCCESS",
		"message": "Success stock entry item",
	})

}

func GetAllEntry(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, gin.H{
				"result":  "FAILED",
				"message": "Request failed: " + message,
			})
		}
	}()

	stock_entry_list, err := controllers.GetStockEntry(config.DB)
	if err != nil {
		panic("get stock entry failed")
	} else if len(stock_entry_list) <= 0 {
		panic("Empty stock entry")
	}

	c.JSON(200, gin.H{
		"result":  "SUCCESS",
		"message": "Success get stock entry",
		"data":    stock_entry_list,
	})

}
