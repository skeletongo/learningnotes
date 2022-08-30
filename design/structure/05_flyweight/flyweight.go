package flyweight

// 享元模式
// 重复利用已有对象；多个对象共同使用一个对象；
// 类似于工厂模式，区别是工厂模式是创建新对象，享元模式是返回已有对象

type Shape interface {
	SetRadius(radius int)
	SetColor(color string)
}

type Circle struct {
	color  string
	radius int
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

type ShapeFactory struct {
	circleMap map[string]Shape
}

//GetCircle 对象不存在则创建
func (sh *ShapeFactory) GetCircle(color string) Shape {
	if sh.circleMap == nil {
		sh.circleMap = make(map[string]Shape)
	}
	if shape, ok := sh.circleMap[color]; ok {
		return shape
	}
	circle := new(Circle)
	circle.SetColor(color)
	sh.circleMap[color] = circle
	return circle
}
