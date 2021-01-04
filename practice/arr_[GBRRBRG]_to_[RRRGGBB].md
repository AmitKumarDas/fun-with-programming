### Array of strictly the characters 'R', 'G', and 'B'
Segregate the values of the array so that all the 
- Rs come first, 
- Gs come second, &
- Bs come last

#### Tags
`2-pointer` `loop` `swap`

#### Note
```bash
- You can only swap elements of the array
- Do this in linear time and in-place
```

#### Sample
```bash
- Given:  ['G', 'B', 'R', 'R', 'B', 'R', 'G']
- Output: ['R', 'R', 'R', 'G', 'G', 'B', 'B']
```

#### How?
```bash
- Linear Time Means 1 Loop or Multiple Parallel Loops
- InPlace & SWAP go hand-in-hand
```

#### Source Code
```go
func CharArrange(str []rune) []rune {
  // ---
  // Use of 2 Pointer Theory
  // Move all 'R' chars to Left
  // ---
  var low int
  for idx, c := range str {
    if c == 'R' && low == idx {
      low++
    } else if c == 'R' && idx > low {
      str[low], str[idx] = str[idx], str[low]
      low++
    }
  }
  
  // ---
  // Repeat above logic for char 'G'
  // i.e. Move all 'G' chars to Left
  // ---
  var size = len(str)
  var second = low
  for low < size {
    if str[low] == 'G' && low == second {
      second++
    } else if str[low] == 'G' && low > second {
      str[low], str[second] = str[second], str[low]
      second++
    }
    low++
  }
  return str
}

// ---
// Test
// ---
func main() {
  fmt.Printf("%c\n", CharArrange([]rune{'G', 'B', 'R', 'R', 'B', 'R', 'G'}))
}
```
#### Can we optimise?
```bash
- Try a single loop with 2 pointers
```
```go
func CharArrange(str []rune) []rune{
  var size = len(str)
  if size <= 1 {
    return str
  }
  
  var low, high int
  high = size - 1
  for idx := range str {
  
  }
}
```
