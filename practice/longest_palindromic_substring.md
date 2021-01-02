### Longest Palindromic Substring
`DP` `Loop` `Teaser` `MemoTable` `Diagonal`

```bash
Given a string.
Find the Longest SubString which is a Palindrome.
```

#### Optimal Approach - DP
- Every single char is Palindrome
- In DP Present Calc Needs Past Results
- In this case Present Calc is Derived from Future Results
- i.e. dp(i) derived from dp(i+1)

#### How to DP?
- Challenge Is To Fill & Use DP Table Instead of isPalin() func

### What can be `dp[i][j]`?
- The cell indicates a substr(i,j)
- The cell value indicates true or false i.e. is a Palindrome?
- _Assumption:_ dp[i][i] = 1 i.e. diagonal is true for any string
- _Since Single Char is Already a Palindrome_

#### Do This Simple Fills
- dp[i][i] = 1
- dp[i][i+1] = 1 if arr[i+1] == arr[i]

#### Derive Generalized Formula
- for j-i >= 2
- dp[i][j] = dp[i+1][j-1]

#### How Did We Arrive At General Formula?
- dp[i][j] is Palin if & only if dp[i+1][j-1] is Palin
  - String is Palin if its Immediate Substring is Palin
  - Substring is Palindrome if its Immediate Substring is Palin
- Here Substring means _first_ & _last_ chars are Removed

#### How To Fill Up Future Cell Before Current Cell?
- Trick Lies in Filling the dp via Custom Nested Loops
- Do Not fill cells on Left of diagonals
- Fill from Bottom to Top
- Fill cells on Diagonal & its Right

```go
var n = len(str)
for i:=n-1; i>=0; i-- {
  for j:=i; j<=n-1; j++ {
  }
}
```
```go
var n = len(str)
var x, y, max int
var dp = make([][]int, n)

for i:=n-1; i>=0; i-- {
  for j:=i; j<=n-1; j++ {
    // Fill
    if i == j {
      dp[i][j] = 1
    } else if arr[i] == [j] && j-i==1 {
      dp[i][j] = 1
    } else {
      dp[i][j] = dp[i+1][j-1]
    }
    // Check
    if dp[i][j] == 1 && j-i>max{
      max=j-i
      x=i
      y=j
    }
  }
}
```

#### References
- https://www.geeksforgeeks.org/fundamentals-of-algorithms/
- https://www.techiedelight.com/ **
- https://iq.opengenus.org/longest-palindromic-substring-dp/ ***
