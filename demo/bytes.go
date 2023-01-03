package demo

import "bytes"

func equal(bs1 []byte, bs2 []byte) {
	bytes.Equal(bs1, bs2)
}
