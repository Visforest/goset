package goset

import (
	"reflect"
	"sync"
)

type StrSet struct {
	m    sync.RWMutex
	data map[string]struct{}
}

// NewStrSet creates a new StrSet
func NewStrSet(v ...string) *StrSet {
	s := &StrSet{data: map[string]struct{}{}}
	if len(v) > 0 {
		s.Add(v...)
	}
	return s
}

// Add adds elements
func (s *StrSet) Add(v ...string) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.data == nil {
		s.data = make(map[string]struct{})
	}
	for _, ele := range v {
		s.data[ele] = struct{}{}
	}
}

// Delete delete elements
func (s *StrSet) Delete(v ...string) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.data == nil {
		s.data = make(map[string]struct{})
	}
	for _, ele := range v {
		delete(s.data, ele)
	}
}

// Clear clears all elements
func (s *StrSet) Clear() {
	defer s.m.Unlock()
	s.m.Lock()
	s.data = make(map[string]struct{})
}

// Copy returns a deep copy of itself
func (s *StrSet) Copy() *StrSet {
	defer s.m.RUnlock()
	s.m.RLock()
	if s.data == nil {
		s.data = make(map[string]struct{})
	}

	data := make(map[string]struct{})
	for v := range s.data {
		data[v] = struct{}{}
	}
	return &StrSet{data: data}
}

// Length returns StrSet length
func (s *StrSet) Length() int {
	return len(s.data)
}

// Has returns whether v exists in StrSet
func (s *StrSet) Has(v string) bool {
	if s.data == nil {
		s.data = make(map[string]struct{})
	}
	_, ok := s.data[v]
	return ok
}

// ToList returns data slice
func (s *StrSet) ToList() []string {
	defer s.m.RUnlock()
	s.m.RLock()

	if s.data == nil {
		s.data = make(map[string]struct{})
	}

	var data = make([]string, s.Length())
	var i int
	for d := range s.data {
		data[i] = d
		i++
	}
	return data
}

// Equals returns whether StrSet s has the same members with StrSet t
func (s *StrSet) Equals(t *StrSet) bool {
	if t == nil {
		return false
	}
	return reflect.DeepEqual(s.data, t.data)
}

// IsSub returns whether it's a part of StrSet t
func (s *StrSet) IsSub(t *StrSet) bool {
	if t == nil {
		return false
	}
	if s == t {
		// compare with itself
		return true
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	if s.data == nil {
		s.data = make(map[string]struct{})
	}
	if t.data == nil {
		t.data = make(map[string]struct{})
	}

	if s.Length() > t.Length() {
		return false
	}
	for v := range s.data {
		if !t.Has(v) {
			return false
		}
	}
	return true
}

// Union unions with StrSet t and returns a new StrSet
//
// for example:
// var a=NewStrSet("1","2","3")
// var b=NewStrSet("2","3","4")
// a.Union(b) returns {"1","2","3","4"}
func (s *StrSet) Union(t *StrSet) *StrSet {
	var r = s.Copy()
	if t == nil || s == t {
		return r
	}
	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	r.Add(t.ToList()...)
	return r
}

// Intersect returns a new StrSet Whose elements exist in both StrSet
//
// for example:
// var a=NewStrSet("1","2","3")
// var b=NewStrSet("2","3","4")
// a.Intersect(b) returns {"2","3"}
func (s *StrSet) Intersect(t *StrSet) *StrSet {
	var r = NewStrSet()
	if t == nil || s.Length() == 0 || t.Length() == 0 {
		return r
	}
	if s == t {
		// intersect itself
		return s.Copy()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	if s.data == nil {
		s.data = make(map[string]struct{})
	}
	if t.data == nil {
		t.data = make(map[string]struct{})
	}

	if s.Length() >= t.Length() {
		for v := range t.data {
			if s.Has(v) {
				r.Add(v)
			}
		}
	} else {
		for v := range s.data {
			if t.Has(v) {
				r.Add(v)
			}
		}
	}
	return r
}

// Subtract returns a new StrSet Whose elements exist in itself but don't exist in StrSet t
//
// for example:
// var a=NewStrSet("1","2","3")
// var b=NewStrSet("2","3","4")
// a.Subtract(b) returns {"1"}
func (s *StrSet) Subtract(t *StrSet) *StrSet {
	if t == nil || t.Length() == 0 {
		return s.Copy()
	}
	var r = NewStrSet()
	if s == t {
		// subtract itself
		return r
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	if s.data == nil {
		s.data = make(map[string]struct{})
	}
	if t.data == nil {
		t.data = make(map[string]struct{})
	}

	for v := range s.data {
		if !t.Has(v) {
			r.Add(v)
		}
	}
	return r
}

// Complement returns a new StrSet Whose elements only exists in one StrSet
//
// for example:
// var a=NewStrSet("1","2","3")
// var b=NewStrSet("2","3","4")
// a.Complement(b) returns {"1","4"}
func (s *StrSet) Complement(t *StrSet) *StrSet {
	if t == nil || s.Length() == 0 || t.Length() == 0 {
		return s.Copy()
	}

	if s == t {
		return NewStrSet()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()
	if s.data == nil {
		s.data = make(map[string]struct{})
	}
	if t.data == nil {
		t.data = make(map[string]struct{})
	}

	var r = s.Union(t)
	if s.Length() >= t.Length() {
		for v := range t.data {
			if s.Has(v) {
				r.Delete(v)
			}
		}
	} else {
		for v := range s.data {
			if t.Has(v) {
				r.Delete(v)
			}
		}
	}
	return r
}
