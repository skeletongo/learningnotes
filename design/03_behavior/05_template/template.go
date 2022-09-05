package template

import "fmt"

// 模板方法模式
// 父类(模板)包含一个对象，这个对象是抽象接口，父类调用这个对象也实现这些抽象接口，也就是将父类方法实现延迟到子类中实现；
// 子类通过组合方式继承父类，然后覆盖父类方法，完成抽象方法的具体实现

type Animal interface {
	Eat()
	Sleep()
}

type P struct {
	Name     string
	Concrete Animal
}

func (p *P) Eat() {
	p.Concrete.Eat()
}

func (p *P) Sleep() {
	p.Concrete.Sleep()
}

type Cat struct {
	P
}

func (c *Cat) Eat() {
	fmt.Println(c.Name, "eat")
}

func (c *Cat) Sleep() {
	fmt.Println(c.Name, "sleep")
}
