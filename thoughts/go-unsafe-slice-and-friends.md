## Teaching unsafe's slice & friends to mortals

### NOTES
```yaml
- Use unsafe.Slice to create a slice whose BACKING array is a memory buffer returned
- from C code
- from call such as syscall.MMap
```

```yaml
- *byte is NOT the ONLY way to POINTER to underlying data
- *uint16, *uint32, *uint64, *unit8 can also be pointers to underlying data
- *uint8 is same as *byte
```

```yaml
- If b is []byte then *b[0] is the POINTER to underlying data
- With *b[0] & len(b) we can construct corresponding string or slice
```

```yaml
- Go string builder code can be a useful reference on managing memory allocations
```

### Example - 101
```yaml
- convert a string to a []byte by reusing the string memory i.e. no allocation
- https://dev.to/jlauinger/sliceheader-literals-in-go-create-a-gc-race-and-flawed-escape-analysis-exploitation-with-unsafe-pointer-on-real-world-code-4mh7
```

```go
func unsafeStringToBytes(s *string) []byte { // given a string pointer
    sh := (*reflect.StringHeader)(unsafe.Pointer(s)) // cast & cast
    sliceHeader := &reflect.SliceHeader{
        Data: sh.Data,
        Len:  sh.Len,
        Cap:  sh.Len,
    }

    // At this point, s is no longer used. 
	// However, there is a copy of the address of its underlying array
	// in sliceHeader.Data. Since sliceHeader was not created from an
	// actual slice, the GC does not treat the address as a reference. 
	// Therefore, IF the GC runs here it will FREE s.
	//
	// When GC runs, it will free string s because it is no longer used
	// When the []byte slice is created in the next line, its Data
	// field will contain an invalid address. It might now point
	// to an unmapped memory page, or simply to some undefined
	// position in the heap that might get reused later on.
    return *(*[]byte)(unsafe.Pointer(sliceHeader)) // star & star
}
```

```go
func main() {
    s := "Hello"
    b := unsafeStringToBytes(&s)
  
    // Attempt to change a read-only memory page
	// The operating system will prevent this 
	// And the result is a SIGSEGV segmentation fault, crashing the program 
    b[1] = "a"
    fmt.Println(b)
}
```

### Example - 1
```yaml
- refer - https://github.com/DanEngelbrecht/golongtail/pull/231
```

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

### Example - 2
```yaml
- https://github.com/tetratelabs/tinymem/pull/3
```

```diff
- buf := *(*[]byte)(unsafe.Pointer(internal.SliceHeader(ptr, size)))
+ buf := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size) 
```

### Example - 3
```yaml
- https://github.com/olivere/elastic/pull/1434
- Use unsafe bytes to string to reuse memory i.e. reduce allocations
```

### Example - 4
```yaml
- https://go.dev/src/strings/builder.go
- Go's string builder makes use of unsafe to reduce allocations
```

### Example - 5
```yaml
- https://github.com/google/brotli/pull/942
```

```diff
- // It is a workaround for non-copying-wrapping of native memory.
- // C-encoder never pushes output block longer than ((2 << 25) + 502).
- // TODO(eustas): use natural wrapper, when it becomes available, see
- //               https://golang.org/issue/13656.
- output := (*[1 << 30]byte)(unsafe.Pointer(result.output_data))[:length:length]
+ output := unsafe.Slice((*byte)(unsafe.Pointer(result.output_data)), length)
```
