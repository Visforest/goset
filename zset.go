package goset

import (
	"cmp"
)

func NewSortedSet[T cmp.Ordered](vals ...T) *SortedSet[T] {
	s := &SortedSet[T]{newLinearSet[T]()}
	s.Add(vals...)
	return s
}

type SortedSet[T cmp.Ordered] struct {
	*linearSet[T]
}

func (l *SortedSet[T]) Add(v ...T) {
	if len(v) == 0 {
		return
	}
	defer l.m.Unlock()
	l.m.Lock()

	var i int
	if l.head == nil {
		// first node
		n := &setNode[T]{
			val: v[i],
		}
		l.head = n
		l.tail = n
		l.data[v[i]] = n
		i++
	}
	for ; i < len(v); i++ {
		if _, ok := l.data[v[i]]; !ok {
			n := &setNode[T]{
				val: v[i],
			}
			if cmp.Less(v[i], l.head.val) {
				// add to head
				n.next = l.head
				l.head.pre = n
				l.head = n
				l.data[v[i]] = n
			} else if cmp.Less(l.tail.val, v[i]) {
				//	add to tail
				n.pre = l.tail
				l.tail.next = n
				l.tail = n
				l.data[v[i]] = n
			} else {
				// search and insert
				left := l.head
				right := left.next
				for right != nil {
					if cmp.Less(v[i], right.val) {
						// insert and break
						left.next = n
						right.pre = n
						n.pre = left
						n.next = right
						l.data[v[i]] = n
						break
					}
					// go on
					right = right.next
					left = left.next
				}
			}
		}
	}
}
