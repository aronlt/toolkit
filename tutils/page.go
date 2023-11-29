package tutils

// MemoryPage 内存分页函数
func MemoryPage[T any](data []T, offset int64, limit int64) []T {
	if offset < 0 || limit < 0 {
		return []T{}
	}
	if offset >= int64(len(data)) {
		return []T{}
	}
	if offset+limit >= int64(len(data)) {
		return data[offset:]
	}
	return data[offset : offset+limit]
}
