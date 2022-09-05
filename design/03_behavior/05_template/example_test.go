package template

func Example() {
	c := new(Cat)
	p := new(P)
	p.Concrete = c
	p.Name = "cat"
	c.P = *p

	p.Eat()
	p.Sleep()
	// Output:
	// cat eat
	// cat sleep
}
