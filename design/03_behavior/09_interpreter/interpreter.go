package interpreter

import "strings"

// 解释器模式

type Exp int

const (
	Eq Exp = iota
	Cont
)

type Expression interface {
	Interpret() bool
}

type Context struct {
	val string
}

func (c *Context) Get() string {
	return c.val
}

type Equal struct {
	left, right Context
}

func (e *Equal) Interpret() bool {
	return e.left.Get() == e.right.Get()
}

type Contain struct {
	left, right Context
}

func (b *Contain) Interpret() bool {
	return strings.Contains(b.left.Get(), b.right.Get())
}

func CreateExpression(exp Exp, left, right Context) Expression {
	switch exp {
	case Eq:
		return &Equal{
			left:  left,
			right: right,
		}

	case Cont:
		return &Contain{
			left:  left,
			right: right,
		}
	default:
		return nil
	}
}
