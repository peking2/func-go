package arrayops

import (
	//"fmt"
	"reflect"

	. "../any"
)

type ArrayOps []Any

func New(arr Any) ArrayOps {
	kind := reflect.ValueOf(arr).Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		arr := reflect.ValueOf(arr)
		res := make(ArrayOps, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			res[i] = arr.Index(i).Interface()
		}
		return res
	} else {
		return ArrayOps{arr}
	}
}

func (arr ArrayOps) Map(lambda func(Any) Any) ArrayOps {
	res := make(ArrayOps, len(arr))
	for i, v := range arr {
		res[i] = lambda(v)
	}
	return res
}

func (arr ArrayOps) Reduce(lambda func(Any, Any) Any) Any {
	res := arr[0]

	for i := 1; i < len(arr); i++ {
		res = lambda(res, arr[i])
	}

	return res
}

func (arr ArrayOps) Filter(lambda func(Any) bool) ArrayOps {
	res := make(ArrayOps, 0)
	for _, v := range arr {
		if lambda(v) {
			res = append(res, v)
		}
	}
	return res
}
