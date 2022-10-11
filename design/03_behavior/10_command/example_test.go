package command

func ExampleInvoker_Execute() {
	i := new(Invoker)
	i.Add(CreateCommand(Acommand, new(ReceiverA)))
	i.Add(CreateCommand(Bcommand, new(ReceiverB)))
	i.Execute()
	// Output:
	// 接收者A处理请求
	// 接收者B处理请求
}
