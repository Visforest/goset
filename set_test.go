package goset

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func getSortedInt(data []interface{}) []int {
	var result = make([]int, len(data))
	for i := range result {
		result[i] = data[i].(int)
	}
	sort.Ints(result)
	return result
}

func TestGoSet(t *testing.T) {
	t.Run("test set add", func(t *testing.T) {
		var set = NewSet()
		set.Add(1, 2, 3, 4, 1)

		var expect = []int{1, 2, 3, 4}
		var result = getSortedInt(set.ToList())
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("set expected:%v, got %v", expect, result)
		}
	})

	t.Run("test set delete", func(t *testing.T) {
		var setA = NewSet(1, 2, 3, 4)
		setA.Delete(3, 4)
		var setB = NewSet(1, 2, 3, 4)
		setB.Delete(4, 5, 1, 2, 3)

		var result1 = getSortedInt(setA.ToList())
		var expect1 = []int{1, 2}

		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("setA expected:%v, got %v", expect1, result1)
		}

		var result2 = getSortedInt(setB.ToList())
		var expect2 = []int{}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("setB expected:%v, got %v", expect2, result2)
		}
	})

	t.Run("test set clear", func(t *testing.T) {
		var set = NewSet(1, 2, 3, 4)
		set.Clear()

		var result = getSortedInt(set.ToList())
		var expect = []int{}

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("set expected %v, got %v", expect, result)
		}
	})

	t.Run("test set length", func(t *testing.T) {
		var set = NewSet(1, 2, 3, 4, 4)
		if set.Length() != 4 {
			t.Errorf("set expected %d, got %d", 4, set.Length())
		}
	})

	t.Run("test set copy", func(t *testing.T) {
		var set = NewSet(1, 2, 3, 4)
		newSet := set.Copy()

		var result1 = getSortedInt(newSet.ToList())
		var expect1 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("the copy expected:%v, got %v", result1, expect1)
		}

		newSet.Delete(2)
		var result2 = getSortedInt(set.ToList())
		var expect2 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("the copy updated, set expected:%v, got %v", expect2, result2)
		}
	})

	t.Run("test set has", func(t *testing.T) {
		var set = NewSet("a", "b", "c", "d")
		if !set.Has("a") {
			t.Errorf("has a, expected true, got false")
		}
		if set.Has("x") {
			t.Errorf("has x, expected false, got true")
		}
		if set.Has(nil) {
			t.Errorf("has nil, expected false, got true")
		}
	})

	t.Run("test set equal", func(t *testing.T) {
		var setA = NewSet(1, 2, 3, 4)
		var setB = NewSet(3, 4, 1, 2)
		var setC = NewSet(9, 10)
		var setD = NewSet()

		if !setA.Equals(setB) {
			t.Error("setA==setB, expected true, got false")
		}

		if setA.Equals(setC) {
			t.Error("setA==setC, expected false, got true")
		}

		if setA.Equals(setD) {
			t.Error("setA==setD, expected false, got true")
		}

		if setA.Equals(nil) {
			t.Error("setA==setC, expected false, got true")
		}
	})

	t.Run("test set isSub", func(t *testing.T) {
		var setA = NewSet(1, 2, 3, 4)
		var setB = NewSet(1, 2, 3, 4, 5)
		var setC = NewSet(9, 10)
		var setD = NewSet()

		if !setA.IsSub(setB) {
			t.Errorf("setA is sub of setB, expected true, got false")
		}

		if setA.IsSub(setC) {
			t.Errorf("setA is sub of setC, expected false, got true")
		}

		if setA.IsSub(setD) {
			t.Errorf("setA is sub of setD, expected false, got true")
		}

		if setA.IsSub(nil) {
			t.Errorf("setA is sub of nil, expected false, got true")
		}
	})

	t.Run("test set union", func(t *testing.T) {
		var setA = NewSet(1, 2, 3, 4)
		var setB = NewSet(4, 5)
		var setC = NewSet()

		var result1 = getSortedInt(setA.Union(setB).ToList())
		var expect1 = []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("setA union setB, expected:%v, got %v", expect1, result1)
		}

		var result2 = getSortedInt(setA.Union(setC).ToList())
		var expect2 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("setA union setC, expected:%v, got %v", expect2, result2)
		}

		var result3 = getSortedInt(setA.Union(setA).ToList())
		var expect3 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("setA union setA, expected:%v, got %v", expect3, result3)
		}

		var result4 = getSortedInt(setA.Union(nil).ToList())
		var expect4 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result4, expect4) {
			t.Errorf("setA union nil, expected:%v, got %v", expect4, result4)
		}
	})

	t.Run("test set subtract", func(t *testing.T) {
		var setA = NewSet(1, 2, 3, 4)
		var setB = NewSet(4, 5)
		var setC = NewSet(8)
		var setD = NewSet()

		var result1 = getSortedInt(setA.Subtract(setB).ToList())
		var expect1 = []int{1, 2, 3}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("setA subtract setB, expected:%v, got %v", expect1, result1)
		}

		var result2 = getSortedInt(setA.Subtract(setC).ToList())
		var expect2 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("setA subtract setC, expected:%v, got %v", expect2, result2)
		}

		var result3 = getSortedInt(setA.Subtract(setD).ToList())
		var expect3 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("setA subtract setD, expected:%v, got %v", expect3, result3)
		}

		var result4 = getSortedInt(setA.Subtract(setA).ToList())
		var expect4 = []int{}
		if !reflect.DeepEqual(result4, expect4) {
			t.Errorf("setA subtract setA, expected:%v, got %v", expect4, result4)
		}

		var result5 = getSortedInt(setA.Subtract(nil).ToList())
		var expect5 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result5, expect5) {
			t.Errorf("setA subtract nil, expected:%v, got %v", expect5, result5)
		}
	})

	t.Run("test set intersect", func(t *testing.T) {
		var setA = NewSet(1, 2, 3, 4)
		var setB = NewSet(4, 5)
		var setC = NewSet(8)
		var setD = NewSet()

		var result1 = getSortedInt(setA.Intersect(setB).ToList())
		var expect1 = []int{4}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("setA intersect setB, expected:%v, got %v", expect1, result1)
		}

		var result2 = getSortedInt(setA.Intersect(setC).ToList())
		var expect2 = []int{}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("setA intersect setC, expected:%v, got %v", expect2, result2)
		}

		var result3 = getSortedInt(setA.Intersect(setD).ToList())
		var expect3 = []int{}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("setA intersect setD, expected:%v, got %v", expect3, result3)
		}

		var result4 = getSortedInt(setA.Intersect(setA).ToList())
		var expect4 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result4, expect4) {
			t.Errorf("setA intersect setA, expected:%v, got %v", expect4, result4)
		}

		var result5 = getSortedInt(setA.Intersect(nil).ToList())
		var expect5 = []int{}
		if !reflect.DeepEqual(result5, expect5) {
			t.Errorf("setA intersect nil, expected:%v, got %v", expect5, result5)
		}
	})

	t.Run("test set complement", func(t *testing.T) {
		var setA = NewSet(1, 2, 3, 4)
		var setB = NewSet(4, 5)
		var setC = NewSet(8)
		var setD = NewSet()

		var result1 = getSortedInt(setA.Complement(setB).ToList())
		var expect1 = []int{1, 2, 3, 5}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("setA complement setB, expected:%v, got %v", expect1, result1)
		}

		var result2 = getSortedInt(setA.Complement(setC).ToList())
		var expect2 = []int{1, 2, 3, 4, 8}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("setA complement setC, expected:%v, got %v", expect2, result2)
		}

		var result3 = getSortedInt(setA.Complement(setD).ToList())
		var expect3 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("setA complement setD, expected:%v, got %v", expect3, result3)
		}

		var result4 = getSortedInt(setA.Complement(setA).ToList())
		var expect4 = []int{}
		if !reflect.DeepEqual(result4, expect4) {
			t.Errorf("setA complement setA, expected:%v, got %v", expect4, result4)
		}

		var result5 = getSortedInt(setA.Complement(nil).ToList())
		var expect5 = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result5, expect5) {
			t.Errorf("setA complement nil, expected:%v, got %v", expect5, result5)
		}
	})

	t.Run("test concurrently read & write", func(t *testing.T) {
		var set = NewSet()
		for i := 0; i < 1000; i++ {
			go func() {
				set.Add(rand.Int(), rand.Int())
			}()
			go func() {
				set.Delete(rand.Int(), rand.Int())

			}()
			go func() {
				set.Union(NewSet(1, 2, 3))
			}()
			go func() {
				set.Intersect(NewSet(1, 2, 3))
			}()
		}
	})
}
