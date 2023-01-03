package demo

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

// 不同类型转换为reader
func stringToReader(str string) io.Reader {
	return strings.NewReader(str)
}
func bytesToReader(bs []byte) io.Reader {
	return bytes.NewReader(bs)
}
func fileToReader(fp *os.File) io.Reader {
	return bufio.NewReader(fp)
}
