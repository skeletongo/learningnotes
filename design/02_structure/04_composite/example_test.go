package composite

import "fmt"

func ExampleNewItemGroup() {
	a1 := NewItem("苹果", "2个苹果", 2.5)
	a2 := NewItem("草莓", "一斤草莓", 3.5)

	a1.Print()
	a2.Print()
	fmt.Println()

	g := NewItemGroup("水果套装", "苹果和草莓", a1, a2)
	g.Print()
	fmt.Println()

	a3 := NewItem("草莓", "一斤草莓", 3.5)
	g.Add(a3)
	g.Print()
	// Output:
	//苹果, ¥2.50, 2个苹果
	//草莓, ¥3.50, 一斤草莓
	//
	//水果套装, 苹果和草莓, ¥6.00
	//------------------------
	//苹果, ¥2.50, 2个苹果
	//草莓, ¥3.50, 一斤草莓
	//
	//水果套装, 苹果和草莓, ¥9.50
	//------------------------
	//苹果, ¥2.50, 2个苹果
	//草莓, ¥3.50, 一斤草莓
	//草莓, ¥3.50, 一斤草莓
}
