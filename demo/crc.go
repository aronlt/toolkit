package demo

import "hash/crc32"

/*
* 计算crc
 */
func crc() {
	CastagnoliCrcTable := crc32.MakeTable(crc32.Castagnoli)
	var buf []byte
	crc32.Checksum(buf, CastagnoliCrcTable)
}
