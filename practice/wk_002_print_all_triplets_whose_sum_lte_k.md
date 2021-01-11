### Print ALL Triplets From Array Whose Sum Is Less Than Equal To K

#### Samples
```bash
- Input: [2, 4, 2, 5, 2, 7, 8]
- Input: 6
- Output: [2, 2, 2]
```

#### How
```bash
[2, 4, 2, 5, 2, 7, 8] k=6
[2, 2, 2, 4, 5, 7, 8]

[1, 4, 2, 5, 2, 7, 1, 8] k=10
[1, 1, 2, 2, 4, 5, 7, 8]

- sort the array // O(NlogN)
- for every element check all other pairs
```

#### Source Code
```go
func PrintTriplets(given []int, k int) {
  sort.Ints(given)
  size := len(given)
  
  for low:=0; low<size-3; low++ {
    var mid = low + 1
    var high = size - 1

    for mid < high {
      sum := arr[low] + arr[mid] + arr[high]
      if sum <= k {
        fmt.Printf("%d %d %d\n", arr[low], arr[mid], arr[high])
        mid++
      } else {
        high--
      }
    }
  }
}
```

