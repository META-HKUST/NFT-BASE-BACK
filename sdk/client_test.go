package sdk

import (
	"NFT-BASE-BACK/config"
	"NFT-BASE-BACK/sdk/pb"
	"NFT-BASE-BACK/sdk/service"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	// startTestServer(t)
	InitClient()

	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal(err)
	}

	// response, err := client.OwnerOf(context.Background(), &pb.OwnerOfRequest{
	// 	TokenId:  "1",
	// 	Username: "zzding",
	// })
	// x509, _ := service.Enroll("zzding")
	// fmt.Printf(response.GetOwner())
	// require.NoError(t, err)
	// require.EqualValues(t, x509, response.GetOwner())
	// var wg sync.WaitGroup
	// wg.Add(1)
	for i := 0; i < 1; i++ {
		// go func() {
		ret, err := Client.PublicMint(context.Background(), &pb.PublicMintRequest{
			Username: "zzding",
		})
		log.Printf(err.Error())
		log.Printf(ret.GetTokenId())
		// wg.Done()
		// }()
	}
	// wg.Wait()
	// response2, err := client.PublicMint(context.Background(), &pb.PublicMintRequest{
	// 	Username: "zzding",
	// })
	// require.NoError(t, err)
	// fmt.Printf(response2.GetTokenId())
}

func startTestServer(t *testing.T) {
	server := service.NewFabricServer()
	grpcServer := grpc.NewServer()
	pb.RegisterFabricSDKServer(grpcServer, server)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9080))
	require.NoError(t, err)
	go grpcServer.Serve(listener)
}
