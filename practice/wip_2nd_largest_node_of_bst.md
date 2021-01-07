### Given The Root of BST Find Its Second Largest Node

#### Facts
```bash
- l<=root
- root<right
- in-order gives the full ordered list
```

#### How
```bash
- Get Root Node
- If Left & Right Not Present then Error
- Traverse Right if Present
- Traverse Left if Right Not Present
- If Left & Right Not Present Then Set Max
- On Return If Max is Set Then Set 2nd Max

- Get Root Node
- If NO L & R then Error
- If R push to Stack
- If No R push Root to Stack
- If no R & L then pop twice & return

- Get Root Node
- If No L & R then Error

- If R Recursion with R
- If No R Add Root to Array
- If L Recursion with L
```

### Source Code
```go
func SecLarg(root *BST) int {
  if root == nil {
    panic("Root BST is Nil")
  }
  if root.Left == nil && root.Right == nil {
    panic("Single node BST")
  }
  
  var max, sec *int
  secLargest(root, max, sec)
  return *sec
}

func secLargest(root *BST, max, sec *int) {
  if root.Right != nil {
    secLargest(root.Right, max, sec)
  } 
  if max != nil && sec != nil {
    return
  }

  var m = root.Val
  if max == nil {
    max = &m
  } else {
    sec = &m
  }
  
  if max != nil && sec != nil {
    return
  }
  
  if root.Left != nil {
    secLargest(root.Left, max, sec)
  }
}
```
