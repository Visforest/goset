package goset

import (
	"reflect"
	"testing"
)

func TestIntSet(t *testing.T) {
	t.Run("test intset add", func(t *testing.T) {
		var set = NewIntSet()
		set.Add(1, 8, 0, 8, 9)

		var expect1 = []int{0, 1, 8, 9}
		var result1 = set.ToList(Asc)
		if !reflect.DeepEqual(expect1, result1) {
			t.Errorf("to asc list,expected:%v,got %v", expect1, result1)
		}

		var expect2 = []int{9, 8, 1, 0}
		var result2 = set.ToList(Desc)
		if !reflect.DeepEqual(expect2, result2) {
			t.Errorf("to desc list,expected:%v,got %v", expect2, result2)
		}

		var expect3 = map[int]int{0: 0, 1: 0, 8: 0, 9: 0}
		var result3 = set.ToList()
		if len(expect3) != len(result3) {
			t.Errorf("to no-order list,got %v", result3)
		}
		for _, v := range result3 {
			if _, ok := expect3[v]; !ok {
				t.Errorf("to no-order list,got %v", result3)
				break
			}
		}
	})

	t.Run("test intset delete", func(t *testing.T) {
		var set1 = NewIntSet(1, 0, 8, 9)
		set1.Delete(8, 9)

		var result1 = set1.ToList(Asc)
		var expect1 = []int{0, 1}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("set1 delete expected %v, got %v", expect1, result1)
		}

		var set2 = NewIntSet(1, 0)
		set2.Delete(0, 1, 2, 3, 4)
		var result2 = set2.ToList(Asc)
		var expect2 = []int{}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("set2 delete expected %v, got %v", expect2, result2)
		}
		if set2.Length() != 0 {
			t.Errorf("set2 delete, length expected %d, got %d", 0, set2.Length())
		}
	})

	t.Run("test intset clear", func(t *testing.T) {
		var set = NewIntSet(1, 0, 8, 9)
		set.Clear()

		var result = set.ToList(Asc)
		var expect = []int{}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("clear expected %v, got %v", expect, result)
		}
		if set.Length() != 0 {
			t.Errorf("clear, length expected %d, got %d", 0, set.Length())
		}
	})

	t.Run("test intset has", func(t *testing.T) {
		var set = NewIntSet(0, 1)

		if !set.Has(1) {
			t.Errorf("set has 1, expected true, got false")
		}
		if set.Has(2) {
			t.Errorf("set has 2, expected false, got true")
		}
	})

	t.Run("test intset length", func(t *testing.T) {
		var set = NewIntSet(1, 0, 8, 9)
		set.Add(8, 10)

		result := set.Length()
		expect := 5
		if result != expect {
			t.Errorf("length expected %d, got %d", expect, result)
		}
	})

	t.Run("test intset equal", func(t *testing.T) {
		var set1 = NewIntSet(1, 0, 8, 9)
		var set2 = NewIntSet(8, 9, 0, 1)
		var set3 = NewIntSet(8, 9, 0, 10)
		var set4 = NewIntSet()

		if !set1.Equals(set2) {
			t.Errorf("set1==set2, expected true, got false")
		}
		if !set1.Equals(set1) {
			t.Errorf("set1==set1, expected true, got false")
		}
		if set1.Equals(set3) {
			t.Errorf("set1==set3, expected false, got true")
		}
		if set1.Equals(set4) {
			t.Errorf("set1==set4, expected false, got true")
		}
		if set1.Equals(nil) {
			t.Errorf("set1==nil, expected false, got true")
		}
	})

	t.Run("test intset is sub", func(t *testing.T) {
		var set1 = NewIntSet(1, 0, 8, 9)
		var set2 = NewIntSet(8, 9, 0, 1, 10)
		var set3 = NewIntSet(8, 9, 0, 10)
		var set4 = NewIntSet()
		if !set1.IsSub(set2) {
			t.Errorf("set1 is sub of set2, expected true, got false")
		}
		if set1.IsSub(set3) {
			t.Errorf("set1 is sub of set3, expected false, got true")
		}
		if !set1.IsSub(set1) {
			t.Errorf("set1 is sub of set1, expected true, got false")
		}
		if set1.IsSub(set4) {
			t.Errorf("set1 is sub of set4, expected false, got true")
		}
		if set1.IsSub(nil) {
			t.Errorf("set1 is sub of nil, expected false, got true")
		}
	})

	t.Run("test intset union", func(t *testing.T) {
		var set1 = NewIntSet(1, 0, 8, 9)
		var set2 = NewIntSet(4, 5, 8, 7)
		var set3 = NewIntSet()

		var result1 = set1.Union(set2).ToList(Asc)
		var expect1 = []int{0, 1, 4, 5, 7, 8, 9}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("set1 union set2, expected %v, got %v", expect1, result1)
		}

		var result2 = set1.Union(set3).ToList(Asc)
		var expect2 = []int{0, 1, 8, 9}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("set1 union set3, expected %v, got %v", expect2, result2)
		}

		var result3 = set1.Union(nil).ToList()
		var expect3 = []int{0, 1, 8, 9}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("set1 union nil ,expected %v, got %v", expect3, result3)
		}
	})

	t.Run("test intset intersect", func(t *testing.T) {
		var set1 = NewIntSet(1, 0, 8, 9)
		var set2 = NewIntSet(4, 5, 8, 7)
		var set3 = NewIntSet(10, 11)
		var set4 = NewIntSet()

		var result1 = set1.Intersect(set2).ToList(Asc)
		var expect1 = []int{8}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("set1 intersect set2, expected %v, got %v", expect1, result1)
		}

		var result2 = set1.Intersect(set3).ToList(Asc)
		var expect2 = []int{}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("set1 intersect set3, expected %v, got %v", expect2, result2)
		}

		var result3 = set1.Intersect(set4).ToList(Asc)
		var expect3 = []int{}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("set1 intersect set4, expected %v, got %v", expect3, result3)
		}

		var result4 = set1.Intersect(nil).ToList(Asc)
		var expect4 = []int{}
		if !reflect.DeepEqual(result4, expect4) {
			t.Errorf("set1 intersect nil, expected %v, got %v", expect4, result4)
		}
	})

	t.Run("test intset substract", func(t *testing.T) {
		var set1 = NewIntSet(1, 0, 8, 9)
		var set2 = NewIntSet(4, 5, 8, 7)
		var set3 = NewIntSet(10, 11)
		var set4 = NewIntSet()

		var result1 = set1.Subtract(set2).ToList(Asc)
		var expect1 = []int{0, 1, 9}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("set1 subtract set2, expected %v, got %v", expect1, result1)
		}

		var result2 = set1.Subtract(set3).ToList(Asc)
		var expect2 = []int{0, 1, 8, 9}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("set1 subtract set3, expected %v, got %v", expect2, result2)
		}

		var result3 = set1.Subtract(set4).ToList(Asc)
		var expect3 = []int{0, 1, 8, 9}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("set1 subtract set4, expected %v, got %v", expect3, result3)
		}

		var result4 = set1.Subtract(nil).ToList(Asc)
		var expect4 = []int{0, 1, 8, 9}
		if !reflect.DeepEqual(result4, expect4) {
			t.Errorf("set1 subtract nil, expected %v, got %v", expect4, result4)
		}
	})

	t.Run("test intset complement", func(t *testing.T) {
		var set1 = NewIntSet(1, 0, 8, 9)
		var set2 = NewIntSet(4, 5, 8, 7)
		var set3 = NewIntSet(10, 11)
		var set4 = NewIntSet()

		var result1 = set1.Complement(set2).ToList(Asc)
		var expect1 = []int{0, 1, 4, 5, 7, 9}
		if !reflect.DeepEqual(result1, expect1) {
			t.Errorf("set1 complement set2, expected %v, got %v", expect1, result1)
		}

		var result2 = set1.Complement(set3).ToList(Asc)
		var expect2 = []int{0, 1, 8, 9, 10, 11}
		if !reflect.DeepEqual(result2, expect2) {
			t.Errorf("set1 complement set3, expected %v, got %v", expect2, result2)
		}

		var result3 = set1.Complement(set4).ToList(Asc)
		var expect3 = []int{0, 1, 8, 9}
		if !reflect.DeepEqual(result3, expect3) {
			t.Errorf("set1 complement set4, expected %v, got %v", expect3, result3)
		}

		var result4 = set1.Complement(nil).ToList(Asc)
		var expect4 = []int{0, 1, 8, 9}
		if !reflect.DeepEqual(result4, expect4) {
			t.Errorf("set1 complement nil, expected %v, got %v", expect4, result4)
		}

	})
}
