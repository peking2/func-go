# func-go
functional programming in Go

## Installation
```bash
go get github.com/peking2/func-go
```
## Usage
```go
import . "github.com/peking2/func-go"

res := NewArray([]int{1, 2, 3, 4, 5}).
	Map(func(i Any) Any { return i.(int) * 2 }).
	Reduce(func(a Any, b Any) Any { return a.(int) + b.(int) })
```

## GoDoc

```go
CONSTANTS

const (
    ERROR_INVALID_PARAMETER = "Invalid parameters"
)

FUNCTIONS

func WrapFunc(f Any) func(...Any) Try
    Wrap two kinds of functions:

	1. func(...Any) (error)     -> func (...Any) Try
	2. func(...Any) (Any error) -> func (...Any) Try

TYPES

type Any interface{}

type Array []Any

func NewArray(args ...Any) *Array
    Create new Array. Examples:

	1. NewArray()
	2. NewArray(1,2,3)
	3. NewArray([]int {1,2,3})

func (arr *Array) Filter(lambda func(Any) bool) *Array

func (arr *Array) ForEach(lambda func(Any))

func (arr *Array) Map(lambda func(Any) Any) *Array

func (arr *Array) Reduce(lambda func(Any, Any) Any) Any

type Map map[Any]Any

func NewMap(args ...Any) *Map
    Create new Map. Examples:

	1. NewMap()
	2. NewMap(map[string]string {"name": "Mike"})

func (m *Map) ForEach(lambda func(Any, Any))

type Try struct {
    // contains filtered or unexported fields
}

func NewTry(v Any) *Try

func (t *Try) FlatMap(lambda func(Any) Try) *Try

func (t *Try) Get() Any

func (t *Try) IsFailure() bool

func (t *Try) IsSuccess() bool
```