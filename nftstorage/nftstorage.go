package nftstorage

import (
	"NFT-BASE-BACK/nftstorage/nftclient"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const api = "https://api.nft.storage"

type NFTService struct {
	ApiKey 		string
	Client *http.Client
}

type RequestBody struct {
	Content 	string 		`json:"content"`
	ReadNum 	int32		`json:"read_num"`
	FileList	[]string	`json:"file_list"`
}
type Value struct {
	Cid     string
	Size    int64  `json:",omitempty"`
	Created string `json:",omitempty"`
	Type    string `json:",omitempty"`
}

type Err struct {
	Name, Message string
}

type Response struct {
	Ok    bool
	Value Value
	Error Err
}

func NewNFTService (apikey string,client *http.Client) *NFTService {

	return &NFTService{
		apikey,
		client,
	}
}


// UploadImage upload content to NFTStorage by given byte slice, it returns an IPFS hash and an error.
func (nft *NFTService) Upload(r io.Reader,fileType string) (Response, error) {
	endpoint := api + "/upload"
	req, err := http.NewRequest(http.MethodPost, endpoint, r)
	if err != nil {
		return Response{}, err
	}

	req.Header.Add("Authorization", "Bearer "+nft.ApiKey)
	client := nftclient.NewClient(nft.Client)
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var res Response
	if err := json.Unmarshal(data, &res); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			return Response{}, fmt.Errorf("json syntax error at byte offset %d", e.Offset)
		}
		return Response{}, err
	}
	res.Value.Type = fileType
	return res, nil
}









