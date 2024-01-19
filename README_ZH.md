[English](README.md) | 中文

一个协程安全的 Golang Set 实现，支持多种数量类型和特性。

goset 基于 Go 泛型实现，因此 Go 版本需 >=1.18.

# 特点

- 协程安全
- 支持元素增、删、包含判断、清空
- 支持集合运算：交、并、差、补
- 支持集合比较：包含、相等
- 支持深拷贝
- 支持多种元素数据类型，如：int8,int16,int32,int64,int,string 等等
- 支持先进先出 Set，先进后出 Set，有序 Set 

# 安装

```
$ go get github.com/Visforest/goset
```

# 使用方法

goset 中的各类 Set 有如下相同的函数：
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

引入 goset:
```go
import "github.com/Visforest/goset/v2"
```

## 普通 Set

指定元素数据类型来使用：
```go
var myset=goset.NewSet[string]("a","b","e")
myset.Add("a","c")
// [ a b e c] 
fmt.Println(myset.ToList())
```

对于 int 和 string 类型的元素，已有预置的 `IntSet` 和 `StrSet` 可供直接使用：
```go
var s1 goset.StrSet
// samsung is in s1? false
fmt.Printf("samsung is in s1? %t", s1.Has("samsung"))
```

随意定制你的 Set 类型，例如：
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

FifoSet 类似于先进先出的队列，只是元素是去重的。

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

FiloSet 类似于先进后出的堆栈，只是元素是去重的。

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

SortedSet 是一个元素有序的 Set。

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

查看 [examples/](examples/) 了解更多用法.

---
欢迎提 issues 和参与进来。