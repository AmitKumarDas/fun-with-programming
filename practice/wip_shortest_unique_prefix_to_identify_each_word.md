### Find Shortest Unique Prefixes
That Identify Each Word Uniquely From A Given Array

```bash
Assumption: No Word Is Prefix Of Another
```

```bash
Given:  [hello, there, how, are, you]
Output: [he, t, ho, a, y]
```

```bash
- How

- Since Trie Can Handle Prefixes Well
- Use Trie DataStructure
```

```bash
- If words Hi, Her & Hello are given
- Take Hi
  - h has 2 Nested Tries - move on
  - i has 0 Nested Trie - so hi is a prefix
- Take Her
  - h has 2 Nested Tries - move on
  - e has 2 Nested Trie - move on
  - r has 0 Nested Trie - so her is a prefix
- Take Hello
  - h has 2 Nested Tries - move on
  - e has 2 Nested Tries - move on
  - l has 1 Nested Trie - so hel is a prefix
```

#### Source Code
```go
// --
// Trie Does Not Need a Val? Really?
// --
type Trie struct {
  Next        []*Trie
  IsEnd       bool
}

func (t *Trie) Add(word string) {
  if word == "" {
    return
  }

  var tmp = t
  for idx, c := range word {
    var pos = c - 'a'
    if tmp.Next[pos] == nil {
      tmp.Next[pos] = &Trie{Next: make([]*Trie, 26)}
    }
    if idx+1 == size(word) {
      tmp.Next[pos].IsEnd = true  
    }
    tmp = tmp.Next[pos]
  }
}

func (t *Trie) GetUniqPrefixFor(word string) string {
  if word == "" {
    return ""
  }
  
  var prefix []rune
  for _, c := range word {
    prefix = append(prefix, c)
    var pos = c - 'a'
    var tre = t.Next[pos]
    if tre == nil || len(tre.Next) <= 1 || tre.IsEnd {
      break
    }
  }

  return string(prefix)
}

func UniqPrefixes(words []string) []string {
  var root = &Trie{Next: make([]*Trie, 26)}
  for _, word := range words {
    root.Add(word)
  }
  
  var resp []string
  for _, word := range words {
    resp = append(resp, root.GetUniqPrefixFor(word))
  }
  
  return resp
}
```

#### Test
```go
func main() {
  fmt.Printf("%v\n", UniqPrefixes([]string{"Hi", "Her", "Hello"}))
}
```

#### References
- https://www.techiedelight.com/shortest-unique-prefix/
