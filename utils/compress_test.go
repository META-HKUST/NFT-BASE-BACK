package utils

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestCompressImg (t *testing.T) {
	for i := 10; i < 500; i += 25 {
		f,err:= os.OpenFile("./img/test.jpg",os.O_RDWR|os.O_CREATE, 0744)
		if err != nil {
			fmt.Println(err)
		}
		img,err := CompressImg(f,f.Name(),uint(i))
		if err != nil {
			log.Println(err)
		}
		fmt.Println(img)
	}
}