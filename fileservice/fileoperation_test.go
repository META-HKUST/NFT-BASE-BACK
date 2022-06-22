package fileservice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUpload (t *testing.T)  {

	err := LoadConfig("../config",&COSCONFIG)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(COSCONFIG)
	name := "unifit/文件名"

	f := strings.NewReader("test")
	resp,err := Upload(name,f)
	//resp, err := client.Object.Put(context.Background(), name, f, nil)
	if err != nil {
	}
	fmt.Printf("%s\n", resp.Status)
}


func TestGeneratePreSignedUrl (t *testing.T)  {
	//var cosConfig CosConfig
	err := LoadConfig("../config",&COSCONFIG)
	if err !=nil{
		fmt.Println(err)
	}
	name := "unifit/文件名"
	presignedURL,err := GeneratePreSignedUrl(name)
	if err != nil {
		fmt.Println(err)
	}
	resp2, err := http.Get(presignedURL.String())
	if err != nil {
		fmt.Println(err)
	}
	bs2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(bs2))
	resp2.Body.Close()
}

