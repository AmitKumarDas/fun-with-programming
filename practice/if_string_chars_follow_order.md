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
  
  var patm = make(map[rune]bool, len(pat))
  for _, c := range pat {
    patm[c] = true
  }
  
  var occurs = map[rune]*Occurence{}
  for idx, t := range text {
    if patm[t] {
      o, ok := occurs[t]
      if !ok {
        occurs[t] = &Occurence{First: idx} // Set Only Once
      } else {
        o.Last = idx   // Keeps Updating When Found
        occurs[t] = o
      }
    }
  }
  
  var last int
  for idx, c := range pat {
    o := occurs[c]
    if idx == 0 {
      last = o.Last
      continue
    }
    if o.min < last {
      return false
    }
    last = o.Last // Reset last to current last
  }
  return true
}
```
