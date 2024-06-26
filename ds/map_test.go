package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapOpAppendValue(t *testing.T) {
	m := make(map[string][]int, 0)
	MapOpAppendValue(m, "a", 1)
	MapOpAppendValue(m, "a", 2)
	MapOpAppendValue(m, "b", 3)
	assert.Equal(t, m["a"], []int{1, 2})
	assert.Equal(t, m["b"], []int{3})

}

func TestMapNativeKeyCompare(t *testing.T) {
	m1 := make(map[string]string, 0)
	m2 := make(map[string]string, 0)
	m1["a"] = "a"
	m1["b"] = "b"
	m1["d"] = "d"

	m2["c"] = "a"
	m2["a"] = "a2"
	m2["b"] = "b"

	assert.Equal(t, MapCmpWithSimpleKey(m1, m2, "a"), LeftLessThanRight)
	assert.Equal(t, MapCmpWithSimpleKey(m1, m2, "b"), Equal)
	assert.Equal(t, MapCmpWithSimpleKey(m1, m2, "c"), LeftKeyMiss)
	assert.Equal(t, MapCmpWithSimpleKey(m1, m2, "d"), RightKeyMiss)
	assert.Equal(t, MapCmpWithSimpleKey(m1, m2, "e"), AllKeyMiss)
}

func TestMapNativeKeyFullCompare(t *testing.T) {
	m1 := make(map[string]string, 0)
	m2 := make(map[string]string, 0)
	m1["a"] = "a"
	m1["b"] = "b"
	m1["d"] = "d"

	m2["c"] = "a"
	m2["a"] = "a2"
	m2["b"] = "b"

	assert.Equal(t, MapCmpFullSimpleKey(m1, m2), NotEqual)

	m3 := map[string]string{"a": "a", "b": "b"}
	m4 := map[string]string{"a": "a", "b": "b"}
	assert.Equal(t, MapCmpFullSimpleKey(m3, m4), Equal)
}

func TestMapComplexKeyCompare(t *testing.T) {
	type T struct {
		A int
	}
	m1 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m2 := map[string]T{
		"a": {
			A: 3,
		},
	}
	assert.Equal(t, MapCmpWithComplexKey(m1, m2, "a"), NotEqual)

	m3 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m4 := map[string]T{
		"a": {
			A: 1,
		},
	}
	assert.Equal(t, MapCmpWithComplexKey(m3, m4, "a"), Equal)
}

func TestMapComplexKeyFullCompare(t *testing.T) {
	type T struct {
		A int
	}
	m1 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m2 := map[string]T{
		"a": {
			A: 3,
		},
	}
	assert.Equal(t, MapCmpFullComplexKey(m1, m2), NotEqual)

	m3 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m4 := map[string]T{
		"a": {
			A: 1,
		},
	}
	assert.Equal(t, MapCmpFullComplexKey(m3, m4), Equal)
}

func TestSortedMap(t *testing.T) {
	m := make(map[int]int, 0)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	tuples := BuildOrderTuples(m)
	for i := 0; i < 10; i++ {
		assert.Equal(t, tuples[i].Key, i)
		assert.Equal(t, tuples[i].Value, i)
	}
	for i := 0; i < 10; i++ {
		m[i] = i + 1
	}
	tuples = BuildOrderTuples(m)
	for i := 0; i < 10; i++ {
		assert.Equal(t, tuples[i].Key, i)
		assert.Equal(t, tuples[i].Value, i+1)
	}
}

func TestMergeMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := map[int]int{1: 2, 4: 4, 5: 5}
	m3 := MapOpMerge(m1, m2)
	assert.Equal(t, MapCmpFullComplexKey(m3, map[int]int{
		1: 2, 2: 2, 3: 3, 4: 4, 5: 5,
	}), Equal)

}

func TestMapOPMergeIfDupFn(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := map[int]int{1: 2, 4: 4, 5: 5}
	m3 := MapOPMergeIfDupFn(m1, m2, func(k int) int {
		v1, _ := m1[k]
		v2, _ := m2[k]
		return v1 + v2
	})
	assert.Equal(t, MapCmpFullComplexKey(m3, map[int]int{
		1: 3, 2: 2, 3: 3, 4: 4, 5: 5,
	}), Equal)
}

