package iterator

// 迭代器模式

type Elementor interface {
	Visit() interface{}
}

type Container struct {
	list []Elementor
}

func (c *Container) Add(e Elementor) {
	c.list = append(c.list, e)
}

func (c *Container) Remove(e Elementor) {
	for k, v := range c.list {
		if v == e {
			c.list = append(c.list[:k], c.list[k+1:]...)
			return
		}
	}
}

type Iterator struct {
	index int
	*Container
}

func (i *Iterator) Next() Elementor {
	e := i.list[i.index]
	i.index++
	return e
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.list)
}
