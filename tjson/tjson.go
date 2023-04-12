package tjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

var KeyNotFoundErr = errors.New("not found key")
var InvalidKeyErr = errors.New("invalid key")
var InvalidSliceIndexErr = errors.New("invalid slice index")

/*
 * 遍历json串中的对象
 * travel 支持按照字符串的模式遍历json，例如传入key1.key2.key3的模式，特别针对数组的情况
 * 你可以使用key1.key2[1].key3的模式传入，函数会帮你解析数组下标
 */
func travel(keys []string, rawEntryMap map[string]json.RawMessage, lastElemHandler func(val json.RawMessage, index int) (interface{}, error)) (interface{}, error) {
	if len(keys) == 0 {
		return nil, InvalidKeyErr
	}
	key := keys[0]
	keys = keys[1:]
	// 从rawMap中找到对应的项，所以这里要生成key字符串
	key = strings.TrimSpace(key)
	if key == "" {
		return nil, InvalidKeyErr
	}
	// 查看是否是数组模式
	index := -1
	every := false
	if key[len(key)-1] == ']' {
		boundary := false
		nums := bytes.NewBuffer([]byte{})
		for i := len(key) - 2; i >= 0; i-- {
			if key[i] == '[' {
				// 取出真正的key
				key = key[0:i]
				var err error
				if nums.String() == "*" {
					every = true
				} else {
					index, err = strconv.Atoi(nums.String())
					if err != nil {
						return nil, InvalidKeyErr
					}
					if index < 0 {
						return nil, InvalidSliceIndexErr
					}
				}
				boundary = true
				break
			} else if key[i] == ' ' {
				continue
			} else {
				nums.WriteByte(key[i])
			}
		}
		if boundary == false {
			return nil, InvalidKeyErr
		}
	}

	// 查看当前key的内容
	if val2, ok := rawEntryMap[key]; ok {
		// 如果已经遍历到最后一个elem，就直接调用处理函数进行处理
		if len(keys) == 0 {
			value, err := lastElemHandler(val2, index)
			return value, err
		}
		if every {
			var rawSliceMap []json.RawMessage
			if err := json.Unmarshal(val2, &rawSliceMap); err != nil {
				return nil, err
			}
			anySlice := make([]interface{}, 0)
			for i := 0; i < len(rawSliceMap); i++ {
				// 解析失败
				if err := json.Unmarshal(rawSliceMap[i], &rawEntryMap); err != nil {
					return nil, err
				}
				v, err := travel(keys, rawEntryMap, lastElemHandler)
				if err != nil {
					return nil, err
				}
				anySlice = append(anySlice, v)
			}
			return anySlice, nil
		}
		// 检查是否是数组
		if index >= 0 {
			var rawSliceMap []json.RawMessage
			if err := json.Unmarshal(val2, &rawSliceMap); err != nil {
				return nil, err
			}
			// 下标不对
			if len(rawSliceMap) <= index {
				return nil, InvalidSliceIndexErr
			}
			// 解析失败
			if err := json.Unmarshal(rawSliceMap[index], &rawEntryMap); err != nil {
				return nil, err
			}
			return travel(keys, rawEntryMap, lastElemHandler)
		}
		// 非数组
		if err := json.Unmarshal(val2, &rawEntryMap); err != nil {
			return nil, err
		}
		return travel(keys, rawEntryMap, lastElemHandler)
	} else {
		// 如果key不存在，就直接返回失败
		return nil, KeyNotFoundErr
	}
}

func GetString(content []byte, key string) (string, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return "", err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		if index < 0 {
			var strVal string
			if err := json.Unmarshal(val, &strVal); err != nil {
				return "", err
			}
			return strVal, nil
		} else {
			var strSliceVal []string
			if err := json.Unmarshal(val, &strSliceVal); err != nil {
				return "", err
			}
			if len(strSliceVal) <= index {
				return "", InvalidSliceIndexErr
			}
			return strSliceVal[index], nil
		}
	})
	if err != nil {
		return "", err
	} else {
		return val.(string), nil
	}
}

func GetStringSlice(content []byte, key string) ([]string, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return []string{}, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		var strSliceVal []string
		if err := json.Unmarshal(val, &strSliceVal); err != nil {
			return []string{}, err
		}
		return strSliceVal, nil
	})
	if err != nil {
		return []string{}, err
	} else {
		return val.([]string), nil
	}
}

func GetEverySlice(content []byte, key string) (interface{}, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return nil, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		return val, nil
	})
	if err != nil {
		return nil, err
	} else {
		return val, nil
	}
}

