package demo

import (
	"crypto/aes"
	"crypto/rand"
)

/*
* 随机数生成器
 */
func random() {
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)
}
