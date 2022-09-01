package observer

func ExamplePerson_Receiver() {
	a := new(Person)
	a.Name = "张三"

	b := new(Person)
	b.Name = "李四"

	s := new(Stock)
	s.Add(a)
	s.Add(b)

	s.Price = 100
	s.Notify()
	// Output:
	// 张三 收到价格变更： 100
	// 李四 收到价格变更： 100
}
