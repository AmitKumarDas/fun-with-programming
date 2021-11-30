### Bitmap index
```yaml
- database index uses bitmaps uses bit arrays

- work well for low-cardinality columns
- which have a modest number of distinct values
- either absolutely, or relative to the number of records that contain the data
- lowest cardinality is Boolean datatype i.e. 'True' and 'False'

- Bitmap indexes use bit arrays
- & answer queries by performing bitwise logical operations on them
- Bitmap indexes have a significant space and performance advantage

- Drawback: less efficient than the B-tree indexes for columns for frequent updates
- unsuitable for online transaction processing applications

- More often employed in read-only systems
- specialized for fast query
- e.g., data warehouses
```

### In-memory Bitmap
```yaml
- Intermediate results produced from bitmaps are also bitmaps
- Can be efficiently reused in further operations to answer more complex queries
- Many programming languages support this, e.g. Java has BitSet class
```

### Bitmap - Speed up DB query
```
- A temporary in-memory bitmap is created with ONE BIT FOR EACH ROW in the table
- 1 MB can thus store over 8 million entries i.e. 80,00,000

- Next, results from each index are combined into the bitmap using bitwise operations
- After all conditions are evaluated
- the bitmap contains a "1" for rows that matched the expression
- Finally, the bitmap is traversed and matching rows are retrieved

- In addition to efficiently combining indexes
- this also improves locality of reference of table accesses
- because all rows are fetched sequentially from the main table
- The internal bitmap is discarded after the query

- If there are too many rows in the table to use 1 bit per row
- a "lossy" bitmap is created instead, with a ONE BIT PER DISK PAGE
- In this case, the bitmap is just used to determine which pages to fetch
- the filter criteria are then applied to all rows in matching pages
```
