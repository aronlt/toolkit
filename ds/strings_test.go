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
