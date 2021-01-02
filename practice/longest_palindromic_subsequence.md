### Length of Longest Palindromic SubSequence
`dp` `palindrome` `substr == dp` `loop` `equation`

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
Q/ Dos vs. Donts for Fn

- Do Not Think In Terms Of:
-- Pointers & Their Movements
-- Fn Depends on Include & Exclude

- Think In Terms of:
-- Fn Depends on Include Fn & Exclude Fn
```
```bash
Q/ Right Way to Derive Formula?

- Think Using Numbers like 0, 1, 2, & N
- Question Talks of Longest
  - Hence max(...) may be used

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
#### Source Code - Recursive
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
  if j-i == 1 && str[i] == str[j] {
    return 2
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
  return max(
    longestLengthPalinSubSeq(str, i, j-1),
    longestLengthPalinSubSeq(str, i+1, j),
  )
}
```
```bash
Q/ Can we optimise?

- Lets Try Like We Optimized Fibonacci
- Note dp[i][j] implies substr(i,j)
- Set dp[i][i] = 1
- Set dp[i][j] = 2 if j == i+1 && str[i]==str[j]
- Set dp[i][j] = 3 if j == i+2 && str[i]==str[j]

- Last One is Not Required
- However Formula Should Accomodate Last Logic
```
```go
func longestLenPalinSubSeq(str string) int {

  var size = len(str)
  var dp = make([][]int, size)
  for i := range dp {
    dp[i] = make([]int, size)
  }

  for i := range dp {
    // ---
    // dp[i][j] represents substring
    // Since i == j is one char
    // Value is 1 since each char is a Palin
    // ---
    dp[i][i] = 1
  }

  for i:=size-1; i>=0; i-- {
    for j:=i+1; j<size; j++ {
      // ---
      // FILL
      // ---
      if j == i+1 && str[i] == str[j] {
        dp[i][j] = 2 // when substring is 2 & same chars
      } else if str[i] == str[j] {
        dp[i][j] = dp[i+1][j-1] + 2
      } else {
        dp[i][j] = max(
          dp[i+1][j],
          dp[i][j-1],
        )
      }
    }
  }

  return dp[0][size-1]
}
```
