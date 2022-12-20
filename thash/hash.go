package thash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"hash/crc32"
	"io"
	"os"

	"github.com/aronlt/toolkit/ttypes"
	"github.com/pkg/errors"
)

func hashFactory(hashTypes ...string) hash.Hash {
	if len(hashTypes) == 0 {
		return md5.New()
	}
	switch hashTypes[0] {
	case ttypes.Sha1:
		return sha1.New()
	case ttypes.Md5:
		return md5.New()
	default:
		return md5.New()
	}
}

// CRC 计算bytes的CRC值
func CRC(content []byte) uint32 {
	return crc32.ChecksumIEEE(content) & ttypes.CrcMask
}

// HashFile 计算文件的MD5值
func HashFile(path string, hashType ...string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "call HashFile fail, open filepath error")
	}

	defer file.Close()
	Hash := hashFactory(hashType...)
	if _, err := io.Copy(Hash, file); err != nil {
		return "", errors.Wrap(err, "call HashFile fail, copy file content error")
	}
	return hex.EncodeToString(Hash.Sum(nil)), nil
}

// HashBytes 计算bytes的CRC值
func HashBytes(content []byte, hashType ...string) string {
	Hash := hashFactory(hashType...)
	Hash.Write(content)
	return hex.EncodeToString(Hash.Sum(nil))
}
