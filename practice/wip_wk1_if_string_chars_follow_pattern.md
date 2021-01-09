### Check If Chars of a String Follow The Order Specified in The Pattern

#### Samples
```bash
Given String:       aeroplane
Given Pattern:      ra
Output:             false
Assumption:         Pattern Has Distinct Chars

Given String:       aeroplane
Given Pattern:      el
Output:             false

Given String:       aeroplane
Given Pattern:      la
Output:             false

Given String:       aeroplane
Given Pattern:      ln
Output:             true
```

#### How
```bash
- Map pattern char with bool
- Store text chars' first & last occurence if they are found in pattern
- loop pattern 
  - if 1st char then set its max as max && continue
  - else if current min < last max then false
  - reset max to current max
```

#### Source Code
```go
type Occurence struct {
  First int
  Last  int
}

func MatchPatternOrder(text, pat string) bool {
  if text == "" {
    return pat == ""
  }
  if pat == "" {
    return true
  }

  var patm = make(map[rune]bool, len(pat))
  for _, c := range pat {
    patm[c] = true
  }
  
  var occurs = map[rune]*Occurence{}
  for idx, t := range text {
    if patm[t] {
      _, found := occurs[t]
      if !found {         // Found For The First Time
        occurs[t] = &Occurence{
          First: idx,     // Set Only Once
          Last: idx,      // Reqd For One Time Appearance Char(s)
        }
      } else {
        occurs[t].Last = idx  // Last Updates When Found Again
      }
    }
  }
  
  var last int
  
  // ---
  // Used As A Toggle
  // Can Not Use Idx == 0 As First Seen
  // Since Pat Char May Not Appear In Text
  // ---
  var firstSeen = true

  for _, c := range pat {
    o, found := occurs[c]
    if !found {
      continue
    }
    if firstSeen {
      last = o.Last
      firstSeen = false
      continue            // No Comparisons For First Pat Char
    }
    if o.First < last {
      return false
    }
    last = o.Last         // Reset to current last
  }
  return true
}
```

#### Test
```go
func main() {
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "la"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "lb"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "zx"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "ae"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "ea"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "el"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "ro"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "re"))
  fmt.Printf("%t\n", MatchPatternOrder("aeroplane", "er"))
}
```

#### Source Code - Repeated Removal

```bash
- Loop text & mutate by removing chars that are not in pattern
- Loop text & remove adjacent chars that are equal
- return text == pat
```

```go
```

#### References
- https://www.techiedelight.com/determine-string-follows-specified-order/
