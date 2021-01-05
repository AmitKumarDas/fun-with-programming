### Compare Leaves of 2 Binary Trees - II
- `O(1) space` `Leaves-As-LinkedList` `Multiple Trees` `Traversal`

#### Tips
```bash
- O(1) space Implies Recursion
```

#### Source Code
```go
func isLeaf(b *BT) bool {
  return b.Left == nil && b.Right == nil
}

func connectLeaves(b *BT, head, tail *BT) {
  if isLeaf(b) {
    if head == nil {
      head = b
      tail = b
    } else {
      tail.Right = b
      tail = b
    }
  }
  if b.Left != nil {
    connectLeaves(b.Left, head, tail)
  }
  if b.Right != nil {
    connectLeaves(b.Right, head, tail)
  }
}

func CompareLeavesOfTrees(a, b *BT) bool {
  ahead, atail := connectLeaves(a)
  bhead, btail := connectLeaves(b)
}
```
