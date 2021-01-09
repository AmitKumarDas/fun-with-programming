### Given A Circular Array, Find Next Greater Element For Each Item

#### Samples
```bash
Input:      [1, 5, 4, 3, 2, 6, 9]
Output:     [
             1 < 5 ~ 5, 
             5 < push push push 6 ~ 6,  
             4 < pop pop 6 ~ 6,  
             3 < push 6 ~ 6,
             2 < pop 6 ~ 6,
             6 < 9 ~ 9,
             9 < push push push push push ~ -1,
             ]
```

#### Snippets ~ Circular Array Loop
```go
var size = len(arr)

for i:=0; i<2*size; i=int((i+1)%size) {
}

// or

for i:=0; i<2*size; i++ {
  j = i%size
}
```
