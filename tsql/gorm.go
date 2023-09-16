package tsql

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

func setDB[T any](db *gorm.DB, table string, condition map[string]interface{}) *gorm.DB {
	var empty T
	db = db.Table(table).Model(&empty)
	for k, v := range condition {
		db = db.Where(fmt.Sprintf("%s = ?", k), v)
	}
	return db
}

func Delete[T any](db *gorm.DB, table string, condition map[string]interface{}) (int64, error) {
	db = setDB[T](db, table, condition)
	var empty T
	result := db.Delete(&empty)
	return result.RowsAffected, result.Error
}

func ExecuteRaw(db *gorm.DB, raw string, values ...interface{}) error {
	err := db.Exec(raw, values).Error
	return err
}

func Update[T any](db *gorm.DB, table string, condition map[string]interface{}, values map[string]interface{}) (int64, error) {
	db = setDB[T](db, table, condition)
	result := db.Updates(values)
	return result.RowsAffected, result.Error
}

func SelectRaw[T any](db *gorm.DB, raw string, values ...interface{}) ([]T, error) {
	var result []T
	err := db.Raw(raw, values).Scan(&result).Error
	return result, err
}

func SelectPluck[T any, K any](db *gorm.DB, table string, condition map[string]interface{}, column string) ([]K, error) {
	db = setDB[T](db, table, condition)
	var result []K
	err := db.Pluck(column, &result).Error
	return result, err
}

func SelectById[T any](db *gorm.DB, table string, id int64, columns ...string) (T, error) {
	var empty T
	var result T
	column := "id"
	if len(columns) != 0 {
		column = columns[0]
	}
	err := db.Table(table).Model(&empty).Where(fmt.Sprintf("%s = ?", column), id).First(&result).Error
	return result, err
}

func Select[T any](db *gorm.DB, table string, condition map[string]interface{}) ([]T, error) {
	var result []T
	db = setDB[T](db, table, condition)
	err := db.Find(&result).Error
	return result, err
}

func Insert[T any](db *gorm.DB, table string, value *T) (int64, error) {
	var empty T
	result := db.Table(table).Model(&empty).Create(value)
	return result.RowsAffected, result.Error
}

func BatchInsert[T any](db *gorm.DB, table string, values []*T) (int64, error) {
	var empty T
	result := db.Table(table).Model(&empty).Create(values)
	return result.RowsAffected, result.Error
}

// SelectAll 生成select a, b, c, d 前缀
func SelectAll[T any](data T, prefix ...string) string {
	parts := make([]string, 0)
	t := reflect.TypeOf(data)
	count := t.NumField()
	for i := 0; i < count; i++ {
		tag := t.Field(i).Tag
		if v, ok := tag.Lookup("gorm"); ok {
			startIndex := strings.Index(v, "column:")
			if startIndex != -1 {
				v = v[startIndex+len("column:"):]
				column := v
				endIndex := strings.Index(v, ";")
				if endIndex != -1 {
					column = v[:endIndex]
				}
				if column == "-" {
					continue
				}
				if len(prefix) > 0 {
					parts = append(parts, prefix[0]+"."+"`"+column+"` ")
				} else {
					parts = append(parts, "`"+column+"` ")
				}
			}
		}
	}
	return strings.Join(parts, ",")
}

// SecTimestamp 用于mysql的秒时间戳类型
type SecTimestamp int64

func NewSecTimestamp(t time.Time) SecTimestamp {
	return SecTimestamp(t.Unix())
}

func (s *SecTimestamp) Time() time.Time {
	return time.Unix(int64(*s), 0)
}

func (s *SecTimestamp) Scan(src interface{}) error {
	if s == nil {
		return nil
	}
	switch t := src.(type) {
	case time.Time:
		if t.IsZero() {
			*s = 0
		} else {
			*s = NewSecTimestamp(t)
		}
	case int64:
		*s = SecTimestamp(t)
	default:
		return fmt.Errorf("converting driver.Value type %T (%q) to a %T: invalid syntax", t, t, *s)
	}
	return nil
}

func (s *SecTimestamp) Value() (driver.Value, error) {
	return s.Time(), nil
}

func IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
