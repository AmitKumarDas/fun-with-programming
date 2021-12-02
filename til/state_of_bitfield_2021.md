
```yaml
- https://github.com/emef/bitfield/blob/master/bitfield.go
- https://golangexample.com/go-structure-annotations-that-supports-encoding-and-decoding/
```

### Define Arbitrary Width Fields Within a Structure
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

### Pack Multiple Values into 1 int
#### Problem
```yaml
- https://stackoverflow.com/questions/6556961/use-of-the-bitwise-operators-to-pack-multiple-values-in-one-int

- int age, gender, height, packed_info;
- Pack as AAAAAAA G HHHHHHH
- packed_info = (age << 8) | (gender << 7) | height;

- Unpack
- height = packed_info & 0x7F;   // 0x7F is 01111111
- gender = (packed_info >> 7) & 1;
- age    = (packed_info >> 8);
```

#### Explanation
```yaml
- Pack the age, gender and height into 15 bits, of the format:
- AAAAAAAGHHHHHHH

- To start with, age has this format:
- age           = 00000000AAAAAAA
- where each A can be 0 or 1

- (age << 8)
- age is now    = AAAAAAA00000000

- gender        = 00000000000000G
- (gender << 7) = 0000000G0000000

- height        = 00000000HHHHHHH

- Combine these into one variable:
- Use | operator
- packed_info   = (age << 8) | (gender << 7) | height
- packed_info   = AAAAAAAGHHHHHHH

- Unpack the bits:
- Use & operator
- Get the height:
- packed_info          = AAAAAAAGHHHHHHH
- 0x7F                 = 000000001111111
- (packed_info & 0x7F) = 00000000HHHHHHH = height
- Get the age:
- To get the age, push everything 8 places to the right
- age = (packed_info >> 8)
- age                  = 0000000AAAAAAAA

- Get the gender:
- push everything 7 places to the right to get rid of the height
- We now care the last bit only
- packed_info            = AAAAAAAGHHHHHHH
- (packed_info >> 7)     = 0000000AAAAAAAG
- 1                      = 000000000000001
- (packed_info >> 7) & 1 = 00000000000000G
```
