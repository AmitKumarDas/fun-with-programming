### Solve the Dutch Flag problem When Number of Valid Chars is 4

#### Tags
- `4-Pointer` `O(n)` `O(1)`

```bash
- Consider A, B, C, D as only valid array chars
- Output should follow the order D then C then A then B
```

```bash
- Input:  []rune{'A', 'C', 'B', 'D', 'D', 'D', 'B', 'C', 'B', 'C', 'A', 'A', 'D'}
- Output: []rune{'D', 'D', 'D', 'D', 'C', 'C', 'C', 'A', 'A', 'A', 'B', 'B', 'B'}
```

```go
func CharArrange(str []rune) []rune {
  var size = len(str)
  if size <= 1 {
    return str
  }
  
  var low, mid1, mid2, high int
  high = size-1
  
  for mid2 <= high {
    if str[mid2] == 'D' {
      str[low], str[mid2] = str[mid2], str[low]
      low++
      mid1++
      mid2++
    } else if str[mid2] == 'C' {
      str[mid1], str[mid2] = str[mid2], str[mid1]
      mid1++
      mid2++
    } else if str[mid2] == 'A' {
      mid2++
    } else {
      str[mid2], str[high] = str[high], str[mid2]
      high--
    }
  }
  return str
}
```
```go
// ---
// Test
// ---
func main() {
  // --
  // %c is used as formatting option
  // --
  fmt.Printf(
    "%c\n", 
    CharArrange([]rune{'A', 'C', 'B', 'D', 'D', 'D', 'B', 'C', 'B', 'C', 'A', 'A', 'D'}),
  )
}
```
