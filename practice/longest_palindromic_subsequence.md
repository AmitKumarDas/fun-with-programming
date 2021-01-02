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

#### Generalized Formula For SubSequence
- Question Talks of Longest
  - Hence maximum(...) to be used

- Let n = len(str)
- Consider str[0] till str[n-1]

- Let F(0, n-1) ~ longest palin subsequence
- if str[0] == str[n-1]
  - F(0, n-1) = F(1, n-2) + 2
- if str[0] != str[n-1]
  - F(0, n-1) = max(F(0, n-2), F(1, n-1))

- Above looks like Fibonacci!
- Seems like we dont need a memo table!
- Should couple of memo variables suffice?

#### Source Code
```go
func LongestLengthPalinSubSeq(str string) int {
  if len(str) <= 1 {
    return len(str)
  }
  return longestLengthPalinSubSeq(str, 0, len(str))
}

func longestLengthPalinSubSeq(str string, i, n int) int {
  if n-i <= 1 {
    return n-i
  }
  if n-i == 2 {
    if str[i] == str[n-1] {
      return 2
    } else {
      return 0
    }
  }
  for i <= n {
    if str[i]==str[n-1] {
      return longestLengthPalinSubSeq(str, i+1, n-1) + 2
    } else {
      return max(
        longestLengthPalinSubSeq(str, i, n-1),
        longestLengthPalinSubSeq(str, i+1, n),
      )
    }
  }
}
```
