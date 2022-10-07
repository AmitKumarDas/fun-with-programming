## RUST for the starters
### Generic
- Rust avoids storing data on heap
- Complier does not implicitly store data on heap

### Owner
- Every value has a single owner that determines it lifetime
- When owner of some value is dropped the value is dropped as well
- When program leaves a block in which variable is declared, that variable will be dropped, dropping its value with it

### Vec
- Vector is dynamic (in size), allows pushing values to it at runtime
- Vector of string objects is interesting; do you notice the nested heap management

### String
- Rust holds String object on stack; where this obj consists of 1/ a pointer to data, 2/ length & 3/ capacity
- Since size of a String object is always fixed i.e. 3 words, its always stored on a stack
- String is a Vec<T> with guarantees of holding only well formed UTF-8 text
- A mutable String is capable of resizing its buffers via `.push_str()`

### string slice
- string slices are mostly references; hence they will always be &str
- There are two cases while working with string slices
- 1/ Create a reference to sub string
- 2/ Or we use string literals
- If "amit" is a &str i.e. string literal then where is the owner
- &str is special
- 1/ They are string slices that refer to a preallocated text
- 2/ Stored in read-only memory as part of the executable
- 3/ Memory that is shipped with the program & does not rely on buffers allocated in the heap
- 4/ A hard coded string defined in the binary of your program
- 5/ A string literal have static lifetime - it lasts as long as program is running

### to_string() vs. to_owned()
- https://users.rust-lang.org/t/to-string-vs-to-owned-for-string-literals/1441/5

### Useful Commands
- cargo new unittest-101 --lib

### References
- https://blog.thoughtram.io/ownership-in-rust/
- https://willcrichton.net/notes/rust-memory-safety/
- https://cooscoos.github.io/blog/stress-about-strings/
