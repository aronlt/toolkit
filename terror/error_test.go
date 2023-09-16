package terror

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func test(err error) error {
	return Wrap(err)
}
func TestError(t *testing.T) {
	err := Wrap(test(errors.New("111")))
	fmt.Printf("%+v", err)
}
