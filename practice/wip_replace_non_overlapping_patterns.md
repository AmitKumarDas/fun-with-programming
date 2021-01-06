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
```
