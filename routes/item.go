package routes

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/hattaaldino/stock_management-be/config"
	"github.com/hattaaldino/stock_management-be/controllers"
)

type Item struct {
	Name string `json:"name"`
	UOM  string `json:"uom"`
}

func AddItem(c *gin.Context) {

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

	var newItem Item
	err = json.Unmarshal(body, &newItem)
	if err != nil {
		panic("parsing body failed")
	}

	_, err = controllers.AddItem(config.DB, newItem.Name, newItem.UOM)
	if err != nil {
		panic("add item failed")
	}

	c.JSON(200, gin.H{
		"result":  "SUCCESS",
		"message": "Success add item",
	})

}

func GetAllItem(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, gin.H{
				"result":  "FAILED",
				"message": "Request failed: " + message,
			})
		}
	}()

	item_list, err := controllers.GetAllItem(config.DB)
	if err != nil {
		panic("get item data failed")
	} else if len(item_list) <= 0 {
		panic("Empty item data")
	}

	c.JSON(200, gin.H{
		"result":  "SUCCESS",
		"message": "Success get item list",
		"data":    item_list,
	})

}
