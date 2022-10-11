package interpreter

import "fmt"

func ExampleCreateExpression() {
	e := CreateExpression(Eq, Context{val: "a"}, Context{val: "a"})
	fmt.Println(e.Interpret())
	e = CreateExpression(Eq, Context{val: "a"}, Context{val: "b"})
	fmt.Println(e.Interpret())
	e = CreateExpression(Cont, Context{val: "abc"}, Context{val: "b"})
	fmt.Println(e.Interpret())
	e = CreateExpression(Cont, Context{val: "abc"}, Context{val: "d"})
	fmt.Println(e.Interpret())
	// Output:
	// true
	// false
	// true
	// false
}
