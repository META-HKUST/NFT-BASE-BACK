package utils

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"time"
)

func GenEmailToken() string {
	// 获取当前时间的时间戳
	t := time.Now().Unix()

	// 生成一个MD5的哈希
	h := md5.New()

	// 将时间戳转换为byte，并写入哈希
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(t))
	h.Write([]byte(b))

	// 将字节流转化为16进制的字符串
	return hex.EncodeToString(h.Sum(nil))
}
