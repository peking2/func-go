package func_go

import "testing"

func TestMap(t *testing.T) {
	res := ArrayOps{1, 2, 3, 4, 5}.
		Map(func(i Any) Any { return i.(int) * 2 }).
		Reduce(func(a Any, b Any) Any { return ToInt(a) + ToInt(b) })

	if res != 30 {
		t.Error("Expected ArrayOps{2, 4, 6, 8, 10}, got ", res)
	}
}
