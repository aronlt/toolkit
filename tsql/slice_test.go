//go:build ignore
// +build ignore

package tsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	db, err := newMockDB()
	table := "student"
	assert.Nil(t, err)
	student := &Student{
		ID:      2,
		Name:    "name2",
		Age:     12,
		Score:   10,
		Friends: Strings{"aa", "bb", "cc"},
	}
	_, _ = Delete[Student](db, table, map[string]interface{}{"id": 2})
	got, err := Insert(db, table, student)
	assert.Nil(t, err)
	assert.Equal(t, got, int64(1))
	s, err := SelectById[Student](db, table, 2)
	assert.Nil(t, err)
	assert.Equal(t, &s, student)
	_, _ = Delete[Student](db, table, map[string]interface{}{"id": 2})
}
