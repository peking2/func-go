package func_go

import (
	//"fmt"
	"reflect"
)

type Map map[Any]Any

/*
Create new Map. Examples:
  1. NewMap()
  2. NewMap(map[string]string {"name": "Mike"})
*/
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

func (m *Map) ForEach(lambda func(Any, Any)) {
	for k, v := range *m {
		lambda(k, v)
	}
}
