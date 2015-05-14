package func_go

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestWrapFunc(t *testing.T) {
	f := func(a, b int) (int, error) {
		if a+b >= 0 {
			return a + b, nil
		} else {
			return 0, errors.New(ERROR_INVALID_PARAMETER)
		}
	}

	wf := WrapFunc(f)

	fmt.Println(reflect.TypeOf(wf))

	res := wf(1, 2)

	if res.IsFailure() {
		t.Error("Expected Success, got ", res)
	}

	if res.Get() != 3 {
		t.Error("Expected 3, got ", res.Get())
	}

	res = wf(1, -2)

	if res.IsSuccess() {
		t.Error("Expected Failure, got ", res)
	}

	if res.value.(error).Error() != ERROR_INVALID_PARAMETER {
		t.Errorf("Expected %s, got %s", ERROR_INVALID_PARAMETER, res.value)
	}
}

func TestNewTry(t *testing.T) {
	tr := NewTry(1)
	if tr.IsFailure() {
		t.Error("Expected Falure, got ", tr)
	}

	tr = NewTry(errors.New("error"))
	if tr.IsSuccess() {
		t.Error("Expected Falure, got ", tr)
	}
}
