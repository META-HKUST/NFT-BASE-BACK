package main

import (
	"NFT-BASE-BACK/api"
	"NFT-BASE-BACK/db"
	"log"

	"NFT-BASE-BACK/util"
)

// @title HKUST-NFT
// @version 1.0
// @description HKUST-NFT Server API
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	store := db.NewStore(config.DBDriver)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	url := "0.0.0.0:8080"
	err = server.Start(url)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
