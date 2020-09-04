package bitset

import (
	"sync"
	"testing"
)

func TestBitsetThreadSaveContains(t *testing.T) {
	set := NewThreadSaveWithSize(SIZE)
	for i := uint(0); i < SIZE; i++ {
		set.Add(i)
	}
	for n := uint(0); n < SIZE; n++ {
		if !set.Contains(n) {
			t.Errorf("Value %d not found.", n)
		}
	}
}

func TestBitsetThreadSaveRemove(t *testing.T) {
	set := NewThreadSaveWithSize(SIZE)
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

func TestBitsetThreadSaveEnumerate(t *testing.T) {
	set := NewThreadSaveWithSize(10)
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

func TestBitThreadSavesetUnion(t *testing.T) {
	a := NewThreadSaveWithSize(15)
	b := NewThreadSaveWithSize(15)
	c := NewThreadSaveWithSize(15)
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

func TestBitsetThreadSaveIntersection(t *testing.T) {
	a := NewThreadSaveWithSize(15)
	b := NewThreadSaveWithSize(15)
	c := NewThreadSaveWithSize(15)
	for i := uint(0); i < 10; i++ {
		a.Add(i)
	}
	for i := uint(5); i < 15; i++ {
		b.Add(i)
	}
	for i := uint(9); i < 19; i++ {
		c.Add(i)
	}
	a.Intersection(b, c)
	sum := uint(0)
	for _, v := range a.Enumerate() {
		sum += v
	}
	if 9 != sum {
		t.Errorf("Wants [9], gets %v", a.Enumerate())
	}
}

func TestBitsetThreadSaveDifference(t *testing.T) {
	a := NewThreadSaveWithSize(15)
	b := NewThreadSaveWithSize(15)
	c := NewThreadSaveWithSize(15)
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

func BenchmarkBitsetThreadSaveAdd(b *testing.B) {
	size := uint(b.N)
	set := NewThreadSaveWithSize(size)
	go func() {
		j := uint(0)
		for n := 0; n < b.N; n++ {
			set.Add(j)
			j++
		}
	}()
	j := uint(0)
	for n := 0; n < b.N; n++ {
		set.Add(j)
		j++
	}
}
func BenchmarkBitsetTreadSaveAddInt(b *testing.B) {
	size := uint(b.N)
	set := NewThreadSaveWithSize(size)
	go func() {
		j := 0
		for n := 0; n < b.N; n++ {
			set.AddInt(j)
			j++
		}
	}()
	j := 0
	for n := 0; n < b.N; n++ {
		set.AddInt(j)
		j++
	}
}

func BenchmarkSliceTreadSaveAdd(b *testing.B) {
	size := uint(b.N)
	sl := make([]uint, size)
	mu := sync.Mutex{}
	go func() {
		j := uint(0)
		for i := 0; i < b.N; i++ {
			mu.Lock()
			sl = append(sl, j)
			mu.Unlock()
			j++
		}
	}()
	j := uint(0)
	for i := 0; i < b.N; i++ {
		mu.Lock()
		sl = append(sl, j)
		mu.Unlock()
		j++
	}
}

func BenchmarkMapTreadSaveAdd(b *testing.B) {
	size := uint(b.N)
	mp := make(map[uint]bool, size)
	mu := sync.Mutex{}
	go func() {
		j := uint(0)
		for i := 0; i < b.N; i++ {
			mu.Lock()
			mp[j] = true
			mu.Unlock()
		}
	}()
	j := uint(0)
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mp[j] = true
		mu.Unlock()
	}
}
