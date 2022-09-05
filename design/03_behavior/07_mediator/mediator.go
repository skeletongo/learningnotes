package mediator

import "fmt"

// 中介者模式
// 通过中介者对象调用其它对象的方法

// IDepartment 被调用者接口
type IDepartment interface {
	Send(message string) // 发送给中介者消息的方法
	Get(message string)  // 实际调用的方法
}

type Technical struct {
	mediator *Mediator
}

func (t *Technical) Send(message string) {
	t.mediator.Forward(t, message)
}

func (t *Technical) Get(message string) {
	fmt.Printf("技术部收到消息: %s\n", message)
}

type Market struct {
	mediator *Mediator
}

func (m *Market) Send(message string) {
	m.mediator.Forward(m, message)
}

func (m *Market) Get(message string) {
	fmt.Printf("市场部部收到消息: %s\n", message)
}

// Mediator 中介者
type Mediator struct {
	*Technical
	*Market
}

func (m *Mediator) Forward(d IDepartment, message string) {
	// 注意这里调用的是中介中包含的具体实例，参数d只是用来判断调用那个实例
	switch d.(type) {
	case *Technical:
		m.Technical.Get(message)
	case *Market:
		m.Market.Get(message)
	}
}
