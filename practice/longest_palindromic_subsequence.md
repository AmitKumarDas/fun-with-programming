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
```bash
Q/ Why not start with multiple loops?
Q/ Perhaps 2 pointers?

- Loop Body needs proper Conditions
- Formula provides Conditions
```
```bash
Q/ Why not start thinking from recursion?

- Recursion needs proper Exit Criteria
- Formula Provides Exit Criteria
```
```bash
Q/ Why Formula?

- Visualize x & y axes
- X axis & Y axis are represented by same string
- This enables comparisons
- x & y represent a 2D Table
- Hence, we need Formula to fill cells
```
```bash
Q/ Should Formula be based on Include / Exclude?

- Do not use Pointers & Their Movements
- Think of Include Equation 
- Think of Exclude Equation
```
```bash
Q/ Right Way to Derive Formula?

- Question Talks of Longest
  - Hence maximum(...) to be used

- Let n = len(str)
- Consider str[0] till str[n-1]

- Let F(0, n-1) ~ longest palin subsequence
- if str[0] == str[n-1]
  - F(0, n-1) = F(1, n-2) + 2
- if str[0] != str[n-1]
  - F(0, n-1) = max(F(0, n-2), F(1, n-1))
```
```bash
Q/ Any Patterns?

- Above looks like Fibonacci!
- Seems like we dont need a memo table!
- Should couple of memo variables suffice?
```
#### Source Code - Attempt 1
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
#### Source Code - Attempt 2
```go
func LongestLengthPalinSubSeq(str string) int {
  if len(str) <= 1 {
    return len(str)
  }
  return longestLengthPalinSubSeq(str, 0, len(str)-1)
}

func longestLengthPalinSubSeq(str string, i, j int) int {
  if i > j {
    return 0
  }
  if i == j { // i.e. 1 char
    return 1
  }
  if j-i == 1 { // i.e. 2 chars
    if str[i] == str[j] {
      return 2
    } else {
      return 0
    }
  }
  // ---
  // Formula Depends on
  // Formula For Match &
  // Formula for MisMatch
  // ---
  
  // Match
  if str[i] == str[j] {
    return longestLengthPalinSubSeq(str, i+1, j-1) + 2
  }
  // MisMatch
  return maximum(
    longestLengthPalinSubSeq(str, i, j-1),
    longestLengthPalinSubSeq(str, i+1, j),
  )
}
```
