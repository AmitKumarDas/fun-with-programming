```yaml
- http://cliffle.com/p/dangerust/
```

### Cast an Integer to a Char In Rust
```yaml
- let character = 97 as char;
```


### Lint
```rust
// silence the opinions about naming conventions
#![allow(
    non_upper_case_globals,
    non_camel_case_types,
    non_snake_case,
)]
```

### Layout - Alignment
```rust
// Rust to lay out the struct exactly like C
// repr is applied to only types e.g. struct(s), variables
// repr can not be applied to values
// Without this, Rust does not care about the ordering of fields in a struct
// And it will optimize them for best packing and alignment
#[repr(C)]
struct body {
  position: [f64; 3],
  velocity: [f64; 3],
  mass: f64,
}
```

### #define to const
```yaml
- `#define` can be replaced by Rust `const`
```

### Array =~ Pointers in C
```yaml
- body bodies[] is same as body *bodies
- i.e. passing an array is same as passing a pointer
- _Note: body is the datatype while bodies is the name_
```

### C Pointer becomes *mut data_type in Rust
```yaml
- body *bodies becomes *mut body in Rust

- where body is a datatype
- bodies is an array of body items
```

### _Superpower_: *mut
```yaml
- Explicitly & Carefully Handled via unsafe
```

### Pointer i.e. Array Index is now
```yaml
- bodies[0] becomes *bodies.add(0) in Rust
```

### unsafe
```yaml
- unsafe fn offset_Momentum(bodies: *mut body) {...}

- unsafe is an explicit statement to the caller
- fn may do dumb things with pointers
- fn might corrupt the passed pointer
```

### In C - bodies[i] is same as *(bodies + i)
```yaml
- C ASSUMES that i is a valid index of the array

- a[0] is same as *a
- a[1] is same as *(a + 1)

- a[1] is NOT *a + 1
```

### *(a + i) in C becomes *a.add(i) in Rust
```yaml
- does not overload arithmetic operators
- Pointers provide add, sub & other operations - WOW!!!
- _Note: Easy to SPOT in CODE REVIEW_
```

### for m in 0..3 {...} in Rust
```yaml
- Resulting indexes are 0, 1 & 2
```

### In C double position_Delta[3] Leads to Unpredicatable Values
```yaml
- Reason for serious bugs in C
- However, sometimes leaving memory Un-Initialized is important for performance
```

### let mut position_Delta = [mem::MaybeUninit::<f64>::uninit(); 3];
```yaml
- Which is shorthand for:
  - let mut position_Delta: [mem::MaybeUninit<f64>; 3] = [mem::MaybeUninit::uninit(); 3];
- An array of three f64s containing arbitrary uninitialized memory
```

### std::mem::MaybeUninit
```yaml
- expressing storage locations i.e. pointers that might be uninitialized

- MaybeUninit<T> is a T that might not be initialized
- MaybeUninit::uninit() is an uninitialized value

- Accessing these variables is unsafe, but that shouldn’t scare you
- You need to be real sure that the variables are initialized before you read them
- The first thing the Rust code does, like C, is to write the array elements with valid data
```

### position_Delta[m].as_mut_ptr().write(value_goes_here);
```yaml
- MaybeUninit::as_mut_ptr()
  
- produces a mutable pointer to the possibly uninitialized memory
- write is another raw pointer operation - WOW !!!
- writes data into the pointed-to memory without reading it first
```

### Read Before Write
```yaml
- Why would writing to memory involve reading it first?
- In Rust, it’s usually because the type you’re storing has a **destructor**
- f64 doesn’t have any destructor. Why?
```

### std::mem::transmute - Why Memory Transmute?
```yaml
- converts between any two types as long as they’re the same size
- without running any code or changing any data
- it simply reinterprets the bits of one type as the other
```

### Convert A Floating Point Number Into A Pointer
```yaml
- or any three-word struct into any other, it is firmly, gloriously unsafe
- Transmute only if target is fully initialized else you get garbage
```

### Transmute in C - From Float to Int
```yaml
- float x = something();
- int y = *(int *) &x; // Why Do You Need to Transmute Via Address Route
```

### let position_Delta: [f64; 3] = mem::transmute(position_Delta);
```yaml
- Consuming using the same name
- This is common idiom in Rust to change 1/ type or 2/ mutability

- **Name Shadowing** may be scary
- You do not have to use it if you do not like
```

```rust
// Align16 is a struct that contains an array of ROUNDED_INTERACTIONS_COUNT f64s
// Aligns to 16 byte boundary
// Its a tuple struct with un-named fields
#[repr(align(16))]
#[derive(Copy, Clone)]
struct Align16([f64; ROUNDED_INTERACTIONS_COUNT]);
```

```rust
// static when within a function implies its initialized only once at 1st fn call
// Hence static mut is not threadsafe
// Hence advanced fn is tagged with unsafe
// Callers need to implement thread safety
unsafe fn advance(bodies: *mut body) {
  // ...

  // look how double array is initialized
  static mut position_Deltas: [Align16; 3] =
    [Align16([0.; ROUNDED_INTERACTIONS_COUNT]); 3];

  static mut magnitudes: Align16 =
    Align16([0.; ROUNDED_INTERACTIONS_COUNT]);
}
```

```yaml
- position_Deltas[m].0[k]
- .0 is accessing the first (& only) unnamed field in the tuple struct Align16
- where Align16 struct is an element of the arrary position_Deltas
- where Align16 is itself an array
```
