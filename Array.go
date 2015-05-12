package func_go

import (
	//"fmt"
	"reflect"
)

type Array []Any

func NewArray(arr Any) *Array {
	kind := reflect.ValueOf(arr).Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		arr := reflect.ValueOf(arr)
		res := make(Array, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			res[i] = arr.Index(i).Interface()
		}
		return &res
	} else {
		return &Array{arr}
	}
}

func (arr *Array) Map(lambda func(Any) Any) *Array {
	res := make(Array, len(*arr))
	for i, v := range *arr {
		res[i] = lambda(v)
	}
	return &res
}

func (arr *Array) Reduce(lambda func(Any, Any) Any) Any {
	res := (*arr)[0]

	for i := 1; i < len(*arr); i++ {
		res = lambda(res, (*arr)[i])
	}

	return res
}

func (arr *Array) Filter(lambda func(Any) bool) *Array {
	res := make(Array, 0)
	for _, v := range *arr {
		if lambda(v) {
			res = append(res, v)
		}
	}
	return &res
}

func (arr *Array) ForEach(lambda func(Any)) {
	for _, v := range *arr {
		lambda(v)
	}
}
