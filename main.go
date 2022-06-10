package main

import (
	"NFT-BASE-BACK/config"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/router"
	"NFT-BASE-BACK/sdk"
	"NFT-BASE-BACK/sdk/pb"
	"NFT-BASE-BACK/sdk/service"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// @title HKUST-NFT
// @version 1.0
// @description HKUST-NFT Server API
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host unifit.ust.hk:8888
// @BasePath /api/v1
func startGinServer() {
	gin.DisableConsoleColor()
	f, _ := os.Create("/home/fabric_release/03_End/zwang/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// gin.DefaultWriter = io.MultiWriter(f)
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	model.InitDB(config)
	sdk.InitClient()
	router := router.InitRouter()
	router.Run(":8889")
}

func startGRPCServer() {
	server := service.NewFabricServer()
	grpcServer := grpc.NewServer()
	pb.RegisterFabricSDKServer(grpcServer, server)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9080))
	if err != nil {
		log.Fatal("cannot start grpc server:", err.Error())
	}
	grpcServer.Serve(listener)
}

func main() {
	serverType := flag.String("server", "gin", "choose gin or grpc server")
	flag.Parse()
	if strings.ToLower(*serverType) == "gin" {
		startGinServer()
	} else if strings.ToLower(*serverType) == "grpc" {
		startGRPCServer()
	}
}
