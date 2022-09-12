package toolkit

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
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

// SecTimestamp 用于mysql的秒时间戳类型
type SecTimestamp int64

func NewSecTimestamp(t time.Time) SecTimestamp {
	return SecTimestamp(t.Unix())
}

func (s *SecTimestamp) Time() time.Time {
	return time.Unix(int64(*s), 0)
}

func (s *SecTimestamp) Scan(src interface{}) error {
	if s == nil {
		return nil
	}
	switch t := src.(type) {
	case time.Time:
		if t.IsZero() {
			*s = 0
		} else {
			*s = NewSecTimestamp(t)
		}
	case int64:
		*s = SecTimestamp(t)
	default:
		return fmt.Errorf("converting driver.Value type %T (%q) to a %T: invalid syntax", t, t, *s)
	}
	return nil
}

func (s *SecTimestamp) Value() (driver.Value, error) {
	return s.Time(), nil
}
