## Introduction

Library `bitset` implements a set data structure for positive numbers. By default set is thread-unsafe. If you need to use a thread save approach, you should use ThreadSaveSet. Keep in mind that the thread save approach works slower.

## Methods:

- Add
- Remove
- Contains
- IsEmpty
- Enumerate
- Union
- Intersection
- Difference
- Size

### Installation

To install this package, you need to setup your Go workspace. The simplest way to install the library is to run:
```powershell
$ go get github.com/noil/bitset
```

## Quick code sample

```go
package main

import (
	"fmt"
	"github.com/noil/bitset"
)

func main() {
    // init size of set, it increse automaticly
    size := uint(10000)
    st := bitset.NewWithSize(size)

    // add single value
    st.Add(uint(100))
    // add sevral values
    st.Add(uint(50), uint(150))
    // add slice
    sl := []uint{1,2,3,4,5}
    st.Add(sl...)
    // check for existing
    if st.Contains(uint(100)) {
        fmt.Println("ok!")
    }
    // remove single value
    st.Remove(uint(150))
    // remove sevral values
    st.Remove(uint(1), uint(2))
    // remove slice
    st.Remove(sl...)
    // get all values in slice
    for _, v := range st.Enumerate() {
        fmt.Println(v)
    }
    // union single set
    st1 := bitset.NewWithSize(100)
    st1.Add(uint(1), uint(2), uint(3))
    st.Unions(st1)
    // union sevral sets
    st2 := bitset.NewWithSize(100)
    st2.Add(uint(4), uint(5), uint(6))
    st3 := bitset.NewWithSize(100)
    sl1 := []uint{1,2,3,4,5,6}
    st3.Add(sl1...)
    st.Unions(st2, st3)
}
```

## Benchmark

```bash
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/noil/bitset
BenchmarkBitsetAdd-4               	265514660	         4.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitsetAddInt-4            	259977708	         4.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAdd-4                	100000000	        43.1 ns/op	      65 B/op	       0 allocs/op
BenchmarkMapAdd-4                  	121503525	        10.0 ns/op	      25 B/op	       0 allocs/op
BenchmarkBitsetThreadSaveAdd-4     	12322563	        98.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitsetTreadSaveAddInt-4   	12684298	        93.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceTreadSaveAdd-4       	21254306	        89.2 ns/op	     137 B/op	       0 allocs/op
BenchmarkMapTreadSaveAdd-4         	17331967	        85.6 ns/op	      22 B/op	       0 allocs/op
PASS
ok  	github.com/noil/bitset	16.617s
```

License
----

MIT