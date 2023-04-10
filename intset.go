package goset

import (
	"sort"
	"sync"
)

type SortOption int8

const (
	Random SortOption = iota
	Asc
	Desc
)

type intNode struct {
	val  int
	pre  *intNode
	next *intNode
}

type IntSet struct {
	m    sync.RWMutex
	data map[int]*intNode
	head *intNode
	tail *intNode
}

// NewIntSet creates a new IntSet
func NewIntSet(vals ...int) *IntSet {
	var s = &IntSet{data: make(map[int]*intNode)}
	if len(vals) > 0 {
		s.Add(vals...)
	}
	return s
}

// Add add elements
func (s *IntSet) Add(vals ...int) {
	if len(vals) == 0 {
		return
	}
	sort.Ints(vals)
	s.m.Lock()
	defer s.m.Unlock()

	cur := s.head
	// put sorted elements into sorted double-linkedlist
	for _, v := range vals {
		if n, ok := s.data[v]; ok {
			// v exists
			cur = n.next
		} else {
			// v doesn't exist
			node := &intNode{val: v}
			if s.head == nil {
				// intset is empty
				s.head = node
				s.tail = node
				s.data[v] = node
			} else {
				if cur == nil {
					// append node
					s.tail.next = node
					node.pre = s.tail
					s.tail = node

					s.data[v] = node
					continue
				}

				if cur.val < v {
					if cur.next == nil {
						// cur is the tail, append node
						cur.next = node
						node.pre = cur
						s.tail = node
						s.data[v] = node
					}
					cur = cur.next
				} else {
					// insert node before cur
					if cur.pre == nil {
						// node is the head
						s.head = node
						node.next = cur
						cur.pre = node
					} else {
						node.pre = cur.pre
						cur.pre.next = node
						node.next = cur
						cur.pre = node
					}
					s.data[v] = node
				}
			}
		}
	}
}

// Delete delete elements
func (s *IntSet) Delete(vals ...int) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, v := range vals {
		if node, ok := s.data[v]; ok {
			if node.pre == nil {
				// node is the head
				s.head = node.next
			} else {
				node.pre.next = node.next
			}
			if node.next == nil {
				// node is the tail
				s.tail = node.pre
			} else {
				node.next.pre = node.pre
			}
			delete(s.data, v)
		}
	}
}

// Clear clears all elements
func (s *IntSet) Clear() {
	s.m.Lock()
	defer s.m.Unlock()
	s.data = make(map[int]*intNode)
	s.head = nil
	s.tail = nil
}

// Copy returns a deep copy of itself
func (s *IntSet) Copy() *IntSet {
	s.m.RLock()
	defer s.m.RUnlock()
	var r = NewIntSet()
	cur := s.head
	for cur != nil {
		node := &intNode{val: cur.val}
		if r.head == nil {
			r.head = node
			r.tail = node
		} else {
			node.pre = r.tail
			r.tail.next = node
			r.tail = node
		}
		r.data[node.val] = node
		cur = cur.next
	}
	return r
}

// Length returns IntSet length
func (s *IntSet) Length() int {
	return len(s.data)
}

// Has returns whether v exists in IntSet
func (s *IntSet) Has(v int) bool {
	_, ok := s.data[v]
	return ok
}

// ToList returns elements slice, and SortOption is optional.
// if opts==Asc, returns an ascending slice
// if opts==Desc, returns a descending slice
// if opts==Random returns random order slice (the order in which elements are retrived from map)
// default is Random
func (s *IntSet) ToList(opts ...SortOption) []int {
	s.m.RLock()
	defer s.m.RUnlock()

	var r = make([]int, s.Length())
	var opt = Random
	if len(opts) > 0 {
		opt = opts[0]
	}
	idx := 0
	switch opt {
	case Random:
		for v := range s.data {
			r[idx] = v
			idx++
		}
	case Asc:
		cur := s.head
		for cur != nil {
			r[idx] = cur.val
			cur = cur.next
			idx++
		}
	case Desc:
		cur := s.tail
		for cur != nil {
			r[idx] = cur.val
			cur = cur.pre
			idx++
		}
	}
	return r
}

// Equals returns whether IntSet s has the same members with IntSet t
func (s *IntSet) Equals(t *IntSet) bool {
	if t == nil || s.Length() != t.Length() {
		return false
	}
	if s == t {
		return true
	}
	node1 := s.head
	node2 := t.head
	for node1 != nil && node2 != nil {
		if node1.val != node2.val {
			return false
		}
		node1 = node1.next
		node2 = node2.next
	}
	return true
}

