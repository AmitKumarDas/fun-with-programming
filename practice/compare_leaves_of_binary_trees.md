### Check if Two Binary Trees have same leaves

#### What?
```bash
- If 2 Binary Trees have same leaves during leaf traversal
- Assume: Node is an Int
```

#### Source Code
```go
type BT struct {
  Val   int
  Left  *BT
  Right *BT
}

func (b *BT) Leaves(arr []int) []int {
  if b.Left == nil && b.Right == nil {
    return append(arr, b.Val)
  }

  if b.Left != nil {
    arr = append(arr, b.Left.Leaves(arr))
  }
  if b.Right != nil {
    arr = append(arr, b.Right.Leaves(arr))
  }
  return arr
}

func (b *BT) CompareLeaves(other *BT) bool {

}
```