func TestMapOPMergeWithFn(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := map[int]int{1: 2, 4: 4, 5: 5}
	m3 := MapOPMergeWithFn(m1, m2, func(k int) int {
		v1, _ := m1[k]
		v2, _ := m2[k]
		return v1 + v2
	})
	assert.Equal(t, MapCmpFullComplexKey(m3, map[int]int{
		1: 3, 2: 2, 3: 3, 4: 4, 5: 5,
	}), Equal)
}

func TestMapConvert(t *testing.T) {
	v := map[int]string{1: "1", 2: "2", 3: "3"}
	s1 := MapConvertValueToSlice(v)
	assert.True(t, SliceCmpLogicEqual(s1, []string{"1", "2", "3"}))
	s2 := MapConvertKeyToSlice(v)
	assert.True(t, SliceCmpLogicEqual(s2, []int{1, 2, 3}))
	v2, err := MapConvertZipSliceToMap([]int{1, 2, 3}, []string{"1", "2", "3"})
	assert.Nil(t, err)
	assert.Equal(t, v2, v)

}

func TestMapOp(t *testing.T) {
	v3 := map[int]int{1: 1, 2: 2, 3: 3}
	v4 := MapOpRemoveValue(v3, 3)
	assert.Equal(t, MapCmpFullSimpleKey(map[int]int{1: 1, 2: 2}, v4), Equal)
}

func TestMapConvertKeyToSet(t *testing.T) {
	v := map[int]string{1: "", 2: "2", 3: "3"}
	v2 := MapConvertKeyToSet(v)
	assert.True(t, SliceCmpLogicEqual(SetToSlice(v2), []int{1, 2, 3}))
}

func TestMapConvertValueToSet(t *testing.T) {
	v := map[int]string{1: "2", 2: "2", 3: "3"}
	v2 := MapConvertValueToSet(v)
	assert.True(t, SliceCmpLogicEqual(SetToSlice(v2), []string{"2", "3"}))
}

func TestMapOpRemoveValueInSlice(t *testing.T) {
	v := map[int][]string{1: {"1", "2"}, 2: {"2"}, 3: {"3"}}
	v = MapOpRemoveValueInSlice(v, "2", true)
	assert.Equal(t, v, map[int][]string{1: {"1"}, 3: {"3"}})

	v = map[int][]string{1: {"1", "2"}, 2: {"2"}, 3: {"3"}}
	v = MapOpRemoveValueInSlice(v, "2", false)
	assert.Equal(t, v, map[int][]string{1: {"1"}, 2: {}, 3: {"3"}})
}

func TestMapOpDeepCopy(t *testing.T) {
	a := 1
	b := 2
	m := map[int]*int{1: &a, 2: &b}
	nm := MapOpDeepCopy(m)
	assert.Equal(t, nm, m)
	*nm[1] = 10
	assert.Equal(t, 1, *m[1])

	m2 := map[int]int{1: 1, 2: 2}
	m3 := MapOpCopy(m2)
	assert.Equal(t, m2, m3)
}

func TestMapOpPop(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2}
	v, ok := MapOpPop(m, "1")
	assert.Equal(t, ok, true)
	assert.Equal(t, 1, v)
	assert.Equal(t, m, map[string]int{"2": 2})

	m = map[string]int{"1": 1, "2": 2}
	v, ok = MapOpPop(m, "3")
	assert.Equal(t, ok, false)
	assert.Equal(t, 0, v)
	assert.Equal(t, m, map[string]int{"1": 1, "2": 2})
}

func TestMapGetDefault(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2}
	v1 := MapGetDefault(m, "1", 100)
	assert.Equal(t, v1, 1)
	v2 := MapGetDefault(m, "3", 100)
	assert.Equal(t, v2, 100)
}

func TestMapOpSetIfEmpty(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2}
	v1, ok := MapOpSetIfEmpty(m, "1", 100)
	assert.Equal(t, v1, 1)
	assert.False(t, ok)
	v2, ok := MapOpSetIfEmpty(m, "3", 100)
	assert.True(t, ok)
	assert.Equal(t, v2, 100)
}
