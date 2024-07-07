package main

import (
	"fmt"
	"github.com/visforest/goset/v2"
)

func main() {
	var s1 = goset.NewSortedSet[int](10, 5, 2)
	var s2 = goset.NewSortedSet[int](10, 0, -4, 20, 5, 2)
	s1.Add(11, 7)
	var s3 = s1.Copy()
	fmt.Println("s1:", s1.ToList())
	fmt.Println("s2:", s2.ToList())
	fmt.Println("s3:", s3.ToList())
	fmt.Println("s1 is sub of s2:", s1.IsSub(s2))
	fmt.Println("s1 union s2:", s1.Union(s2).ToList())
	fmt.Println("s1 intersect s2:", s1.Intersect(s2).ToList())
	fmt.Println("s1 complement s2:", s1.Complement(s2).ToList())
	fmt.Println("s1 subtract s2:", s1.Subtract(s2).ToList())
}
