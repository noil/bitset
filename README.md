## Introduction

Library `bitset` implements a set data structure for positive numbers. There are 3 types of sets in the library. Set depends on the bitness of the platform, Set32 uses 32-bit numbers and Set64 uses 64-bit numbers. All these 3 types are thread-unsafe. If you need to use a thread save approach, you should use ThreadSaveSet, ThreadSaveSet32, and ThreadSaveSet64 accordingly. Keep in mind that the thread save approach works slower.

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
    st := bitset.New(size)

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
    st1 := bitset.New(100)
    st1.Add(uint(1), uint(2), uint(3))
    st.Unions(st1)
    // union sevral sets
    st2 := bitset.New(100)
    st2.Add(uint(4), uint(5), uint(6))
    st3 := bitset.New(100)
    sl1 := []uint{1,2,3,4,5,6}
    st3.Add(sl1...)
    st.Unions(st2, st3)
}
```

License
----

MIT