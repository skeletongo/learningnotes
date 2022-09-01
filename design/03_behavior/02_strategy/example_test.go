package strategy

import "fmt"

func ExampleOperation_Apply() {
	sum := new(Sum)
	mul := new(Mul)

	o := NewOperation(sum)
	fmt.Println(o.Apply(1, 2))

	o = NewOperation(mul)
	fmt.Println(o.Apply(1, 2))
	// Output:
	// 3
	// 2
}
