package nftstorage

import (
	"context"
	"fmt"
	openapiclient "github.com/nftstorage/go-client"
	"net/http"
	"os"
	"testing"
)


func TestNFTStorage(t *testing.T) {
	body,err := os.Open("")
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
		return
	}

	configuration := openapiclient.NewConfiguration()

	api_client := openapiclient.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), openapiclient.ContextAccessToken, "")
	resp, r, err := api_client.NFTStorageAPI.Store(ctx).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTStorageAPI.Store``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Store`: UploadResponse
	//https://ipfs.io/ipfs/XXXXX
	//https://nftstorage.link/ipfs/XXXXXX
	fmt.Fprintf(os.Stdout, "Response from `NFTStorageAPI.Store`: %v\n", resp)
	fmt.Printf("Cid from `NFTStorageAPI.Store`: %v\n", *resp.Value.Cid)
}

func TestNFTService_UploadImage(t *testing.T) {
	apikey := ""
	client := http.Client{}
	nftService := NewNFTService(apikey,&client)
	//str, err := nftService.UploadImage("")
}
