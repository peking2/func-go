package func_go

type ArrayOps []Any

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
