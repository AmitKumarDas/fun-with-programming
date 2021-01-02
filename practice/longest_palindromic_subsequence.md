### Length of Longest Palindromic SubSequence

```bash
Input:    geeksforgeeks
Output:   5
Details:  eekee, eesee, eefee, eeoee, eeree, eegee
```
```bash
Input:    bbabcbcab
Output:   7
Details:  babcbab
```

#### Lets Try Recursion
```go
func longestPalinSubSeq(str string) int {
  if len(str) <= 1 || isPalin(str) {
    return len(str)
  }
  return 0
}

func LongestPalinSubSeq(str string) int {
  var max int
  for idx := range str {
    new := str[0:idx] + str[idx+1:len(str)]
    got = maximum(
      longestPalinSubSeq(new), // exclude char
      longestPalinSubSeq(str), // include char
    )
    if got > max {
      max = got
    }
  }
}
```
