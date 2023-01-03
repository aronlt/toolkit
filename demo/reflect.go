package demo

import "fmt"

/*
* 判断interface的具体类型
 */
func interfaceToSpecific(itr interface{}) {
	switch itr.(type) {
	case string:
		fmt.Printf("%+v", itr.(string))
	case []byte:
		fmt.Printf("%+v", itr.(string))
	}
}
