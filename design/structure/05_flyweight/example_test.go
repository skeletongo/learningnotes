package flyweight

import "fmt"

func ExampleShapeFactory_GetCircle() {
	f := new(ShapeFactory)
	s := f.GetCircle("red")
	s1 := f.GetCircle("red")
	s2 := f.GetCircle("yellow")
	fmt.Println(s == s1)
	fmt.Println(s == s2)
	// Output:
	// true
	// false
}
