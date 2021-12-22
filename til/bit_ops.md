### Bit Operators' Table
```
<<   Left Shift
>>   Right Shift
|    Bitwise OR
&    Bitwise AND
^    Exclusive OR
~    One's Complement i.e. Negate Each Bit
1+~x Two's Complement of x
```

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

### Negative Numbers

#### Signed Representation in Computer
```yaml
- signed data is represented in two's complement
```

#### One's Complement of a Number
```yaml
- ~x is one's complement of x
- negate every binary bit of the number
- ~ negates every bit i.e. all 0s to 1s & all 1s to 0s
```

### Bit Representation - Go vs C
```yaml
- x &^ y // in go
- x & ~y // same thing in c
```

#### Two's Complement of a Number
```yaml
- add 1 to one's complement
```

### Extract jth Bit of a number w
```c
// r can be either 1 or 0
r = (w>>j)&1 // shift w to right by j places & bitwise AND with 1
```

### Even or Odd
```c
if (x&1)
  printf("odd\n");
else
  printf("even\n");
```


### When is Masking used
```yaml
- mask:
- to find specific bit
- to find a group of bits
```

#### Access jth bit
```c
// weird that (1<<j) is defined as mask

#define MASK(j) (1<<j)
int getBit(int w, unsigned j) {
  return (( w & MASK(j)) == 0) ? 0 : 1; // why not return (w & MASK(j))
}
```

#### Get ith bit of w
```c
// aliter: write a macro
#define getBit(w,i) ((w>>i)&1)
```

#### Set Specific Bit of the Word
```c
#define MASK(j) (1<<j)
int setBit(int w, unsigned j, short value){
  if (value == 0) return (w & ~MASK(j));
  else if (value == 1) return w | MASK(j);
  else return w;
}
```

```yaml
- So to set the bit 4 of w = 000000000 to 1, we can use
- w = setBit(w,4,1)
- The result is w = 000010000
```

### 0x0F
```yaml
- 0x0F represents a hex value of 1 byte
- rightmost 4 bits are all 1s
- i.e. first 4 bits are all 1s
- leftmost 4 bits are all 0s
- i.e. next 4 bits are all 0s
- i.e. 00001111
```

#### Work with 0x0F
```c
r = x | 0x0F // set first 4 bits (i.e. right most) of x to 1
r = x & 0x0F // set left most 4 bits of x to 0
```

### Right Shift
```go
  fmt.Println(1 >> 1) // 0
  fmt.Println(2 >> 1) // 1
  fmt.Println(3 >> 1) // 1
  fmt.Println(4 >> 1) // 2
```

```go
  fmt.Println(1 >> 1) // 0
  fmt.Println(1 >> 2) // 0
  fmt.Println(1 >> 3) // 0
  fmt.Println(1 >> 4) // 0
```

```go
  fmt.Println(4 >> 1) // 2
  fmt.Println(4 >> 2) // 1
  fmt.Println(4 >> 3) // 0
  fmt.Println(4 >> 4) // 0
```

### Left Shift
```go
  fmt.Println(4 << 1) // 8, a power of 2
  fmt.Println(4 << 2) // 16, a power of 2
  fmt.Println(3 << 1) // 6, a power of 2
  fmt.Println(3 << 2) // 12, a power of 2
```

### Next Number Power of 2

```yaml
- For n = 12 => Next number power of 2 is 16
- For n = 20 => Next number power of 2 is 32
```

```cpp
unsigned findNextPowerOf2(unsigned n)
{
  // decrement n (to handle the case when n itself is a power of 2)
  n = n - 1;

  // do till only one bit is left
  while (n & n - 1) {
    n = n & n - 1;        // unset rightmost bit
  }

  // n is now a power of two (less than n)
  // return next power of 2
  return n << 1;
}
```

```cpp
// Compute power of two greater than or equal to `n`
unsigned findNextPowerOf2(unsigned n)
{
  // decrement n (to handle the case when n itself is a power of 2)
  n = n - 1;

  // initialize result by 2
  int k = 2;

  // double k and divide n in half till it becomes 0
  while (n >>= 1) {
    k = k << 1;    // double k
  }

  return k;
}
```

```cpp
// Compute power of two greater than or equal to n
unsigned findNextPowerOf2(unsigned n)
{
  // decrement n (to handle the case when `n` itself is a power of 2)
  n = n - 1;

  // calculate the position of the last set bit of n
  int lg = log2(n);

  // next power of two will have a bit set at position `lg+1`
  return 1U << lg + 1;
}
```

```yaml
- set all bits on the right-hand side of the most significant set bit to 1 
- then increment the value by 1 to “rollover” to two’s nearest power

- e.g, consider number 20
- convert its binary representation 00010100 to 00011111 
- and add 1 to it which results in the next power of 2 for number 20
- i.e. (00011111 + 1) = 00100000
```

```cpp
// Compute power of two greater than or equal to `n`
unsigned findNextPowerOf2(unsigned n)
{
  // decrement n (to handle the case when n itself is a power of 2)
  n--;

  // set all bits on the right hand side of the most significant bit set to 1
  n |= n >> 1;
  n |= n >> 2;
  n |= n >> 4;
  n |= n >> 8;
  n |= n >> 16;

  // increment n and return
  return ++n;
}
```

```go
func nextPowerOf2(x uint32) uint32 {
	if x == math.MaxUint32 {
		return x
	}

	if x == 0 {
		return 1
	}

	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16

	return x + 1
}
```

### Bitmask
```yaml
- A bitmask is a small SET of booleans
- Often called flags
- Represented by the bits in a single number
```
```go
type Bits uint8 // a single number

const (
  F0 Bits = 1 << iota // 1 << 0
  F1                  // 1 << 1
  F2                  // 1 << 2
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }     // b & ~flag
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

func main() {
  var b Bits
  b = Set(b, F0)
  b = Toggle(b, F2)
  for i, flag := range []Bits{F0, F1, F2} {
    fmt.Println(i, Has(b, flag))
  }
}
```


### References
- https://www.techiedelight.com/round-next-highest-power-2/
- https://www.techiedelight.com/Tags/Bit-Hacks/
- https://www.techiedelight.com/Category/Binary/
- https://yourbasic.org/golang/bitmask-flag-set-clear/
