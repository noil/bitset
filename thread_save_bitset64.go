package bitset

import (
	"sync"
)

// ThreadSaveSet64 a thred save set of unsigned interger 64 which store unique values, without any particular order
type ThreadSaveSet64 struct {
	set *Set64
	mu  *sync.RWMutex
}

// NewThreadSave64 creates a new ThreadSaveSet64, initially empty set structure
func NewThreadSave64(size uint64) *ThreadSaveSet64 {
	return &ThreadSaveSet64{
		set: &Set64{
			m: make([]uint64, size/64),
		},
		mu: &sync.RWMutex{},
	}
}

// Add adds elements to ThreadSaveSet64, if it is not present already
func (s *ThreadSaveSet64) Add(x ...uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Add(x...)
}

// Remove removes elements from ThreadSaveSet64, if it is present
func (s *ThreadSaveSet64) Remove(x ...uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Remove(x...)
}

// Contains checks whether the value x is in the set ThreadSaveSet64
func (s ThreadSaveSet64) Contains(x uint64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Contains(x)
}

// Enumerate returns an array of all values from ThreadSaveSet64
func (s *ThreadSaveSet64) Enumerate() []uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Enumerate()
}

// Union makes union of ThreadSaveSet64 s with one or more ThreadSaveSet64 ss
func (s *ThreadSaveSet64) Union(ss ...*ThreadSaveSet64) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, n := range ss {
		s.set.Add(n.Enumerate()...)
	}
}

// Intersection makes the intersection of ThreadSaveSet64 s  with one or more ThreadSaveSet64 ss
func (s *ThreadSaveSet64) Intersection(ss ...*ThreadSaveSet64) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, n := range ss {
		for i := uint64(0); i < s.Size(); i++ {
			s.set.m[i] &= n.set.m[i]
		}
	}
}

// Difference makes the difference of ThreadSaveSet64 s with one or more ThreadSaveSet64 ss
func (s *ThreadSaveSet64) Difference(ss ...*ThreadSaveSet64) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tmp := NewThreadSave64(s.Size())
	tmpM := make([]uint64, s.Size())
	copy(tmpM, s.set.m)
	tmp.set.m = tmpM
	s.Union(ss...)
	ss = append(ss, tmp)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			tmp := NewThreadSave64(ss[i].Size())
			tmpM := make([]uint64, s.Size())
			copy(tmpM, ss[i].set.m)
			tmp.set.m = tmpM
			tmp.Intersection(ss[j])
			s.Remove(tmp.Enumerate()...)
		}
	}
}

// IsEmpty checks whether the set ThreadSaveSet64 is empty
func (s ThreadSaveSet64) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IsEmpty()
}

// Size returns the number of elements in ThreadSaveSet64
func (s ThreadSaveSet64) Size() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Size()
}
