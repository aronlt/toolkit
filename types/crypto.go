package types

const Sha1 = "sha1"
const Md5 = "md5"

var CrcMask = uint32(0xffffffff) >> 6

func SetCRCMask(n int) {
	if n >= 32 || n <= 0 {
		return
	}
	CrcMask = uint32(0xffffffff) >> n
}
