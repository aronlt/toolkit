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
	sortedMap := MapNewSortedMap(m)
	for i := 0; i < 10; i++ {
		assert.Equal(t, sortedMap.Tuples[i].Key, i)
		assert.Equal(t, sortedMap.Tuples[i].Value, i)
	}
	for i := 0; i < 10; i++ {
		sortedMap.RawMap[i] = i + 1
	}
	sortedMap.Rebuild()
	for i := 0; i < 10; i++ {
		assert.Equal(t, sortedMap.Tuples[i].Key, i)
		assert.Equal(t, sortedMap.Tuples[i].Value, i+1)
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

func TestMapOpRemoveValueInSlice(t *testing.T) {
	v := map[int][]string{1: {"1", "2"}, 2: {"2"}, 3: {"3"}}
	v = MapOpRemoveValueInSlice(v, "2", true)
	assert.Equal(t, v, map[int][]string{1: {"1"}, 3: {"3"}})

	v = map[int][]string{1: {"1", "2"}, 2: {"2"}, 3: {"3"}}
	v = MapOpRemoveValueInSlice(v, "2", false)
	assert.Equal(t, v, map[int][]string{1: {"1"}, 2: {}, 3: {"3"}})
}
