### Compare Leaves of 2 Binary Trees - II
- `O(1) space` `Leaves-As-LinkedList` `Multiple Trees` `Traversal` 
- `H-T-Pointers` `LinkedList` `Recursion`

#### Tips
```bash
- If O(1) space & Tree Then Think Recursion
```

```bash
- If building a LinkedList To Parse Later
- Then Use `Head` & `Tail` Pointers

- Head Stays during LinkedList construction
- Tail Moves during LinkedList construction
- Head is Used While Parsing LinkedList
- Tail is Not Used While Parsing LinkedList
```

#### Source Code
```go
func isLeaf(b *BT) bool {
  return b.Left == nil && b.Right == nil
}

// --
// Move the Tail
// Stick the Head
// --
func connectLeaves(root *BT, head, tail *BT) {
  if root == nil {
    return
  }

  if isLeaf(root) {
    if head == nil {
      head = root
    } else {
      tail.Right = root
    }
    tail = root
  }
  
  // --
  // Additional Check is still better to avoid stack memory
  // --
  if root.Left != nil {
    connectLeaves(root.Left, head, tail)
  }
  
  // --
  // Additional Check is still better to avoid stack memory
  // --
  if root.Right != nil {
    connectLeaves(root.Right, head, tail)
  }
}


// --
// It is important to have head & tail
// Head Sticks
// Tail Moves
// --

// --
// Head is Used during Comparison / Parsing
// Tail is Not Used during Comparison / Parsing
// --
func CompareLeavesOfTrees(a, b *BT) bool {
  var ahead, atail *BT
  var bhead, btail *BT
  connectLeaves(a, ahead, btail)
  connectLeaves(b, bhead, btail)
  
  for ahead != nil && bhead != nil {
    if ahead.Val != bhead.Val {
      return false
    }
    ahead = ahead.Right
    bhead = bhead.Right
  }
  return ahead == nil && bhead == nil
}
```
