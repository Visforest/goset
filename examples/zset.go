package main

import (
	"fmt"
	"github.com/Visforest/goset"
)

func main() {
	var s1 = goset.NewSortedSet[int]()
	s1.Add(5, 7, 10, 3, -1, 7, 0, 9, 3)
	for _, v := range s1.ToList() {
		fmt.Println(v)
	}

	var s2 = goset.NewSortedSet[string]()
	s2.Add("nokia", "apple", "nubia", "huawei", "honor", "apple", "xiaomi", "samsung", "htc", "oppo", "oneplus", "vivo")
	for _, v := range s2.ToList() {
		fmt.Println(v)
	}
}
