package enum

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestJson(t *testing.T) {
	type Value struct {
		M LeaveType `json:"m"`
	}

	value := Value{M: AnnualLeave}

	bytes, err := json.Marshal(value)
	assert.Nil(t, err)
	assert.Equal(t, string(bytes), "{\"m\":\"AnnualLeave\"}")

	v := &Value{}
	err = json.Unmarshal(bytes, v)
	assert.Nil(t, err)
	assert.Equal(t, AnnualLeave, v.M)
}

func TestGorm(t *testing.T) {
	type User struct {
		Id         int64     `gorm:"column:id"`
		Username   string    `gorm:"column:username"`
		Password   string    `gorm:"column:password"`
		Type       LeaveType `gorm:"column:type"`
		CreateTime int64     `gorm:"column:createtime"`
	}
	// user define dsn
	var dsn string
	//连接MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	//定义一个用户，并初始化数据
	u := User{
		Username:   "tizi365",
		Password:   "123456",
		Type:       AnnualLeave,
		CreateTime: time.Now().Unix(),
	}

	err = db.Create(&u).Error
	assert.Nil(t, err)

	var user User
	err = db.Where("username=?", "tizi365").Find(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Type, AnnualLeave)
}
