package prototype

import "fmt"

func ExampleClone() {
	p := &Phone{
		CPU:    "A15",
		Memory: "8GB",
		IMEI:   "0000",
	}
	fmt.Printf("%+v\n", *p)
	p1 := Clone(p, "1111")
	fmt.Printf("%+v\n", *p1)
	// Output:
	// {CPU:A15 Memory:8GB IMEI:0000}
	// {CPU:A15 Memory:8GB IMEI:1111}
}
