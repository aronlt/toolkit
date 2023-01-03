package demo

import "io"

/*
* 读取固定数量的数据到字节数组中
 */
func readFull() {
	var lenCrcBuf [8]byte
	var r io.Reader
	io.ReadFull(r, lenCrcBuf[:])
}
