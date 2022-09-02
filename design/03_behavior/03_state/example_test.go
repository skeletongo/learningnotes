package state

func ExampleContext_SetState() {
	c := new(Context)
	c.SetState(StateEat)
	c.Done()
	c.SetState(StateSleep)
	c.Done()
	// Output:
	// 吃饭中
	// 睡觉中
}
