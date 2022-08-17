package generator

import "testing"

func TestCount(t *testing.T) {
	var ls []int
	for i := 1; i < 10; i++ {
		ls = append(ls, i)
	}
	ch := Count(1, 9)
	for _, v := range ls {
		if <-ch != v {
			t.Error("not same")
		}
	}
}
