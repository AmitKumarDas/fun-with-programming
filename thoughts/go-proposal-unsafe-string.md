### 53003 (Accepted)
```yaml
- https://github.com/golang/go/issues/53003
```

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

```yaml
- To use above unsafe functions; we might resort to following
```

```go
func StringToBytes(s string) []byte {
    return unsafe.Slice(unsafe.StringData(s), len(s))
}
```
```go
func BytesToString(b []byte) string {
    return unsafe.String(&b[0], len(b))
}
```
