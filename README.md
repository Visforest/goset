English | [中文](README_ZH.md)

A goroutine-safe Set implementation in Golang.

# Features

- Goroutine safe
- Support add/delete/exist/clear elements
- Support set math operations: intersect/union/subtract/complement
- Support set compare: contain/equal
- Support deep copy

# Install

```
$ go get github.com/Visforest/goset
```

# Usages

## Set
Element operations:

```go
// create a new set
var fruits = goset.NewSet("banana", "tomato", "peach")
// add elements
fruits.Add("apple","pear")
// delete elements
fruits.Delete("tomato")
// check whether element exists in set
fruits.Has("apple")
// clear all elements
fruits.Clear()
```

Set view operations:

```go
var fruits = goset.NewSet("banana", "tomato", "peach")
// get elements in form of slice
fruits.ToList() 
// get elements count
fruits.Length()
// get a deep copy of fruits
fruits.Copy()
```

Set math operations:

```go
var fruits = goset.NewSet("banana", "tomato", "peach")
var vegatables = goset.NewSet("tomato", "cabbage")

// fruits,vegatables union: [banana tomato peach cabbage]
fmt.Println("fruits,vegatables union:", fruits.Union(vegatables).ToList())
// fruits,vegatables subtract: [peach banana]
fmt.Println("fruits,vegatables subtract:", fruits.Subtract(vegatables).ToList())
// fruits,vegatables intersect: [tomato]
fmt.Println("fruits,vegatables intersect:", fruits.Intersect(vegatables).ToList())
// fruits,vegatables complement: [peach banana cabbage]
fmt.Println("fruits,vegatables complement:", fruits.Complement(vegatables).ToList())
```

Set compare operations:

```go
var numbers1 = goset.NewSet(1, 3, 0, -3, 5)
var numbers2 = goset.NewSet(0, 3)
var numbers3 = goset.NewSet(3, 0, 3)
// numbers2 is sub set of numbers1 ? true
fmt.Println("numbers2 is sub set of numbers1 ?", numbers2.IsSub(numbers1))
// numbers2 equals numbers3 ? true
fmt.Println("numbers2 equals numbers3 ?", numbers2.Equals(numbers3))
```

## IntSet

IntSet has same methods with Set, but only accecpts `int` elements, and is available to export sorted slice result:

```go
var nums = goset.NewIntSet(9,3,5,7,3,1)
// [1,3,5,7,9]
fmt.Println(nums.Tolist(goset.Asc))
// [9,7,5,3,1]
fmt.Println(nums.Tolist(goset.Desc))
// random order of [1,3,5,7,9]
fmt.Println(nums.Tolist(goset.Random))
// random order of [1,3,5,7,9]
fmt.Println(nums.Tolist())
```