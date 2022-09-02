package state

import "fmt"

// 状态模式
// 根据不同状态对行为进行封装
// 1.行为抽象
// 2.创建行为对象实例
// 3.创建状态控制对象

type Action interface {
	Done()
}

type Eat struct {
}

func (e *Eat) Done() {
	fmt.Println("吃饭中")
}

type Sleep struct {
}

func (s *Sleep) Done() {
	fmt.Println("睡觉中")
}

const (
	StateEat   = 1
	StateSleep = 2
)

type Context struct {
	State Action
	id    int
}

func (c *Context) SetState(id int) {
	// 可以用享元模式优化一下
	switch id {
	case StateEat:
		c.State = new(Eat)
	case StateSleep:
		c.State = new(Sleep)
	}
}

func (c *Context) Done() {
	c.State.Done()
}
