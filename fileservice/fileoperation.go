package fileservice

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"time"
)

const EXPIRETIME = time.Hour
type CosConfig struct {
	URL 		string		`mapstructure:"URL"`
	SecretID 	string		`mapstructure:"Secret_ID"`
	SecretKey	string		`mapstructure:"Secret_Key"`
}

var COSCONFIG CosConfig

func GetCosClient() *cos.Client{
	url, _ := url.Parse(COSCONFIG.URL)
	b := &cos.BaseURL{BucketURL: url}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  COSCONFIG.SecretID,
			SecretKey: COSCONFIG.SecretKey,
		},
	})
	return client
}

func Upload(key string,content io.Reader) (*cos.Response,error){
	client := GetCosClient()
	resp, err := client.Object.Put(context.Background(), key,content,nil)
	if err != nil {
		return resp,err
	}

	return resp,nil
}

func AccessFileObject(preSignedUrl *url.URL) {
	resp, err := http.Get(preSignedUrl.String())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer resp.Body.Close()
	fmt.Println(preSignedUrl.String())
	fmt.Printf("resp:%v\n", resp)
}

func GeneratePreSignedUrl(key string) (*url.URL,error){
	client := GetCosClient()
	presignedURL, err := client.Object.GetPresignedURL(context.Background(),http.MethodGet,key,COSCONFIG.SecretID,COSCONFIG.SecretKey,EXPIRETIME,nil)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(presignedURL)
	return presignedURL,err
}

func LoadConfig(path string,config *CosConfig) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(config)
	return err
}




