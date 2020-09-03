package bitset

import "sync"

// ThreadSaveSet a set of unsigned interger which store unique values, without any particular order
type ThreadSaveSet struct {
	set *Set
	mu  *sync.RWMutex
}

// NewThreadSave creates a new ThreadSaveSet, initially empty set structure
func NewThreadSave(size uint) *ThreadSaveSet {
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

// Remove removes elements from ThreadSaveSet, if it is present
func (s *ThreadSaveSet) Remove(x ...uint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Remove(x...)
}

// Contains checks whether the value x is in the set ThreadSaveSet
func (s ThreadSaveSet) Contains(x uint) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Contains(x)
}

// Enumerate returns an array of all values from ThreadSaveSet
func (s *ThreadSaveSet) Enumerate() []uint {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Enumerate()
}

// Union makes union of ThreadSaveSet s with one or more ThreadSaveSet ss
func (s *ThreadSaveSet) Union(ss ...*ThreadSaveSet) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, n := range ss {
		s.set.Add(n.Enumerate()...)
	}
}

// Intersection makes the intersection of ThreadSaveSet s  with one or more ThreadSaveSet ss
func (s *ThreadSaveSet) Intersection(ss ...*ThreadSaveSet) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, n := range ss {
		for i := uint(0); i < s.Size(); i++ {
			s.set.m[i] &= n.set.m[i]
		}
	}
}

// Difference makes the difference of ThreadSaveSet s with one or more ThreadSaveSet ss
func (s *ThreadSaveSet) Difference(ss ...*ThreadSaveSet) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tmp := NewThreadSave(s.Size())
	tmpM := make([]uint, s.Size())
	copy(tmpM, s.set.m)
	tmp.set.m = tmpM
	s.Union(ss...)
	ss = append(ss, tmp)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			tmp := NewThreadSave(ss[i].Size())
			tmpM := make([]uint, s.Size())
			copy(tmpM, ss[i].set.m)
			tmp.set.m = tmpM
			tmp.Intersection(ss[j])
			s.Remove(tmp.Enumerate()...)
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
