package tutils

import (
	"encoding/json"
	"fmt"
)

func MustMarshal(value any) string {
	str, err := json.MarshalIndent(value, "", " ")
	if err != nil {
		return fmt.Sprintf("%#v", value)
	}
	return string(str)
}
