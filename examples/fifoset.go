package main

import (
	"fmt"

	"github.com/Visforest/goset"
)

func main() {
	var s = goset.NewFifoSet[string]()
	s.Add("e", "a", "b", "a", "c", "b")
	s.Delete("b")
	// e
	// a
	// c
	for _, v := range s.ToList() {
		fmt.Println(v)
	}
}
