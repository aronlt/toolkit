package tutils

import (
	"math/rand"
	"time"

	"github.com/aronlt/toolkit/ds"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

// RandStringBytesMask 获取N个随机字符串
func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func RandPick[T comparable](data []T, blacklist ds.BuiltinSet[T]) (T, bool) {
	if len(data) == 0 {
		var e T
		return e, false
	}
	n := rand.Intn(len(data))
	if !blacklist.Has(data[n]) {
		return data[n], true
	}
	i := (n + 1) % len(data)
	for ; i != n; i = (i + 1) % len(data) {
		if !blacklist.Has(data[i]) {
			return data[i], true
		}
	}
	var e T
	return e, false
}
