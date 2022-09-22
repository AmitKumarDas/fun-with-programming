package tests

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceHeader(t *testing.T) {
	// Slice with pointer 0 and length 0 is exactly the nil slice
	t.Run("SliceHeader{Data: 0, Len: 0, Cap: 0}) is nil slice", func(t *testing.T) {
		// https://go.dev/play/p/hsIOG1wYBhc
		var s []string = nil
		h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		requireTrue(t, h != nil)
		requireTrue(t, *h == reflect.SliceHeader{Data: 0, Len: 0, Cap: 0}) // TIL
	})
}
