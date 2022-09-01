package builder

import (
	"fmt"
)

// 建造者模式
// 把相似对象相同地构建过程单独提取出来

// 1.确认对象属性，struct
// 2.定义创建部分属性值的接口，interface
// 3.创建建造者，控制构建过程，不同对象具有的相同地构建过程，struct
// 4.实现对象的创建方法，struct

// 1.确认对象属性和方法

// Animal 动物对象
type Animal struct {
	Speak func()
	Color string
}

// 2.定义创建部分属性值的接口

// Builder 动物部分属性接口
// SetSpeak, SetColor 提取出的不同创建方法
// GetAnimal 使建造者可以访问具体对象
type Builder interface {
	SetSpeak() Builder
	SetColor() Builder
	GetAnimal() *Animal
}

// 3.创建建造者，控制构建过程，不同对象具有的相同地构建过程

// Director 建造者
type Director struct {
	builder Builder
}

// Construct 提取出的相同的构建过程
func (d *Director) Construct() {
	d.builder.SetColor().SetSpeak()
}

// SetAnimal 设置需要构建的对象
func (d *Director) SetAnimal(b Builder) {
	d.builder = b
}

// 4.实现对象的创建方法

// Dog 狗
// 将创建方法和对象方法分离
// Animal是实例对象
// Dog包含各部分创建方法
type Dog struct {
	Animal
}

func (d *Dog) SetSpeak() Builder {
	d.Animal.Speak = func() {
		fmt.Println("wang wang")
	}
	return d
}

func (d *Dog) SetColor() Builder {
	d.Animal.Color = "white"
	return d
}

func (d *Dog) GetAnimal() *Animal {
	return &d.Animal
}
