package utils

import (
	"fmt"
	"testing"
)

func TestAESCbcEncrypt (t *testing.T) {

	plain := "The text need 被to be encrypt."
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"

	cipherByte, err := Encrypt(plain, key)
	cipherByte = append(cipherByte, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %x\n", plain, cipherByte)
	// 解密
	plainText, err := Decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%x ==> %v\n", cipherByte, plainText)
}

func TestAes(t *testing.T)  {

	ciperText,err := AesEcpt.AesBase64Encrypt("nihaoni你好吗hhh")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ciperText)

	plainText,err := AesEcpt.AesBase64Decrypt(ciperText)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(plainText)
}