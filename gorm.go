package toolkit

import (
	"reflect"
	"strings"
)

func SelectAll[T any](data T, prefix ...string) string {
	parts := make([]string, 0)
	t := reflect.TypeOf(data)
	count := t.NumField()
	for i := 0; i < count; i++ {
		tag := t.Field(i).Tag
		if v, ok := tag.Lookup("gorm"); ok {
			startIndex := strings.Index(v, "column:")
			if startIndex != -1 {
				v = v[startIndex+len("column:"):]
				column := v
				endIndex := strings.Index(v, ";")
				if endIndex != -1 {
					column = v[:endIndex]
				}
				if column == "-" {
					continue
				}
				if len(prefix) > 0 {
					parts = append(parts, prefix[0]+"."+"`"+column+"` ")
				} else {
					parts = append(parts, "`"+column+"` ")
				}
			}
		}
	}
	return strings.Join(parts, ",")
}
