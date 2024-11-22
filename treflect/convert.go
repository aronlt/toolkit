package treflect

import (
	"reflect"
	"strconv"
)

func MustToFloat64(o interface{}) float64 {
	v, _ := ToFloat64(o)
	return v
}

func indirect(o interface{}) (reflect.Value, interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr {
		return v, o
	}
	v = v.Elem()
	return v, v.Interface()
}

func MustToString(o interface{}) string {
	v, _ := ToString(o)
	return v
}

func ToString(o interface{}) (string, bool) {
	if v, ok := o.(string); ok {
		return v, true
	}
	if v, ok := o.([]byte); ok {
		return string(v), true
	}
	v, o := indirect(o)
	if !v.IsValid() {
		return "", false
	}
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), true
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), true
	case reflect.String:
		return v.String(), true
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return string(v.Bytes()), true
		}
	case reflect.Struct:
		if f, ok := o.(reflect.Value); ok {
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			switch f.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return strconv.FormatInt(f.Int(), 10), true
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return strconv.FormatUint(f.Uint(), 10), true
			case reflect.Float32, reflect.Float64:
				return strconv.FormatFloat(f.Float(), 'f', -1, 64), true
			case reflect.String:
				return f.String(), true
			case reflect.Slice:
				if f.Type().Elem().Kind() == reflect.Uint8 {
					return string(f.Bytes()), true
				}
			default:
				return "", false
			}
		}
	default:
		return "", false
	}
	return "", false
}

func ToFloat64(o interface{}) (float64, bool) {
	if i, ok := o.(float64); ok {
		return i, true
	}
	if i, ok := o.(float32); ok {
		return float64(i), true
	}

	v, o := indirect(o)
	if !v.IsValid() {
		return 0, false
	}
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.String:
		if f, err := strconv.ParseFloat(v.String(), 64); err == nil {
			return f, true
		}
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			if f, err := strconv.ParseFloat(string(v.Bytes()), 64); err == nil {
				return f, true
			}
		}
	case reflect.Struct:
		if f, ok := o.(reflect.Value); ok {
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			switch f.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return float64(f.Int()), true
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return float64(f.Uint()), true
			case reflect.Float32, reflect.Float64:
				return f.Float(), true
			case reflect.String:
				if r, err := strconv.ParseFloat(f.String(), 64); err == nil {
					return r, true
				}
			case reflect.Slice:
				if f.Type().Elem().Kind() == reflect.Uint8 {
					if r, err := strconv.ParseFloat(string(f.Bytes()), 64); err == nil {
						return r, true
					}
				}
			default:
				return 0, false
			}
		}
	default:
		return 0, false
	}
	return 0, false
}

func MustToInt64(o interface{}) int64 {
	v, _ := ToInt64(o)
	return v
}

func ToInt64(o interface{}) (int64, bool) {
	if i, ok := o.(int64); ok {
		return i, true
	}
	if i, ok := o.(int); ok {
		return int64(i), true
	}

	v, o := indirect(o)
	if !v.IsValid() {
		return 0, false
	}
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint()), true
	case reflect.String:
		if f, err := strconv.ParseInt(v.String(), 10, 64); err == nil {
			return f, true
		}
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			if f, err := strconv.ParseInt(string(v.Bytes()), 10, 64); err == nil {
				return f, true
			}
		}
	case reflect.Struct:
		if f, ok := o.(reflect.Value); ok {
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			switch f.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return f.Int(), true
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return int64(f.Uint()), true
			case reflect.String:
				if r, err := strconv.ParseInt(f.String(), 10, 64); err == nil {
					return r, true
				}
			case reflect.Slice:
				if f.Type().Elem().Kind() == reflect.Uint8 {
					if r, err := strconv.ParseInt(string(f.Bytes()), 10, 64); err == nil {
						return r, true
					}
				}
			default:
				return 0, false
			}
		}
	default:
		return 0, false
	}
	return 0, false
}
