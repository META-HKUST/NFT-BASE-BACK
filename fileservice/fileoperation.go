package fileservice

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
)


type FileOperation struct {
	Url 		*url.URL
	SecretID 	string
	SecretKey	string
}

func NewFileOperation (url *url.URL,secretID,secretKey string,) *FileOperation {
	return &FileOperation{
		url,
		secretID,
		secretKey,
	}
}

func (fOperation *FileOperation) Upload(key string,content io.Reader) (*cos.Response,error){

	b := &cos.BaseURL{BucketURL: fOperation.Url}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  fOperation.SecretID,
			SecretKey: fOperation.SecretKey,
		},
	})


	resp, err := c.Object.Put(context.Background(), key,content,nil)
	if err != nil {
		return resp,err
	}
	return resp,nil
}


