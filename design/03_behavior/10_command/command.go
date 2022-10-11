package command

import "fmt"

// 命令模式

// Receiver 接受者执行的具体请求
type Receiver interface {
	Execute()
}

// 创建具体的接收者，实现接口方法
type ReceiverA struct {
}

func (a *ReceiverA) Execute() {
	fmt.Println("接收者A处理请求")
}

type ReceiverB struct {
}

func (b *ReceiverB) Execute() {
	fmt.Println("接收者B处理请求")
}

// Command 调用者使用的接口
type Command interface {
	Call()
}

// 创建具体command, 指定接收者
type ConcreteCommandA struct {
	Receiver
}

func (ca *ConcreteCommandA) Call() {
	ca.Receiver.Execute()
}

type ConcreteCommandB struct {
	Receiver
}

func (cb *ConcreteCommandB) Call() {
	cb.Receiver.Execute()
}

// Invoker 调用者
type Invoker struct {
	cmds []Command
}

func (i *Invoker) Add(c Command) {
	i.cmds = append(i.cmds, c)
}

func (i *Invoker) Execute() {
	for _, v := range i.cmds {
		v.Call()
	}
}

//使用工厂方法模式来创建ConcreteCommand
type TYPE string

const (
	Acommand TYPE = "a"
	Bcommand TYPE = "b"
)

func CreateCommand(kind TYPE, receiver Receiver) Command {
	switch kind {
	case Acommand:
		return &ConcreteCommandA{receiver}
	case Bcommand:
		return &ConcreteCommandB{receiver}
	default:
		return nil
	}
}
