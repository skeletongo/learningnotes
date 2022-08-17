package builder

import "fmt"

func ExampleDirector() {
	director := new(Director)

	dog := new(Dog)
	director.SetAnimal(dog)
	director.Construct()

	fmt.Println(dog.Animal.Color)
	dog.Animal.Speak()

	// Output:
	// white
	// wang wang
}

func ExampleOrderDirector_Construct() {
	director := new(OrderDirector)

	order1 := new(Order1)
	director.SetOrderBuilder(order1)
	director.Construct()
	fmt.Println(order1.Order)

	order2 := new(Order2)
	director.SetOrderBuilder(order2)
	director.Construct()
	fmt.Println(order2.Order)

	order3 := new(Order3)
	director.SetOrderBuilder(order3)
	director.Construct()
	fmt.Println(order3.Order)

	// Output:
	// 橙汁,2.00
	// 橙汁,牛奶,6.00
	// 橙汁,牛奶,牛奶,9.00
}
