### Array of strictly the characters 'R', 'G', and 'B'
Segregate the values of the array so that all the 
- Rs come first, 
- Gs come second, &
- Bs come last

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
  var low int
  for idx, c := range str {
    if c == 'R' && low == idx {
      low++
    } else if c == 'R' && idx > low {
      str[low], str[idx] = str[idx], str[low]
      low++
    }
  }
  return str
}
```
