package vistor

func ExampleInfo_Accept() {
	info := &Info{
		Text: "访问者模式",
	}

	v := new(PrintVisit)
	info.Accept(v)
	// Output:
	// 访问者模式
}
