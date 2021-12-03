### BitArray or BitSet

### Golang Types & Sizes
```yaml
- The size of the generic int & uint type is platform dependent
- It is 32 bits wide on a 32-bit system and 64-bits wide on a 64-bit system

- byte and rune that are aliases for uint8 and int32 data types respectively

- byte: alias for uint8
- represent ASCII chars
- mnemonics: byte u into 8 pieces & throw the rest to Sky

- rune: alias for int32
- represent broader set of Unicode chars that are encoded in UTF-8 format
- mnemonics: run for international MaN

- Go does not have char datatype
- It uses byte & rune to represent char values

- Both byte and rune data types are essentially integers
- For example, a byte variable with value 'a' is converted to the integer 97
```

```go
var firstLetter = 'A' // Type inferred as 'rune' (Default type for character values)
var lastLetter byte = 'Z' // explicit declaration
```

#### Grok - 1
```go
// https://github.com/yourbasic/bit/blob/master/set.go
// deals with int
// backed by uint64

const (
  bpw   = 64         // bits per word
  maxw  = 1<<bpw - 1 // maximum value of a word
  shift = 6
  mask  = 0x3f       // 0b111111
)
```

#### Grok - 2
```go
// Given: A set of ints with a maximum value as max
// When: You need to store these ints
// Then: You need following allocation

data: make([]uint64, max>>shift+1)
```

#### Grok - 3
```go
type Set {
  data []uint64
}

if n >= 0 {
  // s.data[bucket] = existing value | current bit value
  // where current bit value = uint(n&mask)
  //
  // n is of int data type
  // i.e. is platform dependent for its size
  // i.e. n can be either 32-bit or 64-bit
  // hence max value of int type can be masked via 0x3F
  // value is a multiple of 2 i.e. << (left shift)
  // note: value is not always a power of 2
  //
  // note: storage array index i.e. bucket index can be a low value
  // i.e. bucket index is representative of division by 2 i.e. >> (right shift)
  //
  // note: multiple values can be set in same bucket
  // how many multiple values can be accommodated in the same bucket?
  // do these bucket values range from n*2 to (n-1)*2 -1
  // these values are OR-ed before being saved
  s.data[n>>shift] |= 1 << uint(n&mask)
}
```

### References
```yaml
- https://github.com/yourbasic/bit
```