// IsSub returns whether Inset s is a sub set of IntSet t
func (s *IntSet) IsSub(t *IntSet) bool {
	if t == nil || s.Length() > t.Length() {
		return false
	}
	if s == t {
		return true
	}

	s.m.RLock()
	t.m.RLock()
	defer s.m.RUnlock()
	defer t.m.RUnlock()

	cur := s.head
	for cur != nil {
		if !t.Has(cur.val) {
			return false
		}
		cur = cur.next
	}
	return true
}

// Union unions with IntSet t and returns a new IntSet
//
// for example:
// var a=NewIntSet(1,2,3)
// var b=NewIntSet(2,3,4)
// a.Union(b) returns {1,2,3,4}
func (s *IntSet) Union(t *IntSet) *IntSet {
	if t == nil || t.Length() == 0 {
		return s.Copy()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	var r = NewIntSet()
	node1 := s.head
	node2 := t.head
	for node1 != nil || node2 != nil {
		var node = &intNode{}
		if node1 == nil {
			// append node2 val
			node.val = node2.val
			// iterate node2
			node2 = node2.next
		} else if node2 == nil {
			// append node1 val
			node.val = node1.val
			// iterate node1
			node1 = node1.next
		} else {
			if node1.val < node2.val {
				node.val = node1.val
				node1 = node1.next
			} else if node1.val > node2.val {
				node.val = node2.val
				node2 = node2.next
			} else {
				node.val = node1.val
				node1 = node1.next
				node2 = node2.next
			}
		}
		// append node
		if r.head == nil {
			r.head = node
			r.tail = node
		} else {
			node.pre = r.tail
			r.tail.next = node
			r.tail = node
		}
		r.data[node.val] = node
	}
	return r
}

// Intersect returns a new IntSet Whose elements exist in both IntSet
//
// for example:
// var a=NewIntSet(1,2,3)
// var b=NewIntSet(2,3,4)
// a.Intersect(b) returns {2,3}
func (s *IntSet) Intersect(t *IntSet) *IntSet {
	var r = NewIntSet()
	if t == nil {
		return r
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	node1 := s.head
	node2 := t.head
	for node1 != nil && node2 != nil {
		if node1.val < node2.val {
			node1 = node1.next
		} else if node1.val > node2.val {
			node2 = node2.next
		} else {
			node := &intNode{val: node1.val}
			if r.head == nil {
				r.head = node
				r.tail = node
			} else {
				node.pre = r.tail
				r.tail.next = node
				r.tail = node
			}
			r.data[node.val] = node
			node1 = node1.next
			node2 = node2.next
		}
	}
	return r
}

// Subtract returns a new IntSet Whose elements exist in itself but don't exist in IntSet t
//
// for example:
// var a=NewIntSet(1,2,3)
// var b=NewIntSet(2,3,4)
// a.Subtract(b) returns {1}
func (s *IntSet) Subtract(t *IntSet) *IntSet {
	if t == nil || t.Length() == 0 {
		return s.Copy()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	var r = NewIntSet()
	cur := s.head
	for cur != nil {
		if !t.Has(cur.val) {
			node := &intNode{val: cur.val}
			if r.head == nil {
				r.head = node
				r.tail = node
			} else {
				node.pre = r.tail
				r.tail.next = node
				r.tail = node
			}
			r.data[node.val] = node
		}
		cur = cur.next
	}
	return r
}

// Complement returns a new IntSet Whose elements only exists in one IntSet
//
// for example:
// var a=NewIntSet(1,2,3)
// var b=NewIntSet(2,3,4)
// a.Complement(b) returns {1,4}
func (s *IntSet) Complement(t *IntSet) *IntSet {
	if t == nil || t.Length() == 0 {
		return s.Copy()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	var r = NewIntSet()
	node1 := s.head
	node2 := t.head
	for node1 != nil || node2 != nil {
		if node1 != nil && node2 != nil && node1.val == node2.val {
			node1 = node1.next
			node2 = node2.next
			continue
		}

		var node = &intNode{}
		if node1 == nil {
			node.val = node2.val
			node2 = node2.next
		} else if node2 == nil {
			node.val = node1.val
			node1 = node1.next
		} else {
			if node1.val < node2.val {
				node.val = node1.val
				node1 = node1.next
			} else {
				node.val = node2.val
				node2 = node2.next
			}
		}

		if r.head == nil {
			r.head = node
			r.tail = node
		} else {
			node.pre = r.tail
			r.tail.next = node
			r.tail = node
		}
		r.data[node.val] = node
	}
	return r
}
