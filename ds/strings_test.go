package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrIsNumber(t *testing.T) {
	a := "abcdef"
	assert.False(t, StrIsNumber(a))
	a = "100111"
	assert.True(t, StrIsNumber(a))
	a = "OX000111"
	assert.False(t, StrIsNumber(a))
}

func TestStrRuneOmit(t *testing.T) {
	a := "abcdef"
	assert.Equal(t, StrRuneOmit(a, 3), "abc...")
	a = "你好，这是一个测试程序"
	assert.Equal(t, StrRuneOmit(a, 3), "你好，...")
	a = "你好a，这是一个测试程序"
	assert.Equal(t, StrRuneOmit(a, 3), "你好a...")
	assert.Equal(t, StrRuneOmit(a, 3, "---"), "你好a---")
}

func TestStrHasSuffixInsensitive(t *testing.T) {
	a := "abceEfg"
	assert.True(t, StrHasSuffixInsensitive(a, "efg"))
	assert.True(t, StrHasSuffixInsensitive(a, "efG"))
}

func TestStrHasPrefixInsensitive(t *testing.T) {
	a := "AbceEfg"
	assert.True(t, StrHasPrefixInsensitive(a, "abc"))
	assert.True(t, StrHasPrefixInsensitive(a, "aBc"))
}

func TestStrHasContainInsensitive(t *testing.T) {
	a := "AbceEfg"
	assert.True(t, StrHasContainInsensitive(a, "Cee"))
	assert.True(t, StrHasContainInsensitive(a, "ef"))
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
	assert.Equal(t, StrSplitNth(a, ",", 3, 2), "")
	assert.Equal(t, StrSplitNth(a, ",", 4, -3), "b")
	assert.Equal(t, StrSplitNth(a, ",", 4, -12), "")
	assert.Equal(t, StrSplitNth(a, ",", 5, 2), "")
	assert.Equal(t, StrSplitNth(a, ",", 4, 12), "")
	assert.Equal(t, StrSplitNth(a, ",", 4, -4), "a")
}

func TestStrReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse odd string",
			args: args{s: "abcde"},
			want: "edcba",
		},
		{
			name: "reverse even string",
			args: args{s: "abcd"},
			want: "dcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StrReverse(tt.args.s), "StrReverse(%v)", tt.args.s)
		})
	}
}
