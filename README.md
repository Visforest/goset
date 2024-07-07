English | [中文](README_ZH.md)

A goroutine-safe Set implementation in Golang, supports various data types and features. 

It's implemented on Go generics features, so your Go version needs to be >=1.18. 

# Features

- Goroutine safe
- Support add/delete/exist/clear elements
- Support set math operations: intersect/union/subtract/complement
- Support set compare: contain/equal
- Support deep copy
- Support various kinds of elements: int8,int16,int32,int64,int,string, and so on
- Support fifo set,filo set,sorted set

# Install

```
$ go get github.com/visforest/goset/v2
```

# Usages

All sets have common functions:
- Add(...vals)
- Delete(...vals)
- Clear()
- Length() int
- Has(v) bool
- Copy() *set
- ToList() []type
- Equals(*set) bool
- IsSub(*set) bool
- Union(*set) *set
- Intersect(*set) *set
- Subtract(*set) *set
- Complement(*set) *set


Import goset:
```go
import "github.com/visforest/goset/v2"
```

## Normal set

specify data type
```go
var myset=goset.NewSet[string]("a","b","e")
myset.Add("a","c")
// [ a b e c] 
fmt.Println(myset.ToList())
```

For integer and string elements, `IntSet`,`Int64Set`,`UintSet`,`Uint64Set` and `StrSet` are predefined already:
```go
var s1 goset.StrSet
// samsung is in s1? false
fmt.Printf("samsung is in s1? %t", s1.Has("samsung"))
```

Feel free to customize your Set type, for example:
```go
type user struct {
	name string
	age  int
}
type userSet = goset.Set[user]

func main() {
	s := userSet{Data: make(map[user]struct{})}
	s.Add(
		user{
			name: "Mickey",
			age:  10,
		},
		user{
			name: "Tom",
			age:  20,
		},
		user{
			name: "Mickey",
			age:  10,
		},
		user{
			name: "Tiana",
			age:  21,
		},
	)
	// {Mickey 10}
	// {Tom 20}
	// {Tiana 21}
	for _, u := range s.ToList() {
		fmt.Println(u)
	}
}
```

## FifoSet

FifoSet is like a fifo queue, but elements are deduplicated.

```go
var s = goset.NewFifoSet[string]()
s.Add("e", "a", "b", "a", "c", "b")
s.Delete("b")
// e
// a
// c
for _, v := range s.ToList() {
    fmt.Println(v)
}
```

## FiloSet

FiloSet is like a filo stack, but elements are deduplicated.

```go
var s = goset.NewFiloSet[string]()
s.Add("e", "a", "b", "a", "c", "b")
s.Delete("b")
// c
// a
// e
for _, v := range s.ToList() {
    fmt.Println(v)
}
```

## SortedSet

SortedSet is a set whose elements are stored in asc order.

```go
var s1 = goset.NewSortedSet[int]()
s1.Add(5, 7, 10, 3, -1, 7, 0, 9, 3)
// -1
// 0
// 3
// 5
// 7
// 9
// 10
for _, v := range s1.ToList() {
    fmt.Println(v)
}
```

Read [examples/](examples/) to learn more.

---

Feel free to make issues and pull request.