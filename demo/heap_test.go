package demo

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	words := []string{
		"a",
		"a",
		"a",
		"a",
		"b",
		"b",
		"c",
		"c",
		"c",
		"c",
		"c",
		"c",
		"c",
		"c",
	}
	res := TopKFrequent(words, 2)
	fmt.Println(res)
}
