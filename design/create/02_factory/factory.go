package factory

import (
	"errors"
	"fmt"
)

// 工厂模式

// 1.定义类型枚举
// 2.提取对象接口
// 3.创建各种对象
// 4.创建工厂方法

// Kind 电器类型
type Kind int

const (
	NilKind = iota
	Television
	Refrigerator
)

// Electrical 电器接口
type Electrical interface {
	GetName() string
}

// Generate 工厂方法
func Generate(k Kind) (Electrical, error) {
	switch k {
	case Television:
		return &MyTelevision{
			Name: "电视",
		}, nil
	case Refrigerator:
		return &MyRefrigerator{
			Name: "冰箱",
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("not found, kind is %v", k))
	}
}

// 各种电器

type MyTelevision struct {
	Name string
}

func (m *MyTelevision) GetName() string {
	return m.Name
}

type MyRefrigerator struct {
	Name string
}

func (m *MyRefrigerator) GetName() string {
	return m.Name
}
