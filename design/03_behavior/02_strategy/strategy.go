package strategy

// 策略模式
// 和桥接模式类似，调用者在使用某个对象时不直接获取对象实例，而是通过一个中间层间接调用目标对象
// 1.抽象需要调用的方法
// 2.创建中间层对象
// 3.调用方通过中间层对象调用抽象方法

// Operate 需要的抽象方法
type Operate interface {
	Apply(int, int) int
}

// Operation 方法包装器
type Operation struct {
	operate Operate
}

func (o *Operation) Apply(a, b int) int {
	return o.operate.Apply(a, b)
}

// Sum 求和算法
type Sum struct {
}

func (s *Sum) Apply(a, b int) int {
	return a + b
}

// Mul 乘积算法
type Mul struct {
}

func (m *Mul) Apply(a, b int) int {
	return a * b
}

func NewOperation(o Operate) *Operation {
	return &Operation{operate: o}
}
