package sdk

import (
	"NFT-BASE-BACK/sdk/pb"
	"log"

	"google.golang.org/grpc"
)

var Client pb.FabricSDKClient

func InitClient() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot creat grpc fabric sdk client")
	}
	Client = pb.NewFabricSDKClient(conn)
}
