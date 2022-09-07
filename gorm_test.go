package toolkit

import "testing"

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
	if sql != expect {
		t.Errorf("SelectAll expected %s actual %s", expect, sql)
	}
}

func TestSelectAllWithPrefix(t *testing.T) {
	b := Blog{}
	sql := SelectAll(b, "blog")
	expect := "blog.`id`,blog.`user`"
	if sql != expect {
		t.Errorf("SelectAll expected %s actual %s", expect, sql)
	}
}
