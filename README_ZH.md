一个协程安全的 Golang Set 实现。

# 特点

- 协程安全
- 支持元素增、删、包含判断、清空
- 支持集合运算：交、并、差、补
- 支持集合比较：包含、相等
- 支持深拷贝

# 安装

```
$ go get github.com/Visforest/goset
```

# 使用方法

元素操作：

```go
// create a new set
var fruits = goset.New("banana", "tomato", "peach")
// add elements
fruits.Add("apple","pear")
// delete elements
fruits.Delete("tomato")
// check whether element exists in set
fruits.Has("apple")
// clear all elements
fruits.Clear()
```

集合操作：

```go
var fruits = goset.New("banana", "tomato", "peach")
// get elements in form of slice
fruits.ToList() 
// get elements count
fruits.Length()
// get a deep copy of fruits
fruits.Copy()
```

集合数学运算：

```go
var fruits = goset.New("banana", "tomato", "peach")
var vegatables = goset.New("tomato", "cabbage")

// fruits,vegatables union: [banana tomato peach cabbage]
fmt.Println("fruits,vegatables union:", fruits.Union(vegatables).ToList())
// fruits,vegatables subtract: [peach banana]
fmt.Println("fruits,vegatables subtract:", fruits.Subtract(vegatables).ToList())
// fruits,vegatables intersect: [tomato]
fmt.Println("fruits,vegatables intersect:", fruits.Intersect(vegatables).ToList())
// fruits,vegatables complement: [peach banana cabbage]
fmt.Println("fruits,vegatables complement:", fruits.Complement(vegatables).ToList())
```

集合比较：

```go
var numbers1 = goset.New(1, 3, 0, -3, 5)
var numbers2 = goset.New(0, 3)
var numbers3 = goset.New(3, 0, 3)
// numbers2 is sub set of numbers1 ? true
fmt.Println("numbers2 is sub set of numbers1 ?", numbers2.IsSub(numbers1))
// numbers2 equals numbers3 ? true
fmt.Println("numbers2 equals numbers3 ?", numbers2.Equals(numbers3))
```
