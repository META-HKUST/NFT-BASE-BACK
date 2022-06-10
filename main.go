package main

import (
	"NFT-BASE-BACK/config"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/router"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// @title HKUST-NFT
// @version 1.0
// @description HKUST-NFT Server API
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host unifit.ust.hk:8888
// @BasePath /api/v1
func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	model.InitDB(config)
	router := router.InitRouter()
	router.Run(":8889")
}
