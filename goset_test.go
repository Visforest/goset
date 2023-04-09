package goset

import (
	"fmt"
	"testing"
)

func TestGoSet(t *testing.T) {
	t.Run("test set add", func(t *testing.T) {
		var set = New()
		set.Add(1, 2, 3, 4, 1)
		fmt.Println("add 1,2,3,4,1 -->", set.ToList())
	})

	t.Run("test set delete", func(t *testing.T) {
		var setA = New(1, 2, 3, 4)
		setA.Delete(3, 4)
		var setB = New(1, 2, 3, 4)
		setB.Delete(4, 5)
		fmt.Println("{1,2,3,4} delete {3,4} -->", setA.ToList())
		fmt.Println("{1,2,3,4} delete {4,5} -->", setB.ToList())
	})

	t.Run("test set clear", func(t *testing.T) {
		var set = New(1, 2, 3, 4)
		set.Clear()
		fmt.Println("{1,2,3,4} clear -->", set.ToList())
	})

	t.Run("test set length", func(t *testing.T) {
		var set = New(1, 2, 3, 4, 4)
		set.Clear()
		fmt.Println("{1,2,3,4,4} length -->", set.Length())
	})

	t.Run("test set copy", func(t *testing.T) {
		var set = New(1, 2, 3, 4)
		newSet := set.Copy()
		fmt.Println("{1,2,3,4} copy -->", newSet.ToList())
	})

	t.Run("test set has", func(t *testing.T) {
		var set = New(1, 2, 3, 4)
		fmt.Println("{1,2,3,4} has 1 ? -->", set.Has(1))
		fmt.Println("{1,2,3,4} has 5 ? -->", set.Has(5))
		fmt.Println("{1,2,3,4} has nil ? -->", set.Has(nil))
	})

	t.Run("test set equal", func(t *testing.T) {
		var setA = New(1, 2, 3, 4)
		var setB = New(3, 4, 1, 2)
		var setC = New(9, 10)
		var setD *Set
		fmt.Println("{1,2,3,4} equals {3,4,1,2} ? -->", setA.Equals(setB))
		fmt.Println("{1,2,3,4} equals {9,10} ? -->", setA.Equals(setC))
		fmt.Println("{1,2,3,4} equals nil ? -->", setA.Equals(setD))
		fmt.Println("{1,2,3,4} equals itself ? -->", setA.Equals(setA))
	})

	t.Run("test set isSub", func(t *testing.T) {
		var setA = New(1, 2, 3, 4)
		var setB = New(1, 2, 3, 4, 5)
		var setC = New(9, 10)
		var setD *Set
		fmt.Println("{1,2,3,4} is sub of {1,2,3,4,5} ? -->", setA.IsSub(setB))
		fmt.Println("{1,2,3,4} is sub of {9,10} ? -->", setA.IsSub(setC))
		fmt.Println("{1,2,3,4} is sub of nil ? -->", setA.IsSub(setD))
		fmt.Println("{1,2,3,4} is sub of itself ? -->", setA.IsSub(setA))
	})

	t.Run("test set union", func(t *testing.T) {
		var setA = New(1, 2, 3, 4)
		var setB = New(4, 5)
		var setC *Set
		fmt.Println("{1,2,3,4} union {4,5} ? -->", setA.Union(setB).ToList())
		fmt.Println("{1,2,3,4} union nil ? -->", setA.Union(setC).ToList())
		fmt.Println("{1,2,3,4} union itself ? -->", setA.Union(setA).ToList())
	})

	t.Run("test set subtract", func(t *testing.T) {
		var setA = New(1, 2, 3, 4)
		var setB = New(4, 5)
		var setC = New(8)
		var setD *Set
		fmt.Println("{1,2,3,4} subtract {4,5} ? -->", setA.Subtract(setB).ToList())
		fmt.Println("{1,2,3,4} subtract {8} ? -->", setA.Subtract(setC).ToList())
		fmt.Println("{1,2,3,4} subtract nil ? -->", setA.Subtract(setD).ToList())
		fmt.Println("{1,2,3,4} subtract itself ? -->", setA.Subtract(setA).ToList())
	})

	t.Run("test set intersect", func(t *testing.T) {
		var setA = New(1, 2, 3, 4)
		var setB = New(4, 5)
		var setC = New(8)
		var setD *Set
		fmt.Println("{1,2,3,4} intersect {4,5} ? -->", setA.Intersect(setB).ToList())
		fmt.Println("{1,2,3,4} intersect {8} ? -->", setA.Intersect(setC).ToList())
		fmt.Println("{1,2,3,4} intersect nil ? -->", setA.Intersect(setD).ToList())
		fmt.Println("{1,2,3,4} intersect itself ? -->", setA.Intersect(setA).ToList())
	})

	t.Run("test set complement", func(t *testing.T) {
		var setA = New(1, 2, 3, 4)
		var setB = New(4, 5)
		var setC = New(8)
		var setD *Set
		fmt.Println("{1,2,3,4} complement {4,5} ? -->", setA.Complement(setB).ToList())
		fmt.Println("{1,2,3,4} complement {8} ? -->", setA.Complement(setC).ToList())
		fmt.Println("{1,2,3,4} complement nil ? -->", setA.Complement(setD).ToList())
		fmt.Println("{1,2,3,4} complement itself ? -->", setA.Complement(setA).ToList())
	})

	t.Run("test concurrently add & write", func(t *testing.T) {
		var set = New()
		for i := 0; i < 1000; i++ {
			go func() {
				set.Add(1, 2, 3)
			}()
		}
		for i := 0; i < 1000; i++ {
			go func() {
				set.Delete(1, 2, 3)
			}()
		}
		fmt.Println("empty set concurrently add 1,2,3 and delete 1,2,3 -->", set.ToList())
	})

	t.Run("test concurrently add & union with others", func(t *testing.T) {
		var setA = New()
		for i := 0; i < 1000; i++ {
			go func() {
				setA.Add(1, 2, 3)
			}()
		}
		for i := 0; i < 1000; i++ {
			go func() {
				setA.Delete(3, 4)
			}()
		}
		var setB = New(0)
		var setC *Set
		for i := 0; i < 1000; i++ {
			go func() {
				setC = setA.Union(setB)
			}()
		}
		fmt.Println("empty set A concurrently write and union with set B {0}, got set C -->", setC.ToList())
	})
}
