package bridge

func ExamplePhone_Run() {
	p := new(Phone)
	p.Set(new(Cpu1))
	p.Run()

	p.Set(new(Cpu2))
	p.Run()

	// Output:
	// cpu1 run...
	// cpu2 run...
}
