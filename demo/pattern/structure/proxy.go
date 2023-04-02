package structure

// 让别人帮助你做你并不关心的事情，叫代理模式
// 为让自己的能力增强，使得增强后的自己能够使用更多的方法，拓展在自己基础之上的功能的，叫装饰器模式
// 代理模式的常见用法有
// Cache 代理
// 防火墙代理

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}
