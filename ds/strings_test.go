package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrHasSuffixInsensitive(t *testing.T) {
	a := "abceEfg"
	assert.True(t, StrHasSuffixInsensitive(a, "efg"))
}

func TestStrHasPrefixInsensitive(t *testing.T) {
	a := "AbceEfg"
	assert.True(t, StrHasPrefixInsensitive(a, "abc"))
}

func TestStrRemoveTail(t *testing.T) {
	a := "abcede"
	assert.Equal(t, StrRemoveTail(a, 3), "abc")
}

func TestStrRemoveHead(t *testing.T) {
	a := "abcede"
	assert.Equal(t, StrRemoveHead(a, 3), "ede")
}

func TestStrSplitNth(t *testing.T) {
	a := "a,b,c,d"
	assert.Equal(t, StrSplitNth(a, ",", -1, 2), "c")
}
