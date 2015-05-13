package func_go

import (
	//"fmt"
	"reflect"
)

type Try struct {
	value Any
}

/*
Wrap two kinds of functions:
  1. func(...Any) (error)     -> func (...Any) Try
  2. func(...Any) (Any error) -> func (...Any) Try
*/
func WrapFunc(f Any) func(...Any) Try {
	v := reflect.ValueOf(f)

	if v.Kind() != reflect.Func {
		panic(ERROR_INVALID_PARAMETER)
	}

	return func(args ...Any) Try {
		in := make([]reflect.Value, len(args))
		for i, v := range args {
			in[i] = reflect.ValueOf(v)
		}

		res := v.Call(in)
		switch len(res) {
		case 1:
			if res[0].IsNil() {
				return Try{nil}
			} else {
				return Try{res[0].Interface().(error)}
			}
		case 2:
			if res[1].IsNil() {
				return Try{res[0].Interface()}
			} else {
				return Try{res[1].Interface().(error)}
			}
		default:
			panic(ERROR_INVALID_PARAMETER)
		}
	}
}

func (t Try) IsFailure() bool {
	switch t.value.(type) {
	case error:
		return true
	default:
		return false
	}
}

func (t Try) IsSuccess() bool {
	return !t.IsFailure()
}

func (t Try) Get() Any {
	return t.value
}

func (t Try) FlatMap(lambda func(Any) Try) Try {
	if t.IsFailure() {
		return t
	}

	return lambda(t.value)
}
