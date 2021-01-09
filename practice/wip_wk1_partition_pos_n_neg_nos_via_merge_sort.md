### Segregate Positive & Negative Numbers using Merge Sort

#### Tags
- `stable` `merge sort` `divide n conquer` `basics` `partition` `learn`

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

#### Source Code - Corrections Needed
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
  fmt.Printf("%v\n", SegPosNeg([]int{10, 5, -9, 7, -20, -25, 20}))    // IMP
  fmt.Printf("%v\n", SegPosNeg([]int{10}))
  fmt.Printf("%v\n", SegPosNeg([]int{0, 0, 0}))
  fmt.Printf("%v\n", SegPosNeg([]int{10, 20}))
  fmt.Printf("%v\n", SegPosNeg([]int{10, -20}))
  fmt.Printf("%v\n", SegPosNeg([]int{-10, -20}))                      // IMP
  fmt.Printf("%v\n", SegPosNeg([]int{-10, -20, 0}))
  fmt.Printf("%v\n", SegPosNeg([]int{0, -10, -20, 0}))
  fmt.Printf("%v\n", SegPosNeg([]int{0, -10, -20, 1, 0, -1}))         // IMP
  fmt.Printf("%v\n", SegPosNeg([]int{9, -3, 5, -2, -8, -6, 1, 3}))    // !! Does Not Work !!
}
```

#### Source Code - Negative Then Positive
```go
// TODO
```

#### Source Code - Via MergeSort (Negative Then Positive)
```go
// --
// How
// 
// - Keep Negatives To Left
// - Keep 0s & Positives to Right

func SegNegPosByMergeSort(given []int) []int {
  var size = len(given)
  if size <= 1 {
    return given
  }
  
  return SegMergeSort(given)
}

func SegMergeSort(given []int) []int {
  if len(given) == 1 {
    return given
  }
  
  var size = len(given)
  var mid = int(size >> 1)      // Divide By 2
  var left = make([]int, mid)
  var right = make([]int, size-mid)
  
  for idx, item := range given {
    if idx < mid {
      // Fill left
      left[idx] = item
    } else {
      // Fill right
      right[idx-mid] = item
    }
  }

  sortedLeft = SegMergeSort(left)           // Recursion
  sortedRight = SegMergeSort(right)         // Recursion
  return SegMerge(sortedLeft, sortedRight)  // Merge Both Halves
}

func SegMerge(left, right []int) []int {
  var sizel = len(left)
  var sizer = len(right)
  
  if sizel == 0 {
    return right
  }
  if sizer == 0 {
    return left
  }
  
  var neg, pos []int

  for i:=0; i<sizel; i++ {
    if left[i] < 0 {
      neg = append(neg, left[i])
    } else {
      pos = append(pos, left[i])
    }
  }
  
  for i:=0; i<sizer; i++ {
    if right[i] < 0 {
      neg = append(neg, right[i])
    } else {
      pos = append(pos, right[i])
    }
  }
  
  for i:=0; i<len(pos); i++ {
    neg = append(neg, pos[i])
  }
  
  return neg
}
```

#### Source Code ~ MergeSort Via Partition ~ Negative Then Positive
```go
// --
// Partition Between Low & High 
// Partition Recursively 2 Times
// Merge at-last Within Partition Itself
// --

// --
// TIP:
//
// - Partition follows post-order traversal scheme
// - 1/ Recursively Partition Left Elements
// - 2/ Recursively Partition Right Elements
// - 3/ Operate on Left Results & Right Results
// --

// --
// NO RETURN OF ARRAY STYLE
// --

// --
// partition  With 4 args ~ orig aux, low high
// merge      With 5 args ~ orig aux, low mid high
// --

func merge(orig, aux []int, low, mid, high int) {
  var k = low
  
  // --
  // Left Negatives to Left
  // --
  for i:=low; i<=mid; i++ {
    if orig[i] < 0 {
      aux[k] = orig[i]
      k++
    }
  }
  
  // --
  // Right Negatives to Left
  // --
  for i:=mid+1; i<=high; i++ {
    if orig[i] < 0 {
      aux[k] = orig[i]
      k++
    }
  }
  
  // --
  // Left Positives to Right
  // --
  for i:=low; i<=mid; i++ {
    if orig[i] >= 0 {
      aux[k] = orig[i]
      k++
    }
  }
  
  // --
  // Right Positives to Right
  // --
  for i:=mid+1; i<=high; i++ {
    if orig[i] >= 0 {
      aux[k] = orig[i]
      k++
    }
  }
}

func partition(orig, aux []int, low, high int) {
  if high == low {
    return                           // LOGIC 1
  }
  
  var mid = low + ((high-low)>>1)    // x>>1 means x/2
  
  partition(orig, aux, low, mid)     // LEFT
  partition(orig, aux, mid+1, high)  // RIGHT
  
  merge(orig, aux, low, mid, high)   // LOGIC 2
}

func SegNegPosMergeSort(given []int) []int {
  var size = len(given)
  if size <= 1 {
    return given
  }
  
  var aux = make([]int, size)

  partition(given, aux, 0, size-1)
  return aux
}
```

#### References
- https://www.techiedelight.com/segregate-positive-negative-integers-using-mergesort/
