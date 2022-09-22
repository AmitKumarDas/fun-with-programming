package tests

// Use unsafe.Slice to create a slice whose backing array is a memory buffer returned from C code
// Use unsafe.Slice  to create a slice whose backing array is a memory buffer returned from call such as syscall.MMap
// refer - https://github.com/DanEngelbrecht/golongtail/pull/231
//
// package xyz
//
// // #cgo CFLAGS: -g -std=gnu99 -m64 -msse4.1 -maes -pthread -O3
// // #include "golongtail.h"
// import "C"
//
// func carray2slice64(array *C.uint64_t, len int) []uint64 {
// 	return unsafe.Slice((*uint64)(array), len)
// }
//
// func carray2slice32(array *C.uint32_t, len int) []uint32 {
// 	return unsafe.Slice((*uint32)(array), len)
// }
//
// func carray2slice16(array *C.uint16_t, len int) []uint16 {
// 	return unsafe.Slice((*uint16)(array), len)
// }
//
// func carray2sliceByte(array *C.uint8_t, len int) []byte {
// 	return unsafe.Slice((*byte)(array), len)
// }
//
