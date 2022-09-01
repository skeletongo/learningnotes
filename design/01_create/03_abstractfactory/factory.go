package abstractfactory

// 抽象工厂模式
// 在工厂模式基础上，将工厂对象再抽象
// 我的理解是以工厂为单位将对象再分类

// 1.定义产品接口
// 2.定义工厂接口
// 3.创建实例工厂
// 4.创建实例产品

type Car interface {
	Speed() int
}

type CarFactory interface {
	Produce() Car
}

type CarA struct {
}

func (c *CarA) Speed() int {
	return 10
}

type CarB struct {
}

func (c *CarB) Speed() int {
	return 20
}

type CarAFactory struct {
}

func (c *CarAFactory) Produce() Car {
	return new(CarA)
}

type CarBFactory struct {
}

func (c *CarBFactory) Produce() Car {
	return new(CarB)
}

// 使用

const (
	Fast = 0
	Slow = 1
)

func NewFactory(tp int) CarFactory {
	switch tp {
	case Slow:
		return new(CarAFactory)
	case Fast:
		return new(CarBFactory)
	default:
		return nil
	}
}
