package main

import (
	"fmt"
	"github.com/Visforest/goset"
)

func main() {
	var s1 goset.StrSet
	fmt.Printf("s1 len:%d \n", s1.Length())

	fmt.Printf("samsung is in s1? %t \n", s1.Has("samsung"))

	s1.Add("apple", "huawei")
	fmt.Printf("s1: %v \n", s1.ToList())
	fmt.Printf("apple is in s1? %t \n", s1.Has("apple"))

	var s2 = goset.NewStrSet()
	s2.Add("xiaomi", "apple")

	fmt.Printf("s2: %v \n", s2.ToList())

	s3 := s1.Union(s2)
	fmt.Printf("s1 union s2,got:%v \n", s3.ToList())

	s4 := s1.Intersect(s2)
	fmt.Printf("s1 intersect s2,got:%v \n", s4.ToList())

	s5 := s1.Subtract(s2)
	fmt.Printf("s1 subtract s2,got:%v \n", s5.ToList())

	s6 := s1.Complement(s2)
	fmt.Printf("s1 complement s2,got:%v \n", s6.ToList())

	s7 := s1.Copy()
	fmt.Printf("s1's copy, s7:%v \n", s7.ToList())
	fmt.Printf("s1 equals s7? %v \n", s1.Equals(s7))

	s7.Delete("apple")
	s7.Add("samsung")
	fmt.Printf("s1 equals s7? %v \n", s1.Equals(s7))
	fmt.Printf("s7:%v \n", s7.ToList())
}
