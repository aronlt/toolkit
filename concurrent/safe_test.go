package concurrent

import (
	"fmt"
	"testing"
)

func TestRunSafe(t *testing.T) {
	handler := func() {
		panic("panic")
	}
	errHandler := func(err any) {
		fmt.Println(err)
	}
	RunSafe(handler, errHandler)

	t.Log("run success")
}
