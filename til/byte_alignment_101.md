#### Refer
```yaml
https://www.eventhelix.com/embedded/byte-alignment-and-ordering/
```

#### Compiler Byte Padding
```yaml
- Compilers have to follow the byte alignment restrictions defined by the target microprocessors
- This means that compilers have to add pad bytes into user defined structures
- So that the structure does not violate any restrictions imposed by the target microprocessor
```

#### User Structure vs Compiler Re-Defined by Compiler
```c
struct Message
{
  short opcode;                // assume 2 byte
  char subfield;               // assume 1 byte
  long message_length;         // assume 4 byte
  char version;
  short destination_processor;
};
```

```c
struct Message
{
  short opcode;         // 2 byte
  char subfield;        // 1 byte
  char pad1;            // Pad to start the long word at a 4 byte boundary
  long message_length;  // 4 byte
  char version;         // 1 byte
  char pad2;            // Pad to start a short at a 2 byte boundary
  short destination_processor;
  char pad3[4];         // Pad to align the complete structure to a 16 byte boundary
};
```

```yaml
- If the above message structure was used in a different compiler/microprocessor combination
- the pads inserted by that compiler might be different
- Thus two applications using the same structure definition header file
- might be incompatible with each other

- Thus it is a good practice to insert pad bytes explicitly in all C-structures
- that are shared in a interface between machines differing in either the compiler
- and/or microprocessor
```

#### Structure Alignment for Efficiency
```yaml
- Sometimes array indexing efficiency can also determine the pad bytes in the structure
- Note that compilers index into arrays by calculating the address of the indexed entry:
  - by the multiplying the index with the size of the structure
- This number is then added to the array base address to obtain the final address
- Since this operation involves a multiply, indexing into arrays can be expensive
- Array indexing can be speeded up by making sure that the structure size is a power of 2
- The compiler can then replace the multiply with a simple shift operation
```
