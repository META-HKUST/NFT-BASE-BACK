package main

import (
	"NFT-BASE-BACK/config"
	"NFT-BASE-BACK/fileservice"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/router"
	"NFT-BASE-BACK/sdk"
	"NFT-BASE-BACK/sdk/pb"
	"NFT-BASE-BACK/sdk/service"
	"flag"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// @title HKUST-NFT
// @version 1.0
// @description HKUST-NFT Server API
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host unifit.ust.hk:8889
// @BasePath /api/v2
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func startGinServer() {
	gin.DisableConsoleColor()
	log.Println(config.CONFIG.LogFilePATH)
	f, err := os.Create(config.CONFIG.LogFilePATH)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	model.InitDB(config.CONFIG)
	sdk.InitClient()
	httprouter := router.InitRouter()
	//err = router.RunTLS("unifit.ust.hk:"+config.CONFIG.GinPort, "config/ssl.pem", "config/ssl.key")
	//if err != nil {
	//	return
	//}
	err = httprouter.Run(":" + config.CONFIG.GinPort)
	if err != nil {
		return
	}
}

func startGRPCServer() {
	server := service.NewFabricServer()
	grpcServer := grpc.NewServer()
	pb.RegisterFabricSDKServer(grpcServer, server)
	listener, err := net.Listen("tcp", ":"+config.CONFIG.GrpcPort)
	if err != nil {
		log.Fatal("cannot start grpc server:", err.Error())
	}
	grpcServer.Serve(listener)
}

func main() {
	//load flag
	serverType := flag.String("server", "gin", "choose gin or grpc server")
	flag.Parse()

	//load config
	err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	err = fileservice.LoadConfig("./config", &fileservice.COSCONFIG)
	if err != nil {
		log.Fatal("cannot load cos and ipfs config", err)
	}

	if strings.ToLower(*serverType) == "gin" {
		startGinServer()
	} else if strings.ToLower(*serverType) == "grpc" {
		startGRPCServer()
	}
}
