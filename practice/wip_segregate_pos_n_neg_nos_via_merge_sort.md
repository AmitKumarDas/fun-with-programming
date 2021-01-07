### Segregate Positive & Negative Numbers using Merge Sort

#### Tags
- `stable order` `sort` `merge sort`

#### What
```bash
- Given an array of Positive & Negative integers
- Segregate them without changing the relative order of elements
```

#### Sample
```bash
Input:  10 5 -9 7 -20 -25 20
Output: 10 5 7 20 -9 -20 -25

Input:  10 5 6 -9 -2 -3 7 4 2 -1
Output: 10 5 6 7 4 2 -9 -2 -3 -1
```

#### How
```bash
- Loop All Items
- If pos item continue
- If neg & not last then swap with next pos item or last if no pos found
```

#### Source Code
```go
func SegPosNeg(given []int) []int {
  var size = len(given)
  if size <= 1 {
    return given
  }
  
  var isLastElemSwapped bool
  for idx, i := range given {
    // --
    // Positives & Zeros are Kept Left
    // --
    if i >= 0 {
      continue
    }
    if idx+1 == size {
      break
    }

    var tIdx = idx+1
    var next = given[tIdx]
    
    // --
    // Loop Further 
    // Till We Get A Positive or Last Val
    // --
    for next < 0 {
      if tIdx+1 == size {
        break
      }
      tIdx++
      next = given[tIdx]
    }
    
    // --
    // next can be Positive or Last Elem
    //
    // If Last Elem is Negative Then All Items
    // from Current Index till Last are Negative
    //
    // Logic Should Exit For Loop Only If 
    // Last Element Was Never Swapped Previously
    // This Maintains the Order of Negative Values
    // --
    if next < 0 && !isLastElemSwapped {
      break     // Order Needs to be Preserved
    }

    if tIdx+1 == size {
      // --
      // At this point 
      // Last Elem Which Is Positive Is Being Swapped
      // Hence Negatives That Appear Before This Last Pos
      // Should Move Past This SWAP Later
      // --
      isLastElemSwapped = true
    }

    // --
    // SWAP
    // --
    given[idx], given[tIdx] = given[tIdx], given[idx]
  }

  return given
}
```

#### Test
```go
func main() {
  fmt.Printf("%v\n", SegPosNeg([]int{10, 5, 6, -9, -2, -3, 7, 4, 2, -1}))
  fmt.Printf("%v\n", SegPosNeg([]int{10, 5, -9, 7, -20, -25, 20}))   // IMP
  fmt.Printf("%v\n", SegPosNeg([]int{10}))
  fmt.Printf("%v\n", SegPosNeg([]int{10, 20}))
  fmt.Printf("%v\n", SegPosNeg([]int{10, -20}))
  fmt.Printf("%v\n", SegPosNeg([]int{-10, -20}))  // IMP
  fmt.Printf("%v\n", SegPosNeg([]int{-10, -20, 0}))
  fmt.Printf("%v\n", SegPosNeg([]int{0, -10, -20, 0}))
  fmt.Printf("%v\n", SegPosNeg([]int{0, -10, -20, 1, 0, -1})) // IMP
}
```
