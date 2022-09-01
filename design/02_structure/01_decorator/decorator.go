package decorator

import (
	"fmt"
	"time"
)

// 装饰器模式
// 使用结构体组合的方式动态添加功能
// 1.定义装饰器接口
// 2.创建装饰器
// 3.在被装饰对象中使用装饰器接口

type Printer interface {
	PrintString(string)
	PrintInt(int)
}

type PersonPrint struct {
}

func (p *PersonPrint) PrintString(name string) {
	fmt.Println(name)
}

func (p *PersonPrint) PrintInt(age int) {
	fmt.Println(age)
}

type Person struct {
	Printer
	Name string
	Age  int
}

func (p *Person) GetName() string {
	p.PrintString(p.Name)
	return p.Name
}

func (p *Person) GetAge() int {
	p.PrintInt(p.Age)
	return p.Age
}

// 方法装饰器

func Consumer(f func()) func() {
	return func() {
		t := time.Now()
		f()
		sub := time.Now().Sub(t)
		fmt.Printf("消耗时间：%d微妙\n", sub.Microseconds())
		fmt.Printf("消耗时间：%d毫秒\n", sub.Milliseconds())
		fmt.Printf("消耗时间：%f秒\n", sub.Seconds())
	}
}
