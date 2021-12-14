## Basics


## Dynamic Programming

### 0-1 Knapsack
#### Problem
```yaml
- each item has a wt. & value
- total weight <= limit
- max value possible with above constraint
```

#### Tips
```yaml
- optimal substructure
- overlapping subproblems
- dynamic programming
- memoize
- problems that puzzles you often are solved via dynamic programming
```

#### Signature
```yaml
- knapsack(val[], wt[], n, W) max:
  - where n is count & W is capacity
```

#### Snippets
##### Basic
```yaml
- knapsack(v[], w[], n, W):
  - if cap < 0 return INT_MIN # so as to compare
  - if n < 0 || W == 0: # no items left
    - return 0
  - include = v[n] + knapsack(v, w, n-1, W - wt[n])
  - exclude = knapsack(v, w, n-1, W)
  - return max(include, exclude)

- main:
  - W = 100
  - n = len(v) # int n = sizeof(v) / sizeof(v[0]); c++
  - knapsack(v, w, n-1, W) # v & w are arrays
```

##### Memoized / Top Down
```yaml
- knapsack(v, w, n, W, lookup): # O(n.W) time complexity
  - ...
  - k = string(n) + "-" + string(W) # pure forms of n & W
  - if _, found := lookup[k]; !found:
    - ...
    - lookup[k] = max(include, exclude) # O(n.W) space
  - return lookup[k]
```

##### Bottom Up
```yaml
- solve smaller then larger
- same runtime complexity as Memoized i.e. O(n.W)
- without recursion overhead
```
```yaml
- knapsack(v, w, n, W) max:
  - # V[i][j] # stores max value that is possible
  - # for wt <= j using first i items
  - #
  - # i-1 in V[][] represents previous / memoized value
  - # i-1 in w[] & v[] represent the current index since index ends at n-1
  - int V[n+1][W+1]
  - for j=0; j<W; j++:
    - V[0][j] = 0
  - for i=1; i<=n; i++:
    - for j=0; j<=W; j++:
      - excluded = V[i-1][j] # previous calculated value that must be <= j
      - if w[i-1] > j:
        - V[i][j] = excluded
      - else:
        - iWt = w[i-1] # w[i-1] is item i's weight
        - iVal = v[i-1] # v[i-1] is item i's value
        - V[i][j] = max(excluded, V[i-1][j-iWt] + iVal)
  - return V[n][W]
```

## Sorting

### Quick Sort
```yaml
- https://www.techiedelight.com/quicksort/
```

#### Signature
```yaml
- quicksort(arr, start, end) # start & end are indexes
- pivot(arr, start, end) # also known as partition # Lomuto partition
```

#### Snippets
```yaml
- quicksort(arr, start, end):
  - if start >= end return # stop based on indexes
  - p = pivot(arr, start, end)
  - quicksort(arr, start, p - 1) # is start redundant # tip: its swapped in pivot
  - quicksort(arr, p + 1, end) # is end redundant # tip: its swapped in pivot
```
```yaml
- pivot(arr, start, end):
  - pv = arr[end] # end value remains at its position # assumed as pivot
  - pi = start # start pivot index
  - for i = start; i < end; i++:
    - if arr[i] <= pv:
      - swap(arr, i, pi) # small values left of pivot
      - pi++
  - swap (arr, pi, end) # end val i.e. the pivot is now set to correct position
  - return pi # return the pivoted index
```

#### Mnemonics
```yaml
- pivot:
  - pi # pivot index ~ assume start index
  - pv # pivot value ~ assume end value
- quicksort:
  - quickly divide & rule both halves & keep doing this
  - di ~ pi = pivot(arr, start, end)
  - rule left half ~ quicksort(arr, start, di - 1)
  - rule right half ~ quicksort(arr, di + 1, end)
```

### Merge Sort

```yaml
- https://www.techiedelight.com/merge-sort/
```

#### Signature
```c
// merge using two arrays & indices
void merge(int arr[], int aux[], int low, int mid, int high)

// calls itself twice before merging the results
// - split left
// - split right
// - merge
void mergesplit(int arr[], int aux[], int low, int high)

// main()
mergesplit(arr, aux, 0, N - 1);
```

#### Snippets
```yaml
mergesplit:
  - if (high == low) { return }
  - int mid = (low + ((high - low) >> 1))
  - mergesplit(arr, aux, low, mid);          // split left half
  - mergesplit(arr, aux, mid + 1, high);     // split right half
  - merge(arr, aux, low, mid, high);         // merge the two halves

merge:
  - int k = low, i = low, j = mid + 1;
  - while (i <= mid && j <= high)
```

### Selection Sort
```yaml
- https://www.techiedelight.com/selection-sort-iterative-recursive/
```

#### Snippets
```yaml
- ssort
  - for (int i = 0; i < n - 1; i++):
    - min = i
    - for (int j = i+1; j < n; j++):
      - extract the index of min value
    - swap min & i indices
```

#### Mnemonics
```yaml
- ssort - ffor
- select index with min value
- swap indices
```

## Search

### Binary Search
```yaml
- https://www.techiedelight.com/binary-search/
- determine the index
- mnemonic:
  - binary - 0 1
  - find the right binary balance between low & high
  - return the balance i.e. mid
  - either tune high or tune low
```

```yaml
- int bsearch(int nums[], int n, int target)
  - low = 0
  - high = n - 1
  - while (low <= high): # loop
    - int mid = (low + high)/2
    - if target == arr[mid]:
      - return mid # found it! yay!
    - else if target < arr[mid]:
      - high = mid - 1
    - else:
      - low = mid + 1
  - return -1 # not found
```
