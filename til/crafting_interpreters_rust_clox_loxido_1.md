
#### References
```yaml
- https://ceronman.com/2021/07/22/my-experience-crafting-an-interpreter-with-rust/
- https://github.com/ceronman/loxido
```

#### Borrow Checker
```yaml
- Is a well-known struggle for Rust beginners
- Read The book "Learn Rust With Entirely Too Many Linked Lists"
```

#### Garbage Collection
```yaml
- Deal with graph-like structures with cycles is to use vector indices as some sort of pointer
- Study popular crates such as:
  - http://id-arena/
  - https://crates.io/crates/typed-arena
  - https://crates.io/crates/generational-arena
```

#### Trait Copy
```yaml
- https://doc.rust-lang.org/std/marker/trait.Copy.html
```

#### Profiling
```yaml
- https://perf.wiki.kernel.org/index.php/Main_Page
```

#### Deref
```yaml
- https://doc.rust-lang.org/std/ops/trait.Deref.html
```

#### Hashing Algorithm
```yaml
- SipHash
- FNV
- aHash
- fxHash
```

#### HashMap
```yaml
- https://github.com/rust-lang/hashbrown
```

#### vs. clox HashMap
```yaml
- HashBrown, which is a general-purpose hash table implementation
- Hash table in clox is tailored for a very specific use case
- Clox‘s Table only uses Lox strings as keys and Lox values as values
- Lox strings are immutable, interned, and they can cache their own hashes
- This means that hashing only occurs once ever for each string
- Comparison is as fast as comparing two pointers
- Additionally, not dealing with generic types has the advantage of ignoring type system edge cases such as Zero Sized Data Types.
```

#### Price of Dynamic Dispatch
```yaml
- I used trait objects to keep track of anything that should be garbage collected
- I copied this idea from crates such as gc and gc_arena
- Trait allowed to keep a list of all the objects necessary for sweeping process
- Trait also would contain methods for tracing an object
- Each object type had different ways of tracing
  - E.g, a string should only mark itself
  - An object instance should mark itself and its fields
  - Each GC-ed type should implement the GcTrace trait with the specific logic to trace it
  - Then, thanks to polymorphism, the tracing part of the GC was as simple as this
```
```rust
fn trace_references(&mut self) {
  while let Some(object) = self.grey_stack.pop() {
    object.trace(self);
  }
}
```
```yaml
- As usual with polymorphism, this was short and elegant
- However, there was a problem
- This kind of polymorphism uses dynamic dispatch, which has a cost
- In this particular case, the compiler is unable to inline the tracing functions
- So every single trace of an object is a function call
- Again, this is usually not a big problem, but when you have millions of traces per second, it shows
- In comparison, clox was simply using a switch statement
- This is less flexible but the compiler inlines it tightly which makes it really fast
```
```yaml
- So instead of using a trait object, rewrote the GC to use an enum instead
- Then wrote a match expression that would do the right tracing logic for each type
- This is a bit less flexible
- It also wastes memory because the enum effectively makes all objects use the same space
- But it improved tracing speed considerably with up to 28% less time for the most problematic benchmark
```
```yaml
- clox uses a different approach called Struct Inheritance
- There is a struct that acts as a header and contains meta-information about each object
- Then each struct that represents an object managed by the GC has this header as its first field
- Using type punning, it’s possible to cast a pointer to the header to a specific object and vice versa
- This is possible because in C structs are laid out in the same order as defined in the source
- Rust by default doesn’t guarantee the order of the data layout
- There are ways to tell the Rust compiler to use the same data layout as C:
  - which is used for compatibility, but I wanted to stay with Rust as intended
```

#### Small unsafety speedups
```yaml
- There were other small details that made clox faster and were related to avoiding safety checks
- E.g, the stack in clox is a fixed-size array and a pointer that indicates the top of the stack
- No underflow or overflow checks are done at all
- In contrast, Vec in Rust have push() and pop() operations which are checked
- I rewrote the stack code like clox and I was able to shave up to 25% of the running time
```
```yaml
- Program counter in clox is a C pointer to the next instruction in the bytecode array
- Increasing the PC is done with pointer arithmetic and no checks
- In Rust implementation, PC was an integer with the index in the position in the bytecode chunk vector
- Changing this to work as clox allowed me to shave up to 11% from the running time
```
```yaml
- Getting speed-up improvements is nice, but the downside is that the code becomes unsafe
- But at least it’s nice that you can be very specific about what parts of your code are unsafe
- This makes it much easier to debug when things go wrong.
```
```yaml
- How can clox live with all these unchecked pointer arithmetic?
- Well, the VM assumes that the bytecode will always be correct
- The VM is not supposed to take any arbitrary bytecode, but only bytecode generated by the compiler
- It’s the compiler’s responsibility to produce correct bytecode:
  - That would not cause underflows or overflows of the stack or frames
- While the clox compiler does prevent a lot of these cases, it does not prevent stack overflows
- The author intentionally decided to skip this check because it would require a lot of boilerplate code
- Real-world interpreters using unsafe structures like these must absolutely do it though
```
