package main

import (
	"fmt"

	"github.com/Visforest/goset"
)

func main() {
	var s = goset.NewFiloSet[string]()
	s.Add("e", "a", "b", "a", "c", "b")
	s.Delete("b")
	// c
	// a
	// e
	for _, v := range s.ToList() {
		fmt.Println(v)
	}
}
