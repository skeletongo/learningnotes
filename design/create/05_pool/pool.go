package pool

// 对象池模式
// 预先生成一些对象

func NewPool(n int, create func() interface{}) chan interface{} {
	ch := make(chan interface{}, n)
	defer close(ch)

	for i := 0; i < n; i++ {
		ch <- create()
	}
	return ch
}
