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
```

```yaml
- https://dev.to/jlauinger/sliceheader-literals-in-go-create-a-gc-race-and-flawed-escape-analysis-exploitation-with-unsafe-pointer-on-real-world-code-4mh7
```

```go
// --------------------------
// !!! Incorrect WAY !!!
// --------------------------

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
    //
    // Assume s is defined as follows before being passed here:
    //
    // reader := bufio.NewReader(strings.NewReader("abcdefgh"))
    // s, _ := reader.ReadString('\n')
    //
    // Let us get back to this line of code. There is no reference to s anymore.
    // Thus, s does not escape to the heap. In addition, sliceHeader does not
    // resemble a valid reference to s by the compiler. Hence, the caller to
    // this function will get garbage data in form of []byte slice due to 
    // invalid lifetime of s.
    //
    // If we had created the string from a string literal, like s := "abcdefgh, 
    // then s would have been allocated neither on the heap nor on the stack.
    // Instead, it would have been a string constant in the constant data 
    // section of the resulting binary, and therefore the reference to 
    // that data would have continued to work after returning to the 
    // caller
    return *(*[]byte)(unsafe.Pointer(sliceHeader)) // star & star
}

func main() {
    s := "Hello"
    b := unsafeStringToBytes(&s) // You get READ ONLY slices

    // Attempt to change a read-only memory page
    // The operating system will prevent this 
    // And the result is a SIGSEGV segmentation fault, crashing the program 
    b[1] = "a"
    fmt.Println(b)
}
```

```go
// ------------------
// Correct Way To CAST slices minus Copying
// ------------------

func saferStringToBytes(s *string) []byte {
    // Create an actual slice
    // This ensures that Go will treat the address stored in sliceHeader.Data
    // as if it were a "real" pointer
    bytes := make([]byte, 0, 0) // An allocation & hence escape to heap?

    // Create the string and slice headers by casting
    stringHeader := (*reflect.StringHeader)(unsafe.Pointer(s))
    sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))

    // Set the slice's length and capacity temporarily to zero
    // This is unnecessary here because the slice is already initialized as zero
    //
    // But if you are reusing a different slice this is important
    // Decreasing the length and capacity is a safe operation
	//
    // ----
    // QUIRK: GC
    // ----
    // It ensures that if the garbage collector runs just after the switch of Data 
    // it will not run past the slice end
    sliceHeader.Len = 0
    sliceHeader.Cap = 0

    // ----
    // QUIRK: Ordering is important
    // ----
    // ORDER: Step 1: Change the slice header data address
    sliceHeader.Data = stringHeader.Data
    // ORDER: Step 2: Set slice capacity
    sliceHeader.Cap = stringHeader.Len
    // ORDER: Step 3: Set slice length
    sliceHeader.Len = stringHeader.Len

    // -----
    // TIL
    // -----
    // Use the keep alive dummy function to make sure the original string s is not 
    // freed up until this point
    //
    // This ensures that the underlying data array will not be freed before it is
    // referenced by the bytes slice.
    runtime.KeepAlive(s)  // or runtime.KeepAlive(*s)

    // Return the valid bytes slice (still read-only though)
    return bytes
}
```

```go
// !!! Why Not A Single Statement Instead i.e. In Place Cast !!!
func stringtoBytes(s *string) []byte {
    stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	
    // Below single statement might be simple & good to handle above GC 
    // quirks. However, s may NOT escape to heap. Hence, there are 
    // chances of garbage data in slice at the caller's end.
    b := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{ // in-place cast
        Data: stringHeader.Data,
        Cap: stringHeader.Len,
        Len: stringHeader.Len,
    }))	
}
```

### Sample Commits - 1 (from Go 1.17 onwards)
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

### Sample Commits - 2 (from Go 1.17 onwards)
```yaml
- https://github.com/tetratelabs/tinymem/pull/3
```

```diff
- buf := *(*[]byte)(unsafe.Pointer(internal.SliceHeader(ptr, size)))
+ buf := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size) 
```

### Sample Commits - 3 (from Go 1.17 onwards)
```yaml
- https://github.com/olivere/elastic/pull/1434
- Use unsafe bytes to string to reuse memory i.e. reduce allocations
```

### Sample Commits - 4 (from Go 1.17 onwards)
```yaml
- https://go.dev/src/strings/builder.go
- Go's string builder makes use of unsafe to reduce allocations
- TODO: Should we have a dedicated thought page on builder.go & its commits?
```

### Sample Commits - 5 (from Go 1.17 onwards)
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
