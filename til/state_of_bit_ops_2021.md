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
