package adapter

func ExampleList() {
	l := new(Linux)
	w := new(Windows)
	wa := &WindowsAdapter{
		w,
	}

	List(l)
	List(wa)
	// Output:
	// linux ls
	// windows dir
}
