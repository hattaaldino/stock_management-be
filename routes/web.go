package routes

import (
	"github.com/gin-gonic/gin"
)

func Regist(router *gin.Engine) {

	item := router.Group("/item")
	{
		item.GET("/list", GetAllItem)

		/*
			{
				"name":"Pin",
				"uom":"PCS"
			}
		*/
		item.POST("/add", AddItem)
	}

	stockEntry := router.Group("/stock-entry")
	{
		stockEntry.GET("/list", GetAllEntry)

		/*
			{
				"type":"IN",
				"details":[
					{
						"item_code":"025723960",
						"expiry_date":"2024-12-12",
						"qty":100
					}
				]
			}
		*/
		stockEntry.POST("/add", AddStockEntry)
	}

	stockLedger := router.Group("/stock-ledger")
	{
		stockLedger.GET("/list", GetAllStockLedger)
	}

}
