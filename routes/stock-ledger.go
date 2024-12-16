package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hattaaldino/stock_management-be/config"
	"github.com/hattaaldino/stock_management-be/controllers"
)

func GetAllStockLedger(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, gin.H{
				"result":  "FAILED",
				"message": "Request failed: " + message,
			})
		}
	}()

	stock_ledger_list, err := controllers.GetStockLedger(config.DB)
	if err != nil {
		panic("get stock ledger failed")
	} else if len(stock_ledger_list) <= 0 {
		panic("Empty stock ledger")
	}

	c.JSON(200, gin.H{
		"result":  "SUCCESS",
		"message": "Success get stock ledger",
		"data":    stock_ledger_list,
	})

}
