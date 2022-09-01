package factory

import (
	"fmt"
)

func ExampleGenerate() {
	r, err := Generate(Television)
	if err != nil {
		fmt.Println(err)
	}
	if r != nil {
		fmt.Println(r.GetName())
	}

	r, err = Generate(Refrigerator)
	if err != nil {
		fmt.Println(err)
	}
	if r != nil {
		fmt.Println(r.GetName())
	}

	r, err = Generate(NilKind)
	if err != nil {
		fmt.Println(err)
	}
	if r != nil {
		fmt.Println(r.GetName())
	}
	// Output:
	// 电视
	// 冰箱
	// not found, kind is 0
}
