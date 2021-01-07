### Given The Root of BST Find Its Second Largest Node

#### Facts
```bash
- l <= root
- root < right
- in-order gives the full ordered list
- last but one from in-order provides the answer
```

#### How
```bash
- Get Root Node
- If No L & R then Error

- If R Recursion with R
- If No R Add Root As Max or Sec Max whichever is Nil
- If Either Max is Nil or Sec Max is Nil && L is not Nil Recursion with L
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

// --
// Recursive
// --
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

#### Source Code - Via Counter
```go
```
