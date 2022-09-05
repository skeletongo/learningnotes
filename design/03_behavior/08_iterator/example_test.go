package iterator

import "fmt"

type People struct {
	Name string
}

func (p *People) Visit() interface{} {
	return p.Name
}

func Example() {
	list := new(Container)

	d := new(Iterator)
	d.Container = list

	fmt.Println(d.HasNext())

	list.Add(&People{
		Name: "Tom",
	})
	list.Add(&People{
		Name: "Jack",
	})

	fmt.Println(d.HasNext())

	for d.HasNext() {
		fmt.Println(d.Next().Visit())
	}
	// Output:
	// false
	// true
	// Tom
	// Jack
}
