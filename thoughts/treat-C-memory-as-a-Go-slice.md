## Treating C memory as a Go slice
In other words Go slice backed by C array

This was part of a very old proposal (_Dec 2015_) that was closed.
Refer - https://github.com/golang/go/issues/13656
Refer - https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices

## Background
- C arrays are typically either null-terminated or have a length kept elsewhere
- Go provides the following function to make a new Go byte slice from a C array
  - `func C.GoBytes(cArray unsafe.Pointer, length C.int) []byte`

### Now to have a Go slice backed by C array WITHOUT allocation
- To create a slice backed by C array WITHOUT COPYING the original data
  - Acquire the length of array at runtime
  - Use a type conversion to a pointer to a VERY BIG ARRAY (**_WHY?_**)
  - Then slice it to the length that you want
  - Remember to set the cap if you're using Go 1.2 or later

```go
import "C"
import "unsafe"

func some() {
  var theCArray *C.YourType = C.getTheArray()
  length := C.getTheArrayLength()
  slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]	
}
```

- With Go 1.17 or later, programs can use unsafe.Slice instead
- Which similarly results in a Go slice backed by a C array
```go
import "C"
import "unsafe"

func some() {
  var theCArray *C.YourType = C.getTheArray()
  length := C.getTheArrayLength()
  slice := unsafe.Slice(theCArray, length) // Go 1.17	
}
```

### Obvious Note / Warning
- The Go garbage collector will not interact with the underlying C array
- And that if it is freed from the C side of things
- Then The behavior of any Go code using the slice is nondeterministic

## An old proposal
- https://github.com/golang/go/issues/13656
- This has some info worth reading
  - Read C memory without making unnecessary copies
  - Able to read (& write) C char* using Go functions that uses []byte

### Technicals - 1
```go
// Convert an unsafe.Pointer p to any []T without copying as follows
(*[(1 << 31) / unsafe.Sizeof(T{})]T)(p)[:len:len]
```
- This isn't portable:
  - 1/ On 64-bit platforms it only expresses sizes up to 2GB
  - 2/ You can't work around it by increasing the constant
    - Since it doesn't work on 32-bit platforms
- This interacts poorly with analysis tools:
- It temporarily produces a pointer to an "array" with elements in invalid memory

### Technicals - 2
- Using SliceHeader is ok
- Note that "cannot be used safely" warning is aimed at trying to store Go pointers in its Data field
- Storing C pointers in the Data field should be ok
- The rest of the warning is basically the same as the one for unsafe.Pointer
- In an ideal world, SliceHeader.Data should be an unsafe.Pointer, not a uintptr (_too late now_)
```go
h := reflect.SliceHeader{uintptr(p), n, n}
s := *(*[]T)(unsafe.Pointer(&h))
```

### Technicals - 3 (Reiterating 1)
```go
// MAX value of C DEPENDS on the SIZE of T
(*[C]T)(unsafe.Pointer(v))[:len:len]

// convert an unsafe.Pointer p to any []T without copying
// Note: These are not portable
(*[(1 << 31) / unsafe.Sizeof(T{})]T)(p)[:len:len]
```