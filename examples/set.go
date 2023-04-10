package main

import (
	"fmt"

	"github.com/Visforest/goset"
)

func main() {
	var s1 = goset.New()
	s1.Add("a", "b", "c")
	// after add a,b,c: [c a b]
	fmt.Println("after add a,b,c:", s1.ToList())

	s1.Delete("a", "b")
	// after delete a b: [c]
	fmt.Println("after delete a b:", s1.ToList())

	// s1 has c ? true
	fmt.Println("s1 has c ?", s1.Has("c"))
	// s1 has d ? false
	fmt.Println("s1 has d ?", s1.Has("d"))

	var s2 = s1.Copy()

	s1.Clear()
	// after clear: []
	fmt.Println("after clear:", s1.ToList())
	// s1's copy after s1 cleared: [c]
	fmt.Println("s1's copy after s1 cleared:", s2.ToList())

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

	var numbers1 = goset.New(1, 3, 0, -3, 5)
	var numbers2 = goset.New(0, 3)
	var numbers3 = goset.New(3, 0, 3)
	// numbers2 is sub set of numbers1 ? true
	fmt.Println("numbers2 is sub set of numbers1 ?", numbers2.IsSub(numbers1))
	// numbers2 equals numbers3 ? true
	fmt.Println("numbers2 equals numbers3 ?", numbers2.Equals(numbers3))
}
