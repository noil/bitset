package bitset

import (
	"testing"
)

const SIZE = 5

func TestBitsetContains(t *testing.T) {
	set := New(SIZE)
	for i := uint(0); i < SIZE; i++ {
		set.Add(i)
	}
	for n := uint(0); n < SIZE; n++ {
		if !set.Contains(n) {
			t.Errorf("Value %d not found.", n)
		}
	}
}

func TestBitsetRemove(t *testing.T) {
	set := New(SIZE)
	for i := uint(0); i < SIZE; i++ {
		set.Add(i)
	}
	for n := uint(0); n < SIZE; n++ {
		set.Remove(n)
	}
	for n := uint(0); n < SIZE; n++ {
		if set.Contains(n) {
			t.Errorf("Value %d wasn't removed.", n)
		}
	}
}

func TestBitsetEnumerate(t *testing.T) {
	set := New(10)
	for i := uint(0); i < 10; i++ {
		if i%2 == 0 {
			set.Add(i)
		}
	}
	sum := uint(0)
	for _, v := range set.Enumerate() {
		sum += v
	}
	if 20 != sum {
		t.Errorf("Wants [0 2 4 6 8], gets %v", set.Enumerate())
	}
}

func TestBitsetUnion(t *testing.T) {
	a := New(15)
	b := New(15)
	c := New(15)
	for i := uint(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint(25); i < 35; i++ {
		c.Add(i)
	}
	a.Union(b, c)
	sum := uint(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 400 != sum {
		t.Errorf("Wants [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 25 26 27 28 29 30 31 32 33 34], gets %v", a.Enumerate())
	}
}

func TestBitsetIntersection(t *testing.T) {
	a := New(15)
	b := New(15)
	c := New(15)
	for i := uint(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint(9); i < 19; i++ {
		c.Add(i)
	}
	a.Intersection(c, b)
	sum := uint(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 9 != sum {
		t.Errorf("Wants [9], gets %v", a.Enumerate())
	}
}

func TestBitsetDifference(t *testing.T) {
	a := New(15)
	b := New(15)
	c := New(15)
	for i := uint(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint(8); i < 18; i++ {
		c.Add(i)
	}
	a.Difference(b, c)
	sum := uint(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 58 != sum {
		t.Errorf("Wants [0 1 2 3 4 15 16 17], gets %v", a.Enumerate())
	}
}

func BenchmarkBitsetAdd(b *testing.B) {
	size := uint(b.N)
	set := New(size)
	for n := uint(0); n < uint(b.N); n++ {
		set.Add(n)
	}
}

func BenchmarkSliceAdd(b *testing.B) {
	size := uint(b.N)
	sl := make([]uint, size)
	for n := uint(0); n < uint(b.N); n++ {
		sl = append(sl, n)
	}
}

func BenchmarkMapAdd(b *testing.B) {
	size := uint(b.N)
	mp := make(map[uint]bool, size)
	for n := uint(0); n < uint(b.N); n++ {
		mp[n] = true
	}
}
