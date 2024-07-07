package goset

import "testing"

var (
	intset1  = NewIntSet(1, -1, 5)
	intset2  = NewIntSet(1)
	intset3  = NewIntSet(10, 45, 3)
	intset4  = NewIntSet(1, -1, 5, 10, 45, 3)
	intset5  = NewIntSet(-1, 5)
	intset6  = NewIntSet(-1, 5)
	intset7  = NewIntSet(1)
	intset8  = NewIntSet(10, 45, 3)
	intset9  = NewIntSet(5)
	intset10 = NewIntSet(5, 6)
	intset11 = NewIntSet()
)

func TestIntSet(t *testing.T) {
	if r := intset1.Length(); r != 3 {
		t.Fatalf("intset1.Length() got unexpected %d", r)
	}
	if r := intset3.Copy(); !r.Equals(intset8) {
		t.Fatalf("intset3.Copy() got unexpected %v", r.ToList())
	}
	if r := intset1.Has(-1); r != true {
		t.Fatalf("intset1.Has(-1) got unexpected %t", r)
	}
	if r := intset1.Has(2); r != false {
		t.Fatalf("intset1.Has(2) got unexpected %t", r)
	}
	if r := intset1.IsSub(intset3); r != false {
		t.Fatalf("intset1.IsSub(intset3) got unexpected %t", r)
	}
	if r := intset2.IsSub(intset1); r != true {
		t.Fatalf("intset2.IsSub(intset1) got unexpected %t", r)
	}
	if r := intset1.Union(intset3); !r.Equals(intset4) {
		t.Fatalf("intset1.Union(intset3) got unexpteced %v", r.ToList())
	}
	if r := intset1.Subtract(intset2); !r.Equals(intset5) {
		t.Fatalf("intset1.Subtract(intset2) got unexpected %v", r.ToList())
	}
	if r := intset1.Complement(intset2); !r.Equals(intset6) {
		t.Fatalf("intset1.Complement(intset2) got unexpteced %v", r.ToList())
	}
	if r := intset1.Intersect(intset2); !r.Equals(intset7) {
		t.Fatalf("intset1.Intersect(intset2) got unexpteced %v", r.ToList())
	}
	intset1.Delete(1, -1)
	if !intset1.Equals(intset9) {
		t.Fatalf("intset1.Delete(1, -1), got unexpected %v", intset1.ToList())
	}
	intset1.Add(6)
	if !intset1.Equals(intset10) {
		t.Fatalf("!intset1.Equals(intset10) got unexpected %v", intset1.ToList())
	}
	intset1.Clear()
	if !intset1.Equals(intset11) {
		t.Fatalf("intset1.Clear() got unexpected %v", intset1.ToList())
	}
}
