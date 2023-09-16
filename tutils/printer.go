package tutils

import "encoding/json"

func MustMarshal[T any](value T) string {
	content, _ := json.MarshalIndent(value, "", " ")
	return string(content)
}
