package terror

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := Wrap(errors.New("111"))
	t.Logf("%+v", err)
}

func TestErrorf(t *testing.T) {
	err := Wrapf(errors.New("1111"), "test wrapf, num:%d", 1)
	fmt.Printf("%+v", err)
}

func Foo(v int) (int, error) {
	if v < 0 {
		return -1, fmt.Errorf("invalid args")
	}
	return v * 2, nil
}

func TestAccept(t *testing.T) {
	v := Accept[int](Foo(10)).Ok(func(value int) error {
		fmt.Println("call Foo success")
		return nil
	}).Error(func(value int) error {
		fmt.Println("call Foo fail")
		return nil
	}).AcceptErr()
	assert.Nil(t, v)

	v2 := AcceptFn[int](func() (int, error) {
		v3, err := Foo(-10)
		if err != nil {
			panic(err)
		}
		return v3, err
	}).Ok(func(value int) error {
		fmt.Println("ok")
		return nil
	}).Result()
	assert.Equal(t, 0, v2.Value())
	assert.NotNil(t, v2.AcceptErr())

	v4 := Accept[int](Foo(10)).Ok(func(value int) error {
		fmt.Println("call Foo success")
		panic("panic in ok")
	}).Error(func(value int) error {
		fmt.Println("call Foo fail")
		return nil
	}).Result()
	assert.NotNil(t, v4.RunningErr())
}
