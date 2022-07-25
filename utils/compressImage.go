package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func CompressImg(file io.Reader,filename string, hight uint) ([]byte,error) {
	var err error
	fmt.Println(filename)
	buf := bytes.Buffer{}
	reg, _ := regexp.Compile(`^.*\.((png)|(jpg))$`)
	if !reg.MatchString(filename) {
		err = errors.New("%s is not a .png or .jpg file")
		log.Println(err)
		return nil,err
	}

	var img image.Image
	switch {
	case strings.HasSuffix(filename, ".png"):
		if img, err = png.Decode(file); err != nil {
			log.Println(err)
			return nil,err
		}
	case strings.HasSuffix(filename, ".jpg"):
		if img, err = jpeg.Decode(file); err != nil {
			log.Println(err)
			return nil,err
		}
		rec := img.Bounds()
		fmt.Printf("r,c = [%v,%v]\n", rec.Max,rec.Min)
	default:
		err = fmt.Errorf("Images %s name not right!", filename)
		log.Println(err)
		return nil,err
	}
	resizeImg := resize.Resize(hight, 0, img, resize.Lanczos3)
	log.Println("The picture was successfully compressed")
	jpeg.Encode(&buf,resizeImg,nil)
	return buf.Bytes(),nil
}

//create a file name for the iamges that after resize
func newName(name string, size int) string {
	dir, file := filepath.Split(name)
	return fmt.Sprintf("%s_%d%s", dir, size, file)
}

