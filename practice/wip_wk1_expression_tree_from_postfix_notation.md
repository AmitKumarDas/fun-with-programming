### Construct a Binary Expression Tree from Postfix Notation

#### What
```bash
Given:     a b + c d e + * *
Output:    Binary Expression Tree
Note:      Leaves are constants or variables
Note:      Non Leaves are the operators
```

#### How
```bash
- var root, left, right *BT
- If alphabet then push to Stack
- If operator then 
  - Create &BT{}
  - If Stack Not Empty Then Pop & Create Right Node or Left Node
  - If Stack Is Empty Then 
```
