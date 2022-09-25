## Teaching unsafe's slice & friends to mortals

### TIPs
- Use unsafe.Slice to create a slice whose BACKING array is a memory buffer returned
  - from C code
  - from call such as syscall.MMap
- *byte is NOT the ONLY way to POINTER to underlying data
  - *uint16, *uint32, *uint64, *unit8 can also be pointers to underlying data
  - *uint8 is same as *byte
- If b is []byte then *b[0] is the POINTER to underlying data
  - With *b[0] & len(b) we can construct corresponding string or slice
- Go string builder code can be a useful reference on managing memory allocations

### Proposal - 1
- https://github.com/golang/go/issues/53003
```go
// StringData returns a POINTER to the BYTES of a string
//
// The bytes MUST NOT be modified
// Doing so can cause the program to crash or behave unpredictably
func StringData(string) *byte

// String constructs a string VALUE from a pointer and a length
//
// The BYTES passed to String must NOT be modified
// Doing so can cause the program to crash or behave unpredictably
func String(*byte, int) string
```

- To use above unsafe functions; we might resort to following
```go
func StringToBytes(s string) []byte {
    return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
    return unsafe.String(&b[0], len(b))
}
```

### Real World Usage - 1
- refer - https://github.com/DanEngelbrecht/golongtail/pull/231
```go
package xyz

// #cgo CFLAGS: -g -std=gnu99 -m64 -msse4.1 -maes -pthread -O3
// #include "golongtail.h"
import "C"

func cArrToSlice64(array *C.uint64_t, len int) []uint64 {
	return unsafe.Slice((*uint64)(array), len) // *uint64 pointer to underlying data
}

func cArrToSlice32(array *C.uint32_t, len int) []uint32 {
	return unsafe.Slice((*uint32)(array), len) // *uint32 pointer to underlying data
}

func cArrToSlice16(array *C.uint16_t, len int) []uint16 {
	return unsafe.Slice((*uint16)(array), len) // *uint16 pointer to underlying data
}

func cArrToSliceByte(array *C.uint8_t, len int) []byte {
	return unsafe.Slice((*byte)(array), len) // *byte pointer to underlying data
}
```

### Real World Usage - 2
- https://github.com/tetratelabs/tinymem/pull/3
```diff
- buf := *(*[]byte)(unsafe.Pointer(internal.SliceHeader(ptr, size)))
+ buf := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size) 
```

### Real World Usage - 3
- https://github.com/olivere/elastic/pull/1434
  - Use unsafe bytes to string to reuse memory i.e. reduce allocations

### Real World Usage - 4
- https://go.dev/src/strings/builder.go
- Go's string builder makes use of unsafe to reduce allocations
