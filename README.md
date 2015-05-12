# func-go
functional programming in Go

## Usage
```go
res := New([]int{1, 2, 3, 4, 5}).
	Map(func(i Any) Any { return i.(int) * 2 }).
	Reduce(func(a Any, b Any) Any { return a.(int) + b.(int) })
```