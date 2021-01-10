### Given An Array Find Next Greater Elem of Each Array Item

#### Assumption
```bash
- When Next Greater Elem Is Not Found Then Set -1
```

#### Samples
```bash
Given:      [1, 4, 2,              1,  6,  2, 4]
Logic:      [
              push 1's idx,
              
              stack top < arr[i]  // 1 < 4
                result[1's idx] = 4 & pop
              push 4's idx,
              
              stack top < arr[i]  // 4 < 2
                --
              push 2's idx,
              
              stack top < arr[i]  // 2 < 1
                --
              push 1's idx,
              
              stack top < arr[i]  // 1 < 6
                result[1's idx] = 6 & pop
                result[2's idx] = 6 & pop
                result[4's idx] = 6 & pop
              push 6's idx,
              
              stack top < arr[i]  // 6 < 2
                --
              push 2's idx,
              
              stack top < arr[i]  // 2 < 4
                result[2's idx] = 4 & pop
              push 4's idx,
            ]
```

