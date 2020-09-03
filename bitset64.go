package bitset

// Set64 a set of unsigned interger 64 which store unique values, without any particular order
type Set64 struct {
	m []uint64
}

// New64 creates a new Set64, initially empty set structure
func New64(size uint64) *Set64 {
	return &Set64{
		m: make([]uint64, size/64),
	}
}

// Add adds elements to Set64, if it is not present already
func (s *Set64) Add(x ...uint64) {
	for _, e := range x {
		i, j := getInteger64AndRemainder(e)
		if uint64(len(s.m)) < i+1 {
			size := uint64(2)
			if i > 0 {
				size = i * 2
			}
			tmpM := make([]uint64, size)
			copy(tmpM, s.m)
			s.m = tmpM
		}
		s.m[i] |= 1 << j
	}
}

// Remove removes elements from Set64, if it is present
func (s *Set64) Remove(x ...uint64) {
	for _, e := range x {
		i, j := getInteger64AndRemainder(e)
		if uint64(len(s.m)) < i+1 {
			return
		}
		s.m[i] &= ^(1 << j)
	}
}

// Contains checks whether the value x is in the set Set64
func (s Set64) Contains(x uint64) bool {
	i, j := getInteger64AndRemainder(x)
	if uint64(len(s.m)) < i+1 {
		return false
	}
	if 1 == s.m[i]>>j&1 {
		return true
	}

	return false
}

// Enumerate returns an array of all values
func (s *Set64) Enumerate() []uint64 {
	result := []uint64{}
	for factor, value := range s.m {
		for i := 0; i < uintSize; i++ {
			if 1 == value>>i&1 {
				result = append(result, uint64((64*factor)+i))
			}
		}
	}

	return result
}

// Union makes union of Set64 s with one or more Set64 ss
func (s *Set64) Union(ss ...*Set64) {
	for _, n := range ss {
		s.Add(n.Enumerate()...)
	}
}

// Intersection makes the intersection of Set64 s  with one or more Set64 ss
func (s *Set64) Intersection(ss ...*Set64) {
	for _, n := range ss {
		for i := uint64(0); i < s.Size(); i++ {
			s.m[i] &= n.m[i]
		}
	}
}

// Difference makes the difference of Set64 s with one or more Set64 ss
func (s *Set64) Difference(ss ...*Set64) {
	tmp := New64(s.Size())
	tmpM := make([]uint64, s.Size())
	copy(tmpM, s.m)
	tmp.m = tmpM
	s.Union(ss...)
	ss = append(ss, tmp)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			tmp := New64(ss[i].Size())
			tmpM := make([]uint64, s.Size())
			copy(tmpM, ss[i].m)
			tmp.m = tmpM
			tmp.Intersection(ss[j])
			s.Remove(tmp.Enumerate()...)
		}
	}
}

// IsEmpty checks whether the set Set64 is empty
func (s Set64) IsEmpty() bool {
	if len(s.m) == 0 {
		return true
	}

	return false
}

// Size returns the number of elements in Set64
func (s Set64) Size() uint64 {
	return uint64(len(s.m))
}

func getInteger64AndRemainder(value uint64) (x uint64, y uint64) {
	x = value / 64
	y = value % 64
	return
}
