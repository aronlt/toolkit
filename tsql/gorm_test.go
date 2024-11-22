//go:build ignore
// +build ignore

package tsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Blog struct {
	ID      int    `gorm:"column:id;primary_key"`
	Author  string `gorm:"embedded"`
	User    string `gorm:"column:user"`
	Upvotes int32  `gorm:"column:-"`
}

type Student struct {
	ID      int     `gorm:"column:id;primary_key"`
	Name    string  `gorm:"column:name"`
	Age     int     `gorm:"column:age"`
	Score   int64   `gorm:"column:score"`
	Friends Strings `gorm:"column:friends"`
}

func TestSelectAll(t *testing.T) {
	b := Blog{}
	sql := SelectAll(b)
	expect := "`id` ,`user` "
	assert.Equal(t, sql, expect)
}

func TestSelectAllWithPrefix(t *testing.T) {
	b := Blog{}
	sql := SelectAll(b, "blog")
	expect := "blog.`id` ,blog.`user` "
	assert.Equal(t, sql, expect)
}

func newMockDB() (*gorm.DB, error) {
	dsn := ""
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func TestBatchInsert(t *testing.T) {
	type args[T any] struct {
		db     *gorm.DB
		table  string
		values []*T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    int64
		wantErr assert.ErrorAssertionFunc
	}
	db, err := newMockDB()
	assert.Nil(t, err, "call newMockDB fail")
	tests := []testCase[Student]{
		{
			name: "batch insert",
			args: args[Student]{
				db:    db,
				table: "student",
				values: []*Student{
					{
						ID:    1,
						Name:  "name",
						Age:   10,
						Score: 100,
					},
				},
			},
			want:    1,
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := int64(1)
			got, err := BatchInsert(tt.args.db, tt.args.table, tt.args.values)
			assert.Nil(t, err)
			assert.Equal(t, got, count)
			value, err := SelectById[Student](tt.args.db, tt.args.table, 1)
			assert.Nil(t, err)
			assert.Equal(t, value.Name, tt.args.values[0].Name)
			got, err = Delete[Student](tt.args.db, tt.args.table, map[string]interface{}{"name": "name"})
			assert.Nil(t, err)
			assert.Equal(t, got, count)

			got, err = Insert[Student](tt.args.db, tt.args.table, tt.args.values[0])
			assert.Nil(t, err)
			assert.Equal(t, got, count)

			names, err := SelectPluck[Student, string](tt.args.db, tt.args.table, map[string]interface{}{
				"age": 10,
			}, "name")
			assert.Nil(t, err)
			assert.Equal(t, names, []string{"name"})

			got, err = Update[Student](tt.args.db, tt.args.table, map[string]interface{}{
				"name": "name",
			}, map[string]interface{}{
				"age": 11,
			})
			assert.Nil(t, err)
			assert.Equal(t, got, count)
			student, err := SelectById[Student](tt.args.db, tt.args.table, 1)
			assert.Nil(t, err)
			assert.Equal(t, student.Age, 11)

			students, err := Select[Student](tt.args.db, tt.args.table, map[string]interface{}{
				"age": 11,
			})
			assert.Nil(t, err)
			assert.Equal(t, students[0].Age, 11)

			students, err = SelectRaw[Student](tt.args.db, "select * from student where id = ?", 1)
			assert.Nil(t, err)
			assert.Equal(t, students[0].Age, 11)

			err = ExecuteRaw(tt.args.db, "delete from student where id > ?", 0)
			assert.Nil(t, err)
		})
	}
}
