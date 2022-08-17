package abstractfactory

import "fmt"

func ExampleNewFactory() {
	fac := NewFactory(Slow)
	car := fac.Produce()
	fmt.Println(car.Speed())
	fac = NewFactory(Fast)
	car = fac.Produce()
	fmt.Println(car.Speed())
	// Output:
	// 10
	// 20
}
