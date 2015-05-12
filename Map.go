package func_go

import (
	//"fmt"
	"reflect"
)

type Map map[Any]Any

func NewMap(args ...Any) *Map {
	if len(args) == 0 {
		return &Map{}
	}

	if len(args) > 1 {
		panic(ERROR_INVALID_PARAMETER)
	}

	m := args[0]

	if reflect.ValueOf(m).Kind() == reflect.Map {
		m := reflect.ValueOf(m)
		res := make(Map)
		keys := m.MapKeys()
		for _, k := range keys {
			res[k.Interface()] = m.MapIndex(k).Interface()
		}
		return &res
	} else {
		panic(ERROR_INVALID_PARAMETER)
	}
}
