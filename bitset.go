package bitset

const uintSize = 32 << (^uint(0) >> 32 & 1)

// Set a set of unsigned interger which store unique values, without any particular order
type Set struct {
	m []uint
}

// New creates a new Set, initially empty set structure
func New(size uint) *Set {
	return &Set{
		m: make([]uint, size/uintSize),
	}
}

// Add adds elements to Set, if it is not present already
func (s *Set) Add(x ...uint) {
	for _, e := range x {
		i, j := getIntegerAndRemainder(e)
		if uint(len(s.m)) < i+1 {
			size := uint(2)
			if i > 0 {
				size = i * 2
			}
			tmpM := make([]uint, size)
			copy(tmpM, s.m)
			s.m = tmpM
		}
		s.m[i] |= 1 << j
	}
}

// Remove removes elements from Set, if it is present
func (s *Set) Remove(x ...uint) {
	for _, e := range x {
		i, j := getIntegerAndRemainder(e)
		if uint(len(s.m)) < i+1 {
			return
		}
		s.m[i] &= ^(1 << j)
	}
}

// Contains checks whether the value x is in the set Set
func (s Set) Contains(x uint) bool {
	i, j := getIntegerAndRemainder(x)
	if uint(len(s.m)) < i+1 {
		return false
	}
	if 1 == s.m[i]>>j&1 {
		return true
	}

	return false
}

// IsEmpty checks whether the set Set is empty
func (s Set) IsEmpty() bool {
	if len(s.m) == 0 {
		return true
	}

	return false
}

// Enumerate returns an array of all values
func (s *Set) Enumerate() []uint {
	result := []uint{}
	for factor, value := range s.m {
		for i := 0; i < uintSize; i++ {
			if 1 == value>>i&1 {
				result = append(result, uint((uintSize*factor)+i))
			}
		}
	}

	return result
}

// Union makes union of set s with one or more set ss
func (s *Set) Union(ss ...*Set) {
	for _, n := range ss {
		s.Add(n.Enumerate()...)
	}
}

// Intersection makes the intersection of set s  with one or more set ss
func (s *Set) Intersection(ss ...*Set) {
	for _, n := range ss {
		for i := uint(0); i < s.Size(); i++ {
			s.m[i] &= n.m[i]
		}
	}
}

// Difference makes the difference of set s with one or more set ss
func (s *Set) Difference(ss ...*Set) {
	tmp := New(s.Size())
	tmpM := make([]uint, s.Size())
	copy(tmpM, s.m)
	tmp.m = tmpM
	s.Union(ss...)
	ss = append(ss, tmp)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			tmp := New(ss[i].Size())
			tmpM := make([]uint, s.Size())
			copy(tmpM, ss[i].m)
			tmp.m = tmpM
			tmp.Intersection(ss[j])
			s.Remove(tmp.Enumerate()...)
		}
	}
}

// Size returns the number of elements in Set
func (s Set) Size() uint {
	return uint(len(s.m))
}

func getIntegerAndRemainder(value uint) (x uint, y uint) {
	x = value / uintSize
	y = value % uintSize
	return
}
