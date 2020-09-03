package bitset

import (
	"testing"
)

func TestBitsetThreadSave64Contains(t *testing.T) {
	set := NewThreadSave64(SIZE)
	for i := uint64(0); i < SIZE; i++ {
		set.Add(i)
	}
	for n := uint64(0); n < SIZE; n++ {
		if !set.Contains(n) {
			t.Errorf("Value %d not found.", n)
		}
	}
}

func TestBitsetThreadSave64Remove(t *testing.T) {
	set := NewThreadSave64(SIZE)
	for i := uint64(0); i < SIZE; i++ {
		set.Add(i)
	}
	for n := uint64(0); n < SIZE; n++ {
		set.Remove(n)
	}
	for n := uint64(0); n < SIZE; n++ {
		if set.Contains(n) {
			t.Errorf("Value %d wasn't removed.", n)
		}
	}
}

func TestBitsetThreadSave64Enumerate(t *testing.T) {
	set := NewThreadSave64(10)
	for i := uint64(0); i < 10; i++ {
		if i%2 == 0 {
			set.Add(i)
		}
	}
	sum := uint64(0)
	for _, v := range set.Enumerate() {
		sum += v
	}
	if 20 != sum {
		t.Errorf("Wants [0 2 4 6 8], gets %v", set.Enumerate())
	}
}

func TestBitThreadSaveset64Union(t *testing.T) {
	a := NewThreadSave64(15)
	b := NewThreadSave64(15)
	c := NewThreadSave64(15)
	for i := uint64(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint64(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint64(25); i < 35; i++ {
		c.Add(i)
	}
	a.Union(b, c)
	sum := uint64(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 400 != sum {
		t.Errorf("Wants [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 25 26 27 28 29 30 31 64 33 34], gets %v", a.Enumerate())
	}
}

func TestBitsetThreadSave64Intersection(t *testing.T) {
	a := NewThreadSave64(15)
	b := NewThreadSave64(15)
	c := NewThreadSave64(15)
	for i := uint64(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint64(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint64(9); i < 19; i++ {
		c.Add(i)
	}
	a.Intersection(c, b)
	sum := uint64(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 9 != sum {
		t.Errorf("Wants [9], gets %v", a.Enumerate())
	}
}

func TestBitsetThreadSave64Difference(t *testing.T) {
	a := NewThreadSave64(15)
	b := NewThreadSave64(15)
	c := NewThreadSave64(15)
	for i := uint64(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint64(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint64(8); i < 18; i++ {
		c.Add(i)
	}
	a.Difference(b, c)
	sum := uint64(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 58 != sum {
		t.Errorf("Wants [0 1 2 3 4 15 16 17], gets %v", a.Enumerate())
	}
}
