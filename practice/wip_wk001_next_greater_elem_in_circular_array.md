### Given A Circular Array, Find Next Greater Element For Each Item

#### Samples
```bash
Input:      [1, 5, 4, 3, 2, 6, 9]
Output:     [
              push 1's idx,
              
              stack top < arr[i]          // 1 < 5
                result[stack top] = arr[i]
                stack pop
              push 5's idx
              
              stack top < arr[i]          // 5 < 4
                result[stack top] = arr[i]
                stack pop
              push 4's idx
              
              ...
            ]
```

#### Snippets ~ Circular Array Loop
```go
var n = len(arr)

for i:=0; i<2*n; i=int((i+1)%n) {
}

// _or_

for i:=0; i<2*n; i++ {
  j := i%n
  // use j
}
```
