package main

import (
	"NFT-BASE-BACK/config"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/router"
	"log"
)

// @title HKUST-NFT
// @version 1.0
// @description HKUST-NFT Server API
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://pascal.idea.ust.hk:8888
// @BasePath /api/v1
func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	model.InitDB(config)
	router := router.InitRouter()
	router.Run(":8888")
}
