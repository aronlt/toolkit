package terror

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
)

// Wrap 封装错误信息，获取错误的堆栈数据
func Wrap(err error, messages ...string) error {
	pc, file, lineNo, ok := runtime.Caller(1)
	if !ok {
		return errors.WithMessage(err, "call runtime Caller fail")
	}
	funcName := runtime.FuncForPC(pc).Name()
	info := fmt.Sprintf("file:%s:%d, function:%s", file, lineNo, funcName)
	if len(messages) != 0 {
		info += " ,message:" + messages[0] + "\n"
	} else {
		info += "\n"
	}
	return errors.WithMessage(err, info)
}
