package toolkit

// ReverseSlice 转置切片
func ReverseSlice[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// UniqueSlice 去重切片
func UniqueSlice[T comparable](data []T) []T {
	record := make(map[T]struct{}, len(data))
	result := make([]T, 0, len(data))
	for i := 0; i < len(data); i++ {
		if _, ok := record[data[i]]; !ok {
			record[data[i]] = struct{}{}
			result = append(result, data[i])
		}
	}
	return result
}

// CopySlice 复制切片
func CopySlice[T any](data []T) []T {
	dst := make([]T, len(data))
	copy(dst, data)
	return dst
}
