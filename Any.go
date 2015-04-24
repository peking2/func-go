package func_go

type Any interface{}

func ToInt(x Any) int {
	return x.(int)
}
