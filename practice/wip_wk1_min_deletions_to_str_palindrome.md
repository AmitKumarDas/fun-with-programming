### Find the Min Number Of Deletions to Convert A String into Palindrome

```bash
- Input: abdcba
- Output: 1 // Remove either d or c
```

#### Tips
```bash
- No need to set max deletions which is m-1
- Diagonal Up
- Fill i,i as well as i,i+1
- Get the Equation correct before solving
```

```bash
- Dynamic Programming

- O(N^2) runtime
- O(N^2) space

- dp[i][j]        // Refers to substring [i..j]
- dp[i][j] = m-1  // m = len(str) // Max deletions to Palin

- dp[i][i] = 0    // Each char is a Palindrome

- dp[i][i+1] = 0  // if str[i] == str[i+1]
- dp[i][i+1] = 1  // if str[i] != str[i+1] i.e. 1 deletion

- dp[i][j] = dp[i+1][j-1]                     // if str[i] == str[j]
- dp[i][j] = 1 + min(dp[i][j-1], dp[i+1][j])  // if str[i] != str[j]

- dp[0][n-1] // has the answer
```

#### Source Code - Dynamic Programming
```go
func MinDelsToPalin(given string) int {
  size := len(given)
  if size <= 1 {
    return 0
  }
  
  var dp = make([][]int, size)
  for i := range dp {
    dp[i] = make([]int, size)
    dp[i][i] = 0
    if i < size-1 {
      if given[i] == given[i+1] {
        dp[i][i+1] = 0 // 0 deletions
      } else {
        dp[i][i+1] = 1 // 1 deletions
      }
    }
  }
  
  for i:=size-2; i>=0; i-- {
    for j:=i+1; j<size; j++ {
      if given[i] == given[j] {
        dp[i][j] = dp[i+1][j-1]
      } else {
        dp[i][j] = 1 + min(
          dp[i+1][j],
          dp[i][j-1],
        )
      }
    }
  }
  return dp[0][size-1]
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}
```

#### Source Code - Recursion

#### Test
```go
func main() {
  fmt.Printf("abcdba  -%d\n", MinDelsToPalin("abcdba"))
  fmt.Printf("abcba   -%d\n", MinDelsToPalin("abcba"))
  fmt.Printf("abba    -%d\n", MinDelsToPalin("abba"))
  fmt.Printf("abb     -%d\n", MinDelsToPalin("abb"))
  fmt.Printf("bbb     -%d\n", MinDelsToPalin("bbb"))
  fmt.Printf("bb      -%d\n", MinDelsToPalin("bb"))
  fmt.Printf("b       -%d\n", MinDelsToPalin("b"))
  fmt.Printf("b b     -%d\n", MinDelsToPalin("b b"))
}
```
