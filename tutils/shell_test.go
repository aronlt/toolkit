package tutils

import (
	"fmt"
	"testing"
)

func TestRunCmd(t *testing.T) {
	cmd := "ls"
	r := RunCmd(cmd, map[string]string{})
	fmt.Printf("%+v", r)
}
