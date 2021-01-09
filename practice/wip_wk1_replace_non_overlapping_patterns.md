### Replace All Non Overlapping Occurences Of The Pattern With a Char

```bash
Given: Text & Pattern & Char
TODO:  Replace All Non Overlapping Patterns With a Char
```

```bash
- Sample

- Given:  aaaa, aa, b
- Output: bb

- Given:  abaa, aa, b
- Output: abb

- Given:  abaa, aa, a
- Output: aba

- Given:  aaaa, aa, a
- Output: aa
```

```bash
- How

- start from low & high & pidx
- for high < size
  - if match i.e. pattern[pidx]==str[high]
    - if complete match: resp[low]=Char, low=high+1, pidx=0
    - if partial match pidx++
  - if no match pidx=0 resp=append(resp, str[low:high+1]), low=high+1
  - high++
```

#### Source Code
```go
func RplPattern(str string, pat string, c rune) []rune {
  var size = len(str)
  if size < len(pat) {
    return []rune(str)
  }
  
  var low, high, pidx int
  
  // --
  // Is []rune a Good Approach?
  // Is Pure Append a Good Approach Here?
  // --
  var resp []rune
  for high < size {
    // --
    // Match
    // --
    if str[high] == pat[pidx] {
      // --
      // is full match?
      // --
      if pidx == len(pat)-1 {
        resp = append(resp, c)
        low = high+1
        pidx=0
      } else {
        pidx++
      }
    } else {
      // --
      // No Match
      // --
      for low <= high {
        resp = append(resp, rune(str[low]))
        low++
      }
      pidx=0
      low=high+1
    }
    high++
  }
  return resp
}
```
#### Test
```go
func main() {
  fmt.Printf("aaaa, aa, b = %c\n", RplPattern("aaaa", "aa", 'b'))
  fmt.Printf("abaa, aa, b = %c\n", RplPattern("abaa", "aa", 'b'))
  fmt.Printf("aa, aa, b = %c\n", RplPattern("aa", "aa", 'b'))
  fmt.Printf("aaaa, a, b = %c\n", RplPattern("aaaa", "a", 'b'))
  fmt.Printf("aaaa, aa, a = %c\n", RplPattern("aaaa", "aa", 'a'))
  fmt.Printf("aabbaabb, a, b = %c\n", RplPattern("aabbaabb", "aa", 'b'))
  fmt.Printf("aaaa, ab, b = %c\n", RplPattern("aaaa", "ab", 'b'))
}
```
