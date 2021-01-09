### Practices

#### Code Scars
```bash
- Tree Tips

- O(1) space ~ Recursion 
- O(1) space ~ DFS Traversal
```

```bash
- Tree Tips

- No Space Constraint
  - Queue i.e. ~ BFS traversal _OR_
  - Stack i.e. ~ BFS traversal
```

```bash
- Tree Tips for Leaves Specific Logic

- Multiple Stacks & ~ BFS                        - O(H)
- Connect Leaves ~ LinkedList & ~ DFS Recursion  - O(1)
```

```bash
- Depth Traversal Without Recursion
- Use Stack - O(H) space
```

```bash 
- Tree & Traversal via Stack (i.e. LIFO)

- pre-order traversal
  - Operate on Root
  - Push Right 
  - Push Left

- in-order traversal
  - Push Right
  - Operate on Root 
  - Push Left

- post-order traversal
  - Push Right
  - Push Left
  - Operate on Root
```

```bash
- Multiple Trees Tips

- If Fail Fast Comparisons
- I.E Fail Fast Instead of Collecting All Nodes & Then Failing

- Then Iterative Approach
- I.E. GetNext(Tree1) & GetNext(Tree2)
- GetNext() Requires Stacks or Queues i.e. data structure
- Since GetNext() on multiple trees
- Hence multiple Stacks or Queues are required

- Stack or Queue Will Help Provide the Next Required Element
- Stack / Queue will hold all the Nodes traversed so far
- Stack ~ Depth Traversal   ~ O(H) ~ Max Height
- Queue ~ Breadth Traversal ~ O(L) ~ Max Level at any height
```

```bash
- Recursion Does its Magic
  - When exit conditions are well defined
- Do not copy the Recursion code blocks
  - Into DP code block AS-IS
```

```bash
- DP & Max / Min - I
  - Major logic should be inside main dp filling block
  - Mostly if-elif-else blocks
  - Minimize Initial DP fill logic

- DP & Max / Min - II
  - Think Where the Max / Min Value will Percolate in DP Table
  - Will it be Top Right Cell of DP Table?
  - Will it be DP Table Edges?
```

```bash
- DP Table & Palindrome
  - Bottom Part of Diagonal is Useless
  - Fill Top Part of Diagonal only
```

```bash
- Do Not use Left, Right Pointer Movements in DP problems
- DP Is About Constructing the Equation
- DP Equation Is Composed Of Other Equations
```

#### DP
- Longest Palindromic Substring 
  - `dp` `diagonal` `loop` `equation`
- Longest Palindromic SubSequence 
  - `dp` `diagonal` `loop` `equation`

#### Array
- [GBRRBRG]_to_[RRRGGBB] i.e. Dutch Flag Problem 
  - `3-pointer` `rune` `swap` `o(n)` `o(1)`
- Dutch Flag with 4 Chars 
  - `4-pointer` `rune` `swap` `o(n)` `o(1)`
- Next Greater Elem In Array 
  - `stack`
  - `Next Great EATER elMENt In tRRAY is sTACKo`
- Next Greater Elem In Circular Array
  - `stack`
  - `CirCle Twice And Use mOOdulOO`
  - `for i:=0; i<2*n; i++ {arr[i%n]}`
- 
