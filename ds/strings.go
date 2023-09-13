package ds

import "strings"

// StrHasSuffixInsensitive 判断str是以subStr结尾, subStr不关注大小写
func StrHasSuffixInsensitive(str string, subStr string) bool {
	return len(str) >= len(subStr) && strings.ToLower(str[len(str)-len(subStr):]) == strings.ToLower(subStr)
}

// StrHasPrefixInsensitive 判断str是以subStr开始, subStr不关注大小写
func StrHasPrefixInsensitive(str string, subStr string) bool {
	return len(str) >= len(subStr) && strings.ToLower(str[0:len(subStr)]) == strings.ToLower(subStr)
}

// StrHasContainInsensitive 大小写无关的包含判断
func StrHasContainInsensitive(str string, subStr string) bool {
	return len(str) >= len(subStr) && strings.Contains(strings.ToLower(str), strings.ToLower(subStr))
}

// StrRemoveTail 删除str末尾n个元素
func StrRemoveTail(s string, n int) string {
	if len(s) < n {
		return ""
	}
	return s[:len(s)-n]
}

// StrRemoveHead 删除str起始n个元素
func StrRemoveHead(s string, n int) string {
	if len(s) < n {
		return ""
	}
	return s[n:]
}

// StrSplitNth 字符串分割按照sep分割，约定total个，返回第nth个元素
// total == -1表示不关注total元素个数
// nth == -1表示返回最后一个元素
func StrSplitNth(str string, sep string, total int, nth int) string {
	values := strings.Split(str, sep)
	if len(values) == total || total == -1 {
		if nth == -1 {
			return SliceGetTail(values, "")
		}
		if len(values) < nth {
			return ""
		} else {
			return values[nth]
		}
	}
	return ""
}

// StrReverse 转置字符串
func StrReverse(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	SliceOpReverse(r)
	return string(r)
}
