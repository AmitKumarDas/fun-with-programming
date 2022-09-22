package tests

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceHeaderOfNilArray(t *testing.T) {
	// https://go.dev/play/p/hsIOG1wYBhc
	var s []string = nil
	h := (*reflect.SliceHeader)(unsafe.Pointer(&s))

	requireTrue(t, *h == reflect.SliceHeader{Data: 0, Len: 0, Cap: 0})
}
