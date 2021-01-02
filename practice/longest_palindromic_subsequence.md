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

#### Lets Try DP Table
- Loop the string twice
  - For each char i you get below filled
  - dp[i][j]=j-i where str[i]==str[j]
  - dp[i][j]=0 where str[i]!=str[j]

- for each char
  - l = max(
      longestPalinSubStrLen(str[i:len], 
      longestPalinSubStrLen(str[i+1:len], 
    )
