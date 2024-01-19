package main

import (
	"fmt"

	"github.com/Visforest/goset/v2"
)

func main() {
	// create an empty IntSet
	var s1 = goset.NewIntSet()

	// add elements
	s1.Add(9, 6, 8, 10, 1)

	// delete elements
	s1.Delete(6)

	// output elements

	// output: [9 8 10 1]
	fmt.Println("output:", s1.ToList())

	// check exists
	// true
	fmt.Println(s1.Has(1))

	// create a deep copy
	var s2 = s1.Copy()
	// s2: [9 8 10 1]
	fmt.Println("s2:", s2.ToList())

	// clear all elements
	s2.Clear()
	// after s2 cleared, s1:  [9 8 10 1]
	fmt.Println("after s2 cleared, s1:", s1.ToList())

	var nums1 = goset.NewIntSet(10, 9, 11, 0, 0)
	var nums2 = goset.NewIntSet(11, 12)

	// union: [0 9 10 11 12]
	fmt.Println("union:", nums1.Union(nums2).ToList())
	// subtract: [0 9 10]
	fmt.Println("subtract:", nums1.Subtract(nums2).ToList())
	// intersect: [11]
	fmt.Println("intersect:", nums1.Intersect(nums2).ToList())
	// complement: [0 9 10 12]
	fmt.Println("complement:", nums1.Complement(nums2).ToList())

	var nums3 = goset.NewIntSet(12, 11)
	var nums4 = goset.NewIntSet(12, 11, 9)
	// true
	fmt.Println(nums2.IsSub(nums4))
	// true
	fmt.Println(nums2.Equals(nums3))
}
