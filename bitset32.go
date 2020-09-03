package bitset

// Set32 a set of unsigned interger 32 which store unique values, without any particular order
type Set32 struct {
	m []uint32
}

// New32 creates a new Set32, initially empty set structure
func New32(size uint32) *Set32 {
	return &Set32{
		m: make([]uint32, size/32),
	}
}

// Add adds elements to Set32, if it is not present already
func (s *Set32) Add(x ...uint32) {
	for _, e := range x {
		i, j := getInteger32AndRemainder(e)
		if uint32(len(s.m)) < i+1 {
			size := uint32(2)
			if i > 0 {
				size = i * 2
			}
			tmpM := make([]uint32, size)
			copy(tmpM, s.m)
			s.m = tmpM
		}
		s.m[i] |= 1 << j
	}
}

// Remove removes elements from Set32, if it is present
func (s *Set32) Remove(x ...uint32) {
	for _, e := range x {
		i, j := getInteger32AndRemainder(e)
		if uint32(len(s.m)) < i+1 {
			return
		}
		s.m[i] &= ^(1 << j)
	}
}

// Contains checks whether the value x is in the set Set32
func (s Set32) Contains(x uint32) bool {
	i, j := getInteger32AndRemainder(x)
	if uint32(len(s.m)) < i+1 {
		return false
	}
	if 1 == s.m[i]>>j&1 {
		return true
	}

	return false
}

// Enumerate returns an array of all values
func (s *Set32) Enumerate() []uint32 {
	result := []uint32{}
	for factor, value := range s.m {
		for i := 0; i < uintSize; i++ {
			if 1 == value>>i&1 {
				result = append(result, uint32((32*factor)+i))
			}
		}
	}

	return result
}

// Union makes union of Set32 s with one or more Set32 ss
func (s *Set32) Union(ss ...*Set32) {
	for _, n := range ss {
		s.Add(n.Enumerate()...)
	}
}

// Intersection makes the intersection of Set32 s  with one or more Set32 ss
func (s *Set32) Intersection(ss ...*Set32) {
	for _, n := range ss {
		for i := uint32(0); i < s.Size(); i++ {
			s.m[i] &= n.m[i]
		}
	}
}

// Difference makes the difference of Set32 s with one or more Set32 ss
func (s *Set32) Difference(ss ...*Set32) {
	tmp := New32(s.Size())
	tmpM := make([]uint32, s.Size())
	copy(tmpM, s.m)
	tmp.m = tmpM
	s.Union(ss...)
	ss = append(ss, tmp)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			tmp := New32(ss[i].Size())
			tmpM := make([]uint32, s.Size())
			copy(tmpM, ss[i].m)
			tmp.m = tmpM
			tmp.Intersection(ss[j])
			s.Remove(tmp.Enumerate()...)
		}
	}
}

// IsEmpty checks whether the set Set32 is empty
func (s Set32) IsEmpty() bool {
	if len(s.m) == 0 {
		return true
	}

	return false
}

// Size returns the number of elements in Set32
func (s Set32) Size() uint32 {
	return uint32(len(s.m))
}

func getInteger32AndRemainder(value uint32) (x uint32, y uint32) {
	x = value / 32
	y = value % 32
	return
}
