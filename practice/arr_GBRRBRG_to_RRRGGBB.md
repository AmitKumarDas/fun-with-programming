### Array of strictly the characters 'R', 'G', and 'B'
Segregate the values of the array so that all the 
- Rs come first, 
- Gs come second, &
- Bs come last

#### Tags
`3-pointer` `loop` `swap` `mid<=high`

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
  // --
  // %c is used as formatting option
  // --
  fmt.Printf("%c\n", CharArrange([]rune{'G', 'B', 'R', 'R', 'B', 'R', 'G'}))
}
```
#### Can we optimise?
```bash
- Try 3 pointers In a Single Loop
- low, mid & high
- low to mid    - contains all Rs
- mid to high   - contains all Gs
- high onwards  - contains all Bs
```
```go
func CharArrange(str []rune) []rune {
  size := len(str)
  if size <= 1 {
    return str
  }
  
  var low, mid, high int
  high = size - 1
  
  // --
  // Think Why The Cond Should Be <=
  // - Even though 'mid' equals 'high'
  // - 'low' & 'mid' might need swaps
  // - Remember there are 3 pointers
  // --
  for mid <= high {
    if str[mid] == 'R' {
      str[low], str[mid] = str[mid], str[low]
      low++
      mid++
    } else if str[mid] == 'G' {
      mid++
    } else if str[mid] == 'B' {
      str[high], str[mid] = str[mid], str[high]
      high--
    }
  }
  return str
}
```
#### Can we solve when there are 4 chars
```bash
- A, B, C, D are only valid array chars
- Output should follow the order D then C then A then B
```
```bash
- Input:  []rune{'A', 'C', 'B', 'D', 'D', 'D', 'B', 'C', 'B', 'C', 'A', 'A', 'D'}
- Output: []rune{'D', 'D', 'D', 'D', 'C', 'C', 'C', 'A', 'A', 'A', 'B', 'B', 'B'}
```
