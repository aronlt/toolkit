package tutils

import "testing"

func TestRandStringBytesMask(t *testing.T) {
	v := RandStringBytesMask(12)
	t.Logf("%s", v)
}