func GetStringWithDefault(content []byte, key string, defaultVal string) string {
	if val, err := GetString(content, key); err != nil {
		return defaultVal
	} else {
		return val
	}
}

func GetInt(content []byte, key string) (int, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return -1, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		if index < 0 {
			var intVal int
			if err := json.Unmarshal(val, &intVal); err != nil {
				return 0, err
			}
			return intVal, nil
		} else {
			var intSliceVal []int
			if err := json.Unmarshal(val, &intSliceVal); err != nil {
				return 0, err
			}
			if len(intSliceVal) <= index {
				return 0, InvalidSliceIndexErr
			}
			return intSliceVal[index], nil
		}
	})
	if err != nil {
		return 0, err
	} else {
		return val.(int), nil
	}
}

func GetIntSlice(content []byte, key string) ([]int, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return []int{}, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		var intSliceVal []int
		if err := json.Unmarshal(val, &intSliceVal); err != nil {
			return []int{}, err
		}
		return intSliceVal, nil
	})
	if err != nil {
		return []int{}, err
	} else {
		return val.([]int), nil
	}
}

func GetIntWithDefault(content []byte, key string, defaultVal int) int {
	if val, err := GetInt(content, key); err != nil {
		return defaultVal
	} else {
		return val
	}
}

func GetFloat(content []byte, key string) (float64, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return -1.0, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		if index < 0 {
			var floatVal float64
			if err := json.Unmarshal(val, &floatVal); err != nil {
				return 0.0, err
			}
			return floatVal, nil
		} else {
			var floatSliceVal []float64
			if err := json.Unmarshal(val, &floatSliceVal); err != nil {
				return 0, err
			}
			if len(floatSliceVal) <= index {
				return 0, InvalidSliceIndexErr
			}
			return floatSliceVal[index], nil
		}
	})
	if err != nil {
		return 0.0, err
	} else {
		return val.(float64), nil
	}
}

func GetFloatSlice(content []byte, key string) ([]float64, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return []float64{}, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		var floatSliceVal []float64
		if err := json.Unmarshal(val, &floatSliceVal); err != nil {
			return []float64{}, err
		}
		return floatSliceVal, nil
	})
	if err != nil {
		return []float64{}, err
	} else {
		return val.([]float64), nil
	}
}

func GetFloatWithDefault(content []byte, key string, defaultVal float64) float64 {
	if val, err := GetFloat(content, key); err != nil {
		return defaultVal
	} else {
		return val
	}
}

func GetBool(content []byte, key string) (bool, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return false, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		if index < 0 {
			var boolVal bool
			if err := json.Unmarshal(val, &boolVal); err != nil {
				return 0.0, err
			}
			return boolVal, nil
		} else {
			var boolSliceVal []bool
			if err := json.Unmarshal(val, &boolSliceVal); err != nil {
				return 0, err
			}
			if len(boolSliceVal) <= index {
				return 0, InvalidSliceIndexErr
			}
			return boolSliceVal[index], nil
		}
	})
	if err != nil {
		return false, err
	} else {
		return val.(bool), nil
	}
}

func GetBoolSlice(content []byte, key string) ([]bool, error) {
	// 先从缓存中取出
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return []bool{}, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		var boolSliceVal []bool
		if err := json.Unmarshal(val, &boolSliceVal); err != nil {
			return []bool{}, err
		}
		return boolSliceVal, nil
	})
	if err != nil {
		return []bool{}, err
	} else {
		return val.([]bool), nil
	}
}

func GetBoolWithDefault(content []byte, key string, defaultVal bool) bool {
	if val, err := GetBool(content, key); err != nil {
		return defaultVal
	} else {
		return val
	}
}

// GetRawMessage 返回原始片段，如果需要自己处理的话，可以自己处理
func GetRawMessage(content []byte, key string) (json.RawMessage, error) {
	rawEntryMap := make(map[string]json.RawMessage, 0)
	err := json.Unmarshal(content, &rawEntryMap)
	if err != nil {
		return json.RawMessage{}, err
	}
	keys := strings.Split(key, ".")
	val, err := travel(keys, rawEntryMap, func(val json.RawMessage, index int) (interface{}, error) {
		var rawVal json.RawMessage
		if err := json.Unmarshal(val, &rawVal); err != nil {
			return json.RawMessage{}, err
		}
		return rawVal, nil
	})
	if err != nil {
		return json.RawMessage{}, err
	} else {
		return val.(json.RawMessage), nil
	}
}
