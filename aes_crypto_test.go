package toolkit

// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*"

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	content          = []byte("pibigstar")
	content_16, _    = base64.StdEncoding.DecodeString("v1jqsGHId/H8onlVHR8Vaw==")
	content_24, _    = base64.StdEncoding.DecodeString("0TXOaj5KMoLhNWmJ3lxY1A==")
	content_32, _    = base64.StdEncoding.DecodeString("qM/Waw1kkWhrwzek24rCSA==")
	content_16_iv, _ = base64.StdEncoding.DecodeString("DqQUXiHgW/XFb6Qs98+hrA==")
	content_32_iv, _ = base64.StdEncoding.DecodeString("ZuLgAOii+lrD5KJoQ7yQ8Q==")
	// iv 长度必须等于blockSize，只能为16
	iv         = []byte("Hello My GoFrame")
	key_16     = []byte("1234567891234567")
	key_17     = []byte("12345678912345670")
	key_24     = []byte("123456789123456789123456")
	key_32     = []byte("12345678912345678912345678912345")
	keys       = []byte("12345678912345678912345678912346")
	key_32_err = []byte("1234567891234567891234567891234 ")

	// cfb模式blockSize补位长度, add by zseeker
	padding_size      = 16 - len(content)
	content_16_cfb, _ = base64.StdEncoding.DecodeString("oSmget3aBDT1nJnBp8u6kA==")
)

func TestEncrypt(t *testing.T) {
	data, err := Encrypt(content, key_16)
	assert.Nil(t, err)
	assert.Equal(t, data, content_16)

	data, err = Encrypt(content, key_24)
	assert.Nil(t, err)
	assert.Equal(t, data, content_24)

	data, err = Encrypt(content, key_32)
	assert.Nil(t, err)
	assert.Equal(t, data, content_32)

	data, err = Encrypt(content, key_16, iv)
	assert.Nil(t, err)
	assert.Equal(t, data, content_16_iv)

	data, err = Encrypt(content, key_32, iv)
	assert.Nil(t, err)
	assert.Equal(t, data, content_32_iv)
}

func TestDecrypt(t *testing.T) {
	decrypt, err := Decrypt(content_16, key_16)
	assert.Nil(t, err)
	assert.Equal(t, decrypt, content)

	decrypt, err = Decrypt(content_24, key_24)
	assert.Nil(t, err)
	assert.Equal(t, decrypt, content)

	decrypt, err = Decrypt(content_32, key_32)
	assert.Nil(t, err)
	assert.Equal(t, decrypt, content)

	decrypt, err = Decrypt(content_16_iv, key_16, iv)
	assert.Nil(t, err)
	assert.Equal(t, decrypt, content)

	decrypt, err = Decrypt(content_32_iv, key_32, iv)
	assert.Nil(t, err)
	assert.Equal(t, decrypt, content)

	decrypt, err = Decrypt(content_32_iv, keys, iv)
	assert.Equal(t, err.Error(), "invalid padding")
}

func TestPKCS5UnPaddingErr(t *testing.T) {
	_, err := PKCS5UnPadding(content, 0)
	assert.NotNil(t, err)

	// PKCS5UnPadding src len zero
	_, err = PKCS5UnPadding([]byte(""), 16)
	assert.NotNil(t, err)

	// PKCS5UnPadding src len > blockSize
	_, err = PKCS5UnPadding(key_17, 16)
	assert.NotNil(t, err)

	// PKCS5UnPadding src len > blockSize
	_, err = PKCS5UnPadding(key_32_err, 32)
	assert.NotNil(t, err)
}

func TestEncryptCFB(t *testing.T) {
	var padding int = 0
	data, err := EncryptCFB(content, key_16, &padding, iv)
	assert.Nil(t, err)
	assert.Equal(t, padding, padding_size)
	assert.Equal(t, data, content_16_cfb)
}

func TestDecryptCFB(t *testing.T) {
	decrypt, err := DecryptCFB(content_16_cfb, key_16, padding_size, iv)
	assert.Nil(t, err)
	assert.Equal(t, decrypt, content)
}
