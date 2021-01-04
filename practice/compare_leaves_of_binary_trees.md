### Check if Two Binary Trees have same leaves

#### What?
```bash
- If 2 Binary Trees have same leaves during leaf traversal
- Assume: Node is an Int
```

#### How?
```bash
- Try a Queue
- BT1 will enqueue its leaves to Q
- BT1 will enqueue a dummy node to Q to mark end of leaves
- BT2 will dequeue from Q & compare with its leaf
```

#### Source Code
```go
type BT struct {
  Val   int
  Left  *BT
  Right *BT
}
```
