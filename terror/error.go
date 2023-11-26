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

// Wrapf 封装错误信息，获取错误的堆栈数据
func Wrapf(err error, format string, a ...any) error {
	message := fmt.Sprintf(format, a...)
	return Wrap(err, message)
}

// ProcessChain 链式处理器,简化对错误处理，增加panic的recover机制
type ProcessChain[T any] struct {
	value       T
	acceptedErr error
	runningErr  error
}

// Ok 如果接收到成功数据，则执行改逻辑
func (p *ProcessChain[T]) Ok(handler func(value T) error) (processChain *ProcessChain[T]) {
	defer func() {
		processChain = p
	}()
	defer func() {
		if r := recover(); r != nil {
			p.runningErr = errors.Errorf("call Ok panic, error:%+v", r)
		}
	}()
	if p.acceptedErr == nil {
		p.runningErr = handler(p.value)
	}
	return
}

// Error 如果接收到失败数据，则执行改逻辑
func (p *ProcessChain[T]) Error(handler func(value T) error) (processChain *ProcessChain[T]) {
	defer func() {
		processChain = p
	}()
	defer func() {
		if r := recover(); r != nil {
			p.runningErr = errors.Errorf("call Error panic, error:%+v", r)
		}
	}()
	if p.acceptedErr != nil {
		p.runningErr = handler(p.value)
	}
	return
}

// AllErr 取出所有的错误，包括接收的错误和运行时发生的错误
func (p *ProcessChain[T]) AllErr() (error, error) {
	return p.acceptedErr, p.runningErr
}

// AcceptErr 取出接收的错误
func (p *ProcessChain[T]) AcceptErr() error {
	return p.acceptedErr
}

// RunningErr 取出运行时的错误
func (p *ProcessChain[T]) RunningErr() error {
	return p.runningErr
}

// Value 取出运行时值
func (p *ProcessChain[T]) Value() T {
	return p.value
}
func (p *ProcessChain[T]) Result() *ProcessChain[T] {
	return p
}

func Accept[T any](value T, err error) *ProcessChain[T] {
	return &ProcessChain[T]{
		value:       value,
		acceptedErr: err,
	}
}

// AcceptFn 接收一个闭包函数，返回对应的链路
func AcceptFn[T any](fn func() (T, error)) (processChain *ProcessChain[T]) {
	processChain = &ProcessChain[T]{}
	defer func() {
		if r := recover(); r != nil {
			processChain.acceptedErr = errors.Errorf("call fn panic, error:%+v", r)
		}
	}()
	value, err := fn()
	processChain.value = value
	processChain.acceptedErr = err
	return
}
