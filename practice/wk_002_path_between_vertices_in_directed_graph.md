### Find Path Between Given Vertices In A Directed Graph

#### Explain
```bash
- Given a Graph
- Given Two Vertices
- Determine if Dest Vertex is Reachable from Source Vertex
- Print the Path
```

#### How
```bash
        a
  e  f  b   d
        c   g
- a, c

- Path(src, dest G, path []int) []int
  - adjs := src.Adjs()
  - for _, a := range adjs
    - if a.Val == dest.Val
      - return append(path, a.Val)
    - else
      - newpath := Path(a, dest)
      - if newpath[len(newpath)-1] == dest.Val
        return newpath
```

#### Reference
- https://www.techiedelight.com/Category/Backtracking/
