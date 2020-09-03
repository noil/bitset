package bitset

import (
	"testing"
)

func TestBitset32Contains(t *testing.T) {
	set := New32(SIZE)
	for i := uint32(0); i < SIZE; i++ {
		set.Add(i)
	}
	for n := uint32(0); n < SIZE; n++ {
		if !set.Contains(n) {
			t.Errorf("Value %d not found.", n)
		}
	}
}

func TestBitset32Remove(t *testing.T) {
	set := New32(SIZE)
	for i := uint32(0); i < SIZE; i++ {
		set.Add(i)
	}
	for n := uint32(0); n < SIZE; n++ {
		set.Remove(n)
	}
	for n := uint32(0); n < SIZE; n++ {
		if set.Contains(n) {
			t.Errorf("Value %d wasn't removed.", n)
		}
	}
}

func TestBitset32Enumerate(t *testing.T) {
	set := New32(10)
	for i := uint32(0); i < 10; i++ {
		if i%2 == 0 {
			set.Add(i)
		}
	}
	sum := uint32(0)
	for _, v := range set.Enumerate() {
		sum += v
	}
	if 20 != sum {
		t.Errorf("Wants [0 2 4 6 8], gets %v", set.Enumerate())
	}
}

func TestBit32setUnion(t *testing.T) {
	a := New32(15)
	b := New32(15)
	c := New32(15)
	for i := uint32(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint32(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint32(25); i < 35; i++ {
		c.Add(i)
	}
	a.Union(b, c)
	sum := uint32(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 400 != sum {
		t.Errorf("Wants [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 25 26 27 28 29 30 31 32 33 34], gets %v", a.Enumerate())
	}
}

func TestBitset32Intersection(t *testing.T) {
	a := New32(15)
	b := New32(15)
	c := New32(15)
	for i := uint32(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint32(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint32(9); i < 19; i++ {
		c.Add(i)
	}
	a.Intersection(c, b)
	sum := uint32(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 9 != sum {
		t.Errorf("Wants [9], gets %v", a.Enumerate())
	}
}

func TestBitset32Difference(t *testing.T) {
	a := New32(15)
	b := New32(15)
	c := New32(15)
	for i := uint32(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint32(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint32(8); i < 18; i++ {
		c.Add(i)
	}
	a.Difference(b, c)
	sum := uint32(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 58 != sum {
		t.Errorf("Wants [0 1 2 3 4 15 16 17], gets %v", a.Enumerate())
	}
}
