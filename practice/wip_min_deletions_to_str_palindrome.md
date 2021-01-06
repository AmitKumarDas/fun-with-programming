### Find the Min Number Of Deletions to Convert A String into Palindrome

```bash
- Input: abdcba
- Output: 1 // Remove either d or c
```

```bash
- Palindrome - Diagonal & Up
- dp[i][j]        // refers to substring
- dp[i][j] = m-1  // m = len(str) // no of deletions to Palin
- dp[i][i] = 0
- dp[i][i+1] = 0  if str[i] == str[i+1]
- dp[i][i+1] = 1  if str[i] != str[i+1]

- dp[i][j] = dp[i+1][j-1] if str[i] == str[j]
- dp[i][j] = 1 + min(dp[i][j-1], dp[i+1][j]) if str[i] != str[j]

- dp[0][n-1] // has the answer
```
