package demo

import (
	"bytes"
	"io"
	"strings"
)

// 不同类型转换为reader
func stringToReader(str string) io.Reader {
	return strings.NewReader(str)
}

func bytesToReader(bs []byte) io.Reader {
	return bytes.NewReader(bs)
}
