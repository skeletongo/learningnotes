package decorator

func ExamplePerson_GetName() {
	pt := &PersonPrint{}

	p := Person{
		Printer: pt,
		Name:    "Tom",
		Age:     18,
	}

	p.GetName()
	p.GetAge()
	// Output:
	// Tom
	// 18
}
