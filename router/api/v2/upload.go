package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/fileservice"
	"NFT-BASE-BACK/nftstorage"
	"NFT-BASE-BACK/service"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

var PictureMax int64 = 2 << 20
var MusicMax int64 = 10 << 20
var VideoMax int64 = 400 << 20

type CosResponse struct {
	Url          string `json:"url"`
	UrlSignature string `json:"url_signature"`
}

type CosAndIpfsResponse struct {
	UrlIpfs      string `json:"url_ipfs"`
	UrlCos       string `json:"url_cos"`
	UrlSignature string `json:"url_signature"`
}

// UploadToCos @Description  upload data: 上传数据到图片服务器cos
// @Tags         upload
// @param 		 data  formData  file  true    "任何数据"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /upload/cos [POST]
func UploadToCos(ctx *gin.Context) {
	resp := base.Response{}
	file, header, _ := ctx.Request.FormFile("data")

	// check file format
	_, fileType := GetFileType(header.Filename)
	if fileType == "" {
		ctx.JSON(http.StatusOK, resp.SetCode(base.FileTypeError))
		return
	}

	// check file size
	if CheckFileSize(file, fileType) == false {
		ctx.JSON(http.StatusOK, resp.SetCode(base.FileSizeError))
		return
	}

	name := fileservice.DIRECTORY + "/" + header.Filename

	//_, url, err := fileservice.Upload(name, file)
	Url, encryptUrl, err := service.Upload(name, file)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(1000, resp)
		return
	}
	//UrlSignature, err := utils.Encrypt(url.String(),fileservice.COSCONFIG.CryptoKey)
	//if err != nil {
	//	fmt.Println(err)
	//}
	resp.Code = 0
	resp.Msg = "Operation Succeed"
	resp.Data = CosResponse{
		Url:          Url,
		UrlSignature: encryptUrl,
	}
	ctx.JSON(http.StatusOK, resp)
}

// UploadToIpfs @Description  UploadToIpfs: 编辑用户的个人资料
// @Tags         upload
// @param 		 data  formData  file  true    "任何数据"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /upload/ipfs-and-cos [POST]
func UploadToIpfs(ctx *gin.Context) {
	resp := base.Response{}
	file, header, _ := ctx.Request.FormFile("data")

	// check file format
	_, fileType := GetFileType(header.Filename)
	if fileType == "" {
		ctx.JSON(http.StatusOK, base.FileTypeError)
	}

	// check file size
	if CheckFileSize(file, fileType) == false {
		ctx.JSON(http.StatusOK, base.FileSizeError)
	}

	name := fileservice.DIRECTORY + "/" + header.Filename
	Url, encryptUrl, _ := service.Upload(name, file)
	//apikey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDE0REE3N0E4Y0VFZWIwNmY2OTZEQUIzZjFCMkQzODZCZTRiMUNjOTkiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY1MTQ1MDIyNjQ1MiwibmFtZSI6Im5mdC1zdG9yYWdlLXRlc3QifQ.cua-DSWuivlAVRSVxzVOR6pwCaavf5VVifai4zUyG9g"
	client := http.Client{}
	nftService := nftstorage.NewNFTService(fileservice.COSCONFIG.ApiKey, &client)

	fileReader, _ := header.Open()
	ipfsResp, _ := nftService.Upload(fileReader, "video")

	//UrlSignature, err := utils.Encrypt(url.String(),fileservice.COSCONFIG.CryptoKey)
	//if err != nil {
	//	fmt.Println(err)
	//}

	fileResp := CosAndIpfsResponse{
		ipfsResp.Value.Cid,
		Url,
		encryptUrl,
	}
	resp.Code = 0
	resp.Msg = "Operation Succeed"
	resp.Data = fileResp
	ctx.JSON(http.StatusOK, resp)
}

func GetFileType(filename string) (string, string) {
	if strings.Contains(filename, ".jpg") == true {
		return "jpg", "picture"
	}
	if strings.Contains(filename, ".png") == true {
		return "png", "picture"
	}
	if strings.Contains(filename, ".jpeg") == true {
		return "jpeg", "picture"
	}
	if strings.Contains(filename, ".gif") == true {
		return "gif", "picture"
	}
	if strings.Contains(filename, ".jfif") == true {
		return "jfif", "picture"
	}
	if strings.Contains(filename, ".webp") == true {
		return "webp", "picture"
	}
	if strings.Contains(filename, ".mp3") == true {
		return "mp3", "music"
	}
	if strings.Contains(filename, ".flac") == true {
		return "flac", "music"
	}
	if strings.Contains(filename, ".mp4") == true {
		return "mp4", "video"
	}
	if strings.Contains(filename, ".avi") == true {
		return "avi", "video"
	}
	return "", ""
}

func CheckFileSize(file multipart.File, filetype string) bool {

	var filesMax int64
	if filetype == "picture" {
		filesMax = PictureMax
	} else if filetype == "music" {
		filesMax = MusicMax
	} else if filetype == "video" {
		filesMax = VideoMax
	} else {
		log.Println("Error: could not resolve file type")
	}

	var buf bytes.Buffer
	n, err := io.CopyN(&buf, file, filesMax+1)
	if err != nil && err != io.EOF {
		log.Println(err)
		return false
	}

	filesMax -= n
	fmt.Println("filesMax: ", filesMax)
	if filesMax < 0 {
		return false
	}
	return true
}
