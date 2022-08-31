package bridge

import "fmt"

// 桥接模式
// 在实例对象和抽象对象直接添加一层

type Runner interface {
	Run()
}

type Cpu1 struct {
}

func (c *Cpu1) Run() {
	fmt.Println("cpu1 run...")
}

type Cpu2 struct {
}

func (c *Cpu2) Run() {
	fmt.Println("cpu2 run...")
}

// CpuBridge 桥阶层对象
type CpuBridge struct {
	Cpu Runner
}

func (p *CpuBridge) Set(c Runner) {
	p.Cpu = c
}

func (p *CpuBridge) Run() {
	p.Cpu.Run()
}

// Phone 实例对象
type Phone struct {
	CpuBridge
}

func (p *Phone) Run() {
	p.CpuBridge.Run()
}
