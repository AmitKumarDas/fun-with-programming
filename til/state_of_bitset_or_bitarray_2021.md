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

### References
```yaml
- https://github.com/yourbasic/bit
```

