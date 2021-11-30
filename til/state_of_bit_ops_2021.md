### Crash Course - Bit Operations
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

```go
  fmt.Println(4 << 1) // 8, a power of 2
  fmt.Println(4 << 2) // 16, a power of 2
  fmt.Println(3 << 1) // 6, a power of 2
  fmt.Println(3 << 2) // 12, a power of 2
```

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

### References
- https://www.techiedelight.com/round-next-highest-power-2/
- https://www.techiedelight.com/Tags/Bit-Hacks/
- https://www.techiedelight.com/Category/Binary/
