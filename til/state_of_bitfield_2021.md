
```yaml
- https://github.com/emef/bitfield/blob/master/bitfield.go
- https://golangexample.com/go-structure-annotations-that-supports-encoding-and-decoding/
```

### defining arbitrary-width fields within a structure
```go
// https://stackoverflow.com/questions/5793098/go-bitfields-and-bit-packing

type my_chunk uint32

func (c my_chunk) A() uint16 {
  return uint16((c & 0xffff0000) >> 16) 
}

func (c *my_chunk) SetA(a uint16) {
  v := uint32(*c)
  *c = my_chunk((v & 0xffff) | (uint32(a) << 16))
}

func main() {
  x := my_chunk(123)
  x.SetA(12)
  fmt.Println(x.A())
}
```
