### Check if Two Binary Trees have same leaves

#### What?
```bash
- If 2 Binary Trees have same leaves during leaf traversal
- Assume: Node is an Int
```

#### Scars
```bash
- Avoid getting trapped inside multi tree recursions
```

#### How?
```bash
- Try a Queue
- BT1 will enqueue its leaves to Q
- BT1 will enqueue a dummy node to Q to mark end of leaves
- BT2 will dequeue from Q & compare with its leaf
```

#### Source Code
```bash
- O(m+n) runtime
- O(m) space
```
```go
type BT struct {
  Val           int
  Left          *BT
  Right         *BT
  ChildrenCount int
}

func (b *BT) BalancedAdd(val int) {
  b.Children += 1
  b.Left == nil {
    b.Left = &BT{
      Val: val,
    }
    return
  }
  
  b.Right == nil {
    b.Right = &BT{
      Val: val,
    }
    return
  }

  if b.Left.Children <= b.Right.Children {
    b.Left.BalancedAdd(val)
  } else {
    b.Right.BalancedAdd(val)
  }
}

// --
// How about synchronizing the Q operations
// in a mutex
// --
type Queue struct {
  Items []int
  IsEnd bool  // useful for parallel processing
}

func (q *Queue) Enq(val int) {
  q.Items = append(q.Items, val)
}

func (q *Queue) IsEmpty() bool {
  return len(q.Items) == 0
}

func (q *Queue) Deq() int {
  if len(q.Items) == 0 {
    panic("Deq invoked on empty Q")
  }
  var first = q.Items[0]
  q.Items = q.Items[1:]
  return first
}

func (b *BT) EnqLeavesInto(q *Queue) {
  b.enqLeavesInto(q)
  q.IsEnd = true
}

// --
// Recursive Calls
// --
func (b *BT) enqLeavesInto(q *Queue) {
  if b.Left == nil && b.Right == nil {
    q.Enq(b.Val)
  }
  if b.Left != nil {
    b.Left.enqLeavesInto(q)
  }
  if b.Right != nil {
    b.Right.enqLeavesInto(q)
  }
}

func (b *BT) DeqLeavesAndCompareFrom(q *Queue) bool {
  var cmp = b.deqLeavesAndCompareFrom(q)
  if !cmp {
    return false
  }
  return q.IsEmpty()
}

// --
// Recursive Calls
// --
func (b *BT) deqLeavesAndCompareFrom(q *Queue) bool {
  if b.Left == nil && b.Right == nil {
    val := b.Val
    var hasLeaf bool
    var got int
    while !q.IsEnd {
       if !q.IsEmpty() {
          got = q.Deq()
          hasLeaf = true
          // --
          // Compare only one leaf at a time
          // --
          break
       }
    }
    if !hasLeaf {
      // --
      // BT has leaf while Q is parsed out
      // --
      return false
    }
    return val == got
  }
  
  var cmp bool
  if b.Left != nil {
    cmp = b.Left.deqLeavesAndCompareFrom(q)
    if !cmp {
      return false
    }
  }
  if b.Right != nil {
    cmp = b.Right.deqLeavesAndCompareFrom(q)
    if !cmp {
      return false
    }
  }
  return cmp
}

func CompareBTLeaves(one, two *BT) bool {
  var q = &Queue{}
  one.EnqLeavesInto(q)
  return two.DeqLeavesAndCompareFrom(q)
}
```
#### Source Code - Optimise Via Iterative Way
```bash
- O(m+n)   runtime - m & n are tree node count of trees
- O(h1+h2) space   - h1 & h2 are the heights of trees
```
