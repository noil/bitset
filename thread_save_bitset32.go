package bitset

import (
	"sync"
)

// ThreadSaveSet32 a thred save set of unsigned interger 32 which store unique values, without any particular order
type ThreadSaveSet32 struct {
	set *Set32
	mu  *sync.RWMutex
}

// NewThreadSave32 creates a new ThreadSaveSet32, initially empty set structure
func NewThreadSave32(size uint32) *ThreadSaveSet32 {
	return &ThreadSaveSet32{
		set: &Set32{
			m: make([]uint32, size/32),
		},
		mu: &sync.RWMutex{},
	}
}

// Add adds elements to ThreadSaveSet32, if it is not present already
func (s *ThreadSaveSet32) Add(x ...uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Add(x...)
}

// Remove removes elements from ThreadSaveSet32, if it is present
func (s *ThreadSaveSet32) Remove(x ...uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Remove(x...)
}

// Contains checks whether the value x is in the set ThreadSaveSet32
func (s ThreadSaveSet32) Contains(x uint32) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Contains(x)
}

// Enumerate returns an array of all values from ThreadSaveSet32
func (s *ThreadSaveSet32) Enumerate() []uint32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Enumerate()
}

// Union makes union of ThreadSaveSet32 s with one or more ThreadSaveSet32 ss
func (s *ThreadSaveSet32) Union(ss ...*ThreadSaveSet32) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, n := range ss {
		s.set.Add(n.Enumerate()...)
	}
}

// Intersection makes the intersection of ThreadSaveSet32 s  with one or more ThreadSaveSet32 ss
func (s *ThreadSaveSet32) Intersection(ss ...*ThreadSaveSet32) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, n := range ss {
		for i := uint32(0); i < s.Size(); i++ {
			s.set.m[i] &= n.set.m[i]
		}
	}
}

// Difference makes the difference of ThreadSaveSet32 s with one or more ThreadSaveSet32 ss
func (s *ThreadSaveSet32) Difference(ss ...*ThreadSaveSet32) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tmp := NewThreadSave32(s.Size())
	tmpM := make([]uint32, s.Size())
	copy(tmpM, s.set.m)
	tmp.set.m = tmpM
	s.Union(ss...)
	ss = append(ss, tmp)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			tmp := NewThreadSave32(ss[i].Size())
			tmpM := make([]uint32, s.Size())
			copy(tmpM, ss[i].set.m)
			tmp.set.m = tmpM
			tmp.Intersection(ss[j])
			s.Remove(tmp.Enumerate()...)
		}
	}
}

// IsEmpty checks whether the set ThreadSaveSet32 is empty
func (s ThreadSaveSet32) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IsEmpty()
}

// Size returns the number of elements in ThreadSaveSet32
func (s ThreadSaveSet32) Size() uint32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Size()
}
