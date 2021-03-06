package bitset

import (
	"sync"
)

// ThreadSaveSet a set of unsigned interger which store unique values, without any particular order
type ThreadSaveSet struct {
	set *Set
	mu  *sync.RWMutex
}

// NewThreadSave creates a new ThreadSaveSet, initially empty set structure
func NewThreadSave() *ThreadSaveSet {
	return NewThreadSaveWithSize(uintSize)
}

// NewThreadSaveWithSize creates a new ThreadSaveSet, initially empty set structure
func NewThreadSaveWithSize(size uint) *ThreadSaveSet {
	return &ThreadSaveSet{
		set: &Set{
			m: make([]uint, size/uintSize),
		},
		mu: &sync.RWMutex{},
	}
}

// Add adds elements to ThreadSaveSet, if it is not present already
func (s *ThreadSaveSet) Add(x ...uint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Add(x...)
}

// AddInt adds elements to ThreadSaveSet, if it is not present already
func (s *ThreadSaveSet) AddInt(x ...int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.AddInt(x...)
}

// AddInt64 adds elements to ThreadSaveSet, if it is not present already
func (s *ThreadSaveSet) AddInt64(x ...int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.AddInt64(x...)
}

// Remove removes elements from ThreadSaveSet, if it is present
func (s *ThreadSaveSet) Remove(x ...uint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Remove(x...)
}

// RemoveInt removes elements from ThreadSaveSet, if it is present
func (s *ThreadSaveSet) RemoveInt(x ...int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.RemoveInt(x...)
}

// RemoveInt64 removes elements from ThreadSaveSet, if it is present
func (s *ThreadSaveSet) RemoveInt64(x ...int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.RemoveInt64(x...)
}

// Contains checks whether the value x is in the set ThreadSaveSet
func (s ThreadSaveSet) Contains(x uint) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Contains(x)
}

// ContainsInt checks whether the value x is in the set ThreadSaveSet
func (s ThreadSaveSet) ContainsInt(x int) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.ContainsInt(x)
}

// ContainsInt64 checks whether the value x is in the set ThreadSaveSet
func (s ThreadSaveSet) ContainsInt64(x int64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.ContainsInt64(x)
}

// Enumerate returns an array of all values from ThreadSaveSet
func (s *ThreadSaveSet) Enumerate() []uint {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Enumerate()
}

// Union makes union of ThreadSaveSet s with one or more ThreadSaveSet ss
func (s *ThreadSaveSet) Union(ss ...*ThreadSaveSet) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, n := range ss {
		n.mu.RLock()
		s.set.Add(n.Enumerate()...)
		n.mu.RUnlock()
	}
}

// Intersection makes the intersection of ThreadSaveSet s  with one or more ThreadSaveSet ss
func (s *ThreadSaveSet) Intersection(ss ...*ThreadSaveSet) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, n := range ss {
		for i := uint(0); i < s.set.Size(); i++ {
			n.mu.RLock()
			s.set.m[i] &= n.set.m[i]
			n.mu.RUnlock()
		}
	}
}

// Difference makes the difference of ThreadSaveSet s with one or more ThreadSaveSet ss
func (s *ThreadSaveSet) Difference(ss ...*ThreadSaveSet) {
	s.mu.Lock()
	defer s.mu.Unlock()
	tmp := NewThreadSaveWithSize(s.set.Size())
	tmpM := make([]uint, s.set.Size())
	copy(tmpM, s.set.m)
	tmp.set.m = tmpM
	for _, n := range ss {
		n.mu.RLock()
		s.set.Add(n.set.Enumerate()...)
		n.mu.RUnlock()
	}
	ss = append(ss, tmp)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			tmp := NewThreadSaveWithSize(ss[i].Size())
			tmpM := make([]uint, s.set.Size())
			ss[i].mu.RLock()
			copy(tmpM, ss[i].set.m)
			ss[i].mu.RUnlock()
			tmp.set.m = tmpM
			tmp.Intersection(ss[j])
			s.set.Remove(tmp.Enumerate()...)
		}
	}
}

// IsEmpty checks whether the set ThreadSaveSet is empty
func (s ThreadSaveSet) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IsEmpty()
}

// Size returns the number of elements in ThreadSaveSet
func (s ThreadSaveSet) Size() uint {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Size()
}
