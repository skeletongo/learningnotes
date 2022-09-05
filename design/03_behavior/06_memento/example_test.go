package memento

import "fmt"

func Example() {
	m := new(MemoManager)
	e := new(State)
	e.SetState("hello")

	memo := m.Create(e)
	s := m.GetStater(memo)

	fmt.Println(e == s)
	fmt.Println(s.GetState())
	// Output:
	// true
	// hello
}
