package tests

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestCMemAsGoSlice(t *testing.T) {
	s := []byte("Hello, ⚛")
	n := len(s)
	p := reflect.ValueOf(s).Pointer()

	// Old way
	slice := (*[1 << 30]byte)(unsafe.Pointer(p))[:n:n]
	fmt.Println("slice old way", slice)         // [72 101 108 108 111 44 32 226 154 155]
	fmt.Println("slice old way", string(slice)) // Hello, ⚛

	// New way // Go 1.17
	sliceNew := unsafe.Slice(&s[0], n)             // [0] else it is two-dimensional
	fmt.Println("slice new way", sliceNew)         // [72 101 108 108 111 44 32 226 154 155]
	fmt.Println("slice new way", string(sliceNew)) // Hello, ⚛

	// Old way
	var i uint64 = 0xdeedbeef01020304
	slice2 := (*[1 << 30]byte)(unsafe.Pointer(&i))[:8:8]
	fmt.Println("another slice old way", slice2)         // [4 3 2 1 239 190 237 222]
	fmt.Println("another slice old way", string(slice2)) // ����

	// New way
	slice2New := unsafe.Slice(&i, 8)
	fmt.Println("another slice new way", slice2New[0]) // 16063705379623797508 // a rune
}
