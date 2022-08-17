package pool

import "fmt"

func ExampleNewPool() {
	n := 0
	numbers := NewPool(5, func() interface{} {
		n++
		return n
	})

	for v := range numbers {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
