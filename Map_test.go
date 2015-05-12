package func_go

import (
	"fmt"
	//"reflect"
	"testing"
)

func TestNewMapWithZeroParameter(t *testing.T) {
	m := *NewMap()
	m["name"] = "Gary"
	fmt.Println(m)

	if res := m["name"]; res != "Gary" {
		t.Error("Expected Gary, got ", res)
	}
}

func TestNewMapWithTwoParameters(t *testing.T) {
	defer func() {
		str := recover()
		if str != ERROR_INVALID_PARAMETER {
			t.Error("Expected ", ERROR_INVALID_PARAMETER)
		}
	}()

	NewMap(1, 2)

	t.Error("Expected panic")
}

func TestNewMapWithOneInvalidParameter(t *testing.T) {
	defer func() {
		str := recover()
		if str != ERROR_INVALID_PARAMETER {
			t.Error("Expected ", ERROR_INVALID_PARAMETER)
		}
	}()

	NewMap(1)

	t.Error("Expected panic")
}

func TestNewMapWithOneValidParameter(t *testing.T) {
	m := *NewMap(map[string]string{
		"name": "Gary",
	})

	if res := m["name"]; res != "Gary" {
		t.Error("Expected Gary, got ", res)
	}
}
