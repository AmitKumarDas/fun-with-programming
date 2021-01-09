### Find all Distinct SubStrings Containing Exactly K Distinct Chars

```bash
Input: How do you do, k=4
Output: 'how ', 'ow do '
```

```bash
- What about spaces
- map[rune]int
- low pointer that moves only when
  - a substring is found && more substrings can't be possible
```

```go
func DSubStr(given string, k int) []string {
  var size = len(given)
  if size < k {
    return nil
  }
  
  var low int
  var resp []string

  var count = map[rune]int{}
  var idx, runSum int
  
  for idx < size {
    c := rune(given[idx])
    count[c] += 1
    if count[c] == 1 {
      // New Char
      runSum++
    }

    // --
    // Commented Approach is Wrong
    // --
    //for runSum > k {
    //  lowc := rune(given[low])
    //  if count[lowc] == 1 {
    //    runSum--
    //  }
    //  count[lowc] -= 1
    //  low++
    //}
    //if runSum == k {
    //  resp = append(resp, given[low:idx+1])
    //}
    
    // --
    // Since Multiple Substrings are Possible
    // With K Distinct Chars Then Loop Till
    // runSum == k remains as-is
    //
    // Since SubString & Not SubSequence low
    // Pointer Needs to Move
    // --
    var tmpIdx = idx
    for runSum == k {
      resp = append(resp, given[low:tmpIdx+1])
      tmpIdx++
      curr := rune(given[tmpIdx])
      if count[curr] > 0 {
        // existing
        count[curr] += 1
        continue
      } else {
        lowc := rune(given[low])
        if count[lowc] == 1 {
          low++
        }
        count[lowc] -= 1
      }
    }
  }
  return resp
}
```
```go
func main() {
  fmt.Printf("%v\n", DSubStr("abcadce", 4))
  fmt.Printf("%v\n", DSubStr("abcbd", 3))
  fmt.Printf("%v\n", DSubStr("aa", 1))
}
```

#### References:
- https://www.techiedelight.com/find-all-substrings-containing-exactly-k-distinct-characters/
