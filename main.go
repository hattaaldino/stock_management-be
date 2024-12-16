package main

import (
	"fmt"
	"os"

	"github.com/hattaaldino/stock_management-be/config"
	"github.com/hattaaldino/stock_management-be/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	godotenv.Load()

	port := os.Getenv("server_port")

	if port == "" {
		port = "8080"
		fmt.Printf("Defaulting to port %s\n", port)
	}

	err = config.InitDB()
	if err != nil {
		fmt.Printf("DB initialization error: %s\n", err.Error())
		os.Exit(0)
	} else {
		fmt.Println("DB initialization success!")
	}

	r := gin.Default()

	routes.Regist(r)

	err = r.Run(":" + port)
	if err != nil {
		fmt.Printf("Start server error: %s\n", err.Error())
	} else {
		fmt.Printf("Listening on port %s\n", port)
	}

}
