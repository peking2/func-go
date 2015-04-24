package func_go

import "testing"

func TestToInt(t *testing.T) {
	var x Any = 1
	res := ToInt(x)
	if res != 1 {
		t.Error("Expected 1, got ", res)
	}
}
