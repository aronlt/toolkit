package demo

import "fmt"

// interface to specific
func interfaceToSpecific(itr interface{}) {
	switch itr.(type) {
	case string:
		fmt.Printf("%+v", itr.(string))
	case []byte:
		fmt.Printf("%+v", itr.(string))
	}
}
