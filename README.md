# func-go
functional programming in Go

## Installation
```bash
go get github.com/peking2/func-go
```
## Usage
```go
res := NewArrayOps([]int{1, 2, 3, 4, 5}).
	Map(func(i Any) Any { return i.(int) * 2 }).
	Reduce(func(a Any, b Any) Any { return a.(int) + b.(int) })
```