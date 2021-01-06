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
- Store Freq Of Visits Based on Prefix Sequence
```

#### Tips
```bash
- If ASCII set then size = 128
- If only small case English chars then:
  - size = 26
  - pos = rune - 'a'
- If only capital case English chars then:
  - size = 26
  - pos = rune - 'A'
- If both case English chars then:
  - size = 128 // since small & cap cases are not together
  - pos = rune - 'A'
```

```bash
- trie.Val DoesNot Exist
```

```bash
- Use []*Trie         // With Position Calcs _OR_
- Use map[rune]*Trie
```

#### Source Code
```go
// --
// Trie Does Not Need a Val? Really?
// --
type Trie struct {
  Seen        int       // Freq of Visits
  Next        []*Trie
  IsEnd       bool      // Good To Have
}

func (t *Trie) Add(word string) {
  if word == "" {
    return
  }

  var tmp = t
  for idx, c := range word {
    var pos = c - 'A'
    
    // --
    // Start off With Next
    // As Per the Data Structure
    // Root Has No Role Here
    // --
    if tmp.Next[pos] == nil {
      tmp.Next[pos] = &Trie{Next: make([]*Trie, 128), Seen: 1,}
    } else {
      tmp.Next[pos].Seen += 1   // Freq of Visits In Prefix Sequence
    }
    if idx+1 == len(word) {
      tmp.Next[pos].IsEnd = true  
    }
    tmp = tmp.Next[pos]
  }
}

func (t *Trie) GetUniqPrefixFor(word string) string {
  if word == "" {
    return ""
  }
  
  if len(t.Next) == 0 {
    return ""
  }
  
  var prefix []rune
  var tmp = t

  for _, c := range word {
    prefix = append(prefix, c)
    var pos = c - 'A'
    tmp = tmp.Next[pos]   // Its Never The Root But Next
    
    // --
    // If A Word can be a Prefix of Another Word
    // Then Use Below Logic:
    //
    // if tmp == nil || tmp.Seen == 1 {..}
    // --
    if tmp == nil || tmp.Seen == 1 || tmp.IsEnd {
      break
    }
  }

  return string(prefix)
}

func UniqPrefixes(words []string) []string {
  var root = &Trie{Next: make([]*Trie, 128)}
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
  // [H hi her hel fella fello A a z Z]
  fmt.Printf(
    "%v\n", 
    UniqPrefixes(
      []string{"Hi", "hi", "her", "hello", "fella", "fellow", "Andy", "am", "ze", "Zero"}),
    ),
}
```

#### References
- https://www.techiedelight.com/shortest-unique-prefix/
