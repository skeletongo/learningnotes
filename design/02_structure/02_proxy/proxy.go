package proxy

import "fmt"

// 代理模式
// 通过代理访问真实对象
// 1.创建代理对象
// 2.代理对象实现真实对象的行为
// 3.代理对象访问真实对象

type Router interface {
	Handler(string)
}

type route struct {
}

func (r *route) Handler(s string) {
	fmt.Println(s)
}

type RouteProxy struct {
	*route
}

func (r *RouteProxy) Handler(s string) {
	r.route.Handler(fmt.Sprint("user_", s)) // 例如添加前缀，只是举个例子
}
