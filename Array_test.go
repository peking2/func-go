package func_go

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	res := NewArray(1).
		Map(func(i Any) Any { return i.(int) * 2 }).
		Reduce(func(a Any, b Any) Any { return a.(int) + b.(int) })

	if res != 2 {
		t.Error("Expected 2, got ", res)
	}
}

func TestArray(t *testing.T) {
	res := NewArray([]int{1, 2, 3, 4, 5}).
		Map(func(i Any) Any { return i.(int) * 2 }).
		Reduce(func(a Any, b Any) Any { return a.(int) + b.(int) })

	if res != 30 {
		t.Error("Expected ArrayOps{2, 4, 6, 8, 10}, got ", res)
	}
}

func TestFilter(t *testing.T) {
	res := NewArray([]int{1, 2, 3, 4, 5}).
		Filter(func(i Any) bool { return i.(int)%2 == 0 })

	if len(*res) != 2 {
		t.Error("Expected ArrayOps{2, 4, 6, 8, 10}, got ", res)
	}
}

func TestForEach(t *testing.T) {
	NewArray([]int{1, 2, 3, 4, 5}).
		ForEach(func(i Any) { fmt.Println(i) })
}
