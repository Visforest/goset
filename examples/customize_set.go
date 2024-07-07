package main

import (
	"fmt"
	"github.com/visforest/goset/v2"
)

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
