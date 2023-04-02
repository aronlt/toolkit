package structure

import "fmt"

/*
API 为 facade 模块的外观接口，大部分代码使用此接口简化对 facade 类的访问。
facade 模块同时暴露了 a 和 b 两个 Module 的 NewXXX 和 interface，其它代码如果需要使用细节功能时可以直接调用。
*/

// API is facade interface of facade package
type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// apiImpl facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

// 具体实现
func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
