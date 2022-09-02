package vistor

import "fmt"

// 访问者模式
// 数据和方法解耦，对象结构数据很少改变，但经常需要在此对象结构上定义新的操作

// Visitor 访问者接口
type Visitor interface {
	Visit(*Info)
}

// PrintVisit 访问者实例
type PrintVisit struct {
}

func (p *PrintVisit) Visit(info *Info) {
	fmt.Println(info.Text)
}

// BeVisitor 被访问者接口
type BeVisitor interface {
	Accept(Visitor)
}

// Info 被访问者实例
type Info struct {
	Text string
}

func (i *Info) Accept(v Visitor) {
	v.Visit(i)
}
