package tsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Blog struct {
	ID      int    `gorm:"column:id;primary_key"`
	Author  string `gorm:"embedded"`
	User    string `gorm:"column:user"`
	Upvotes int32  `gorm:"column:-"`
}

func TestSelectAll(t *testing.T) {
	b := Blog{}
	sql := SelectAll(b)
	expect := "`id`,`user`"
	assert.Equal(t, sql, expect)
}

func TestSelectAllWithPrefix(t *testing.T) {
	b := Blog{}
	sql := SelectAll(b, "blog")
	expect := "blog.`id`,blog.`user`"
	assert.Equal(t, sql, expect)
}
