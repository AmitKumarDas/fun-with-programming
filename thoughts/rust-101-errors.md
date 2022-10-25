## My Rust 101 on errors

### Enum, Variants of an Enum & Generics
```rust
// T can be any value
// E can be any error
enum Result<T, E> {
    Ok(T), // Ok is a variant of enum Result; globally available
    Err(E), // Err is a variant of enum Result; globally available
}
```

### Dealing with Errors
```rust
use std::io;
use std::fs::File;

fn read_username_from_file(path: &str) -> Result<String, io::Error> {
    let f = File::open(path); // straight to business

    let mut f = match f { // Shadowing is okay
        Ok(file) => file, // file has no collisons with module name
        Err(e) => return Err(e), // Error vs. Err
    };

    let mut s = String::new();
    match f.read_to_string(&mut s) { // Explicit &mut teaches your memory management
        Ok(_) => Ok(s),
        Err(err) => Err(err),
    } // No semicolon so return
}
```

```yaml
- Above code is not at all verbose
- It looks verbose since entire logic is explicit
- Overall LOC is smaller compared to Golang equivalent
```

### Unwrap makes your code terse but can panic
```rust
fn read_username_from_file(path: &str) -> Result<String, io::Error> {
    let mut f = File::open(path).unwrap();
    let mut s = String::new();
    f.read_to_string(&mut s).unwrap();
    Ok(s)
}
```

```yaml
- I might think of using above logic to replace bash scripts
- rust implements unwrap as follows
- rust also implements expect similar to unwrap
```

```rust
// result.rs

impl<T, E: fmt::Debug> Result<T, E> {
    pub fn unwrap(&self) -> T {
        match self {
            Ok(t) => t,
            Err(e) => unwrap_failed("called `Result::unwrap()` on an `Err` value", &e),
        }
    }
}
```

```yaml
- unwrap_failed is a shortcut to the panic! macro
- This means if you use .unwrap() then on error software crashes
- Can we implement a function that is similar to unwrap but returns the error?
```

### My own panic messages
```rust
fn read_username_from_file(path: &str) -> Result<String, io::Error> {
    let mut f = File::open(path).expect("error opening file");
    let mut s = String::new();
    f.read_to_string(&mut s).unwrap("error reading file to string");
    Ok(s) 
}
```

```rust
impl<T, E: fmt::Debug> Result<T, E> {
    pub fn expect(self, msg: &str) -> T {
        match self {
            Ok(t) => t,
            Err(e) => unwrap_failed(msg, &e),
        }
    }
}
```

### Error with fallback values
```rust
fn read_username_from_file(path: &str) -> Result<String, io::Error> {
    let mut f = File::open(path).expect("error opening file");
    let mut s = String::new();
    f.read_to_string(&mut s).unwrap_or("guest"); // eliminate templating
    Ok(s)
}
```
```yaml
- Above features in a lang might eliminate templating & pains associated with them
- There is also .unwrap_or_else which takes in a closure
```

### Error propagation i.e. write code for happy path here
```rust
fn read_username_from_file(path: &str) -> Result<String, io::Error> {
    let mut f = File::open(path)?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}
```

```rust
fn main() {
    match read_username_from_file("user.txt") {
        Ok(username) => println!("Welcome {}", username),
        Err(err) => eprintln!("Whoopsie! {}", err)
    };
}
```

```yaml
- Did you notice the question marks `?`
- Error propagation makes rust logic as terse as a bash & yet be correct at compile time
- Hey there is `eprintln!`
- [extreme][opinion] rust entry point full of matches while library full of happy paths
- [mnemonic] {} for expressions {} for println! {} for function blocks
```

### Is Error not an interface (aka error interface in golang)
```yaml
- We can not use ? for every function if they return different types of error
```

```rust
use std::error;

fn read_number_from_file(filename: &str) -> Result<u64, Box<dyn error::Error>> {
    let mut file = File::open(filename)?; // io::Error

    let mut buffer = String::new();
    file.read_to_string(&mut buffer)?; // io::Error

    let parsed: u64 = buffer.trim().parse()?; // ParseIntError
    Ok(parsed)
}
```

```yaml
- We tell rust that something that implements Error trait is coming along
- We dont know the type of Error trait at compile time
- Hence we make it a trait object! Wait what! An object of an interface
- dyn std::error::Error is a trait object
- We also know know the size of the error at compile time. Hence we wrap it in a Box
- Box is a smart pointer that points to data that will eventually be on heap
```

```yaml
- NOTE: SMART POINTER points to data that will EVENTUALLY be on HEAP
```

### Downcast to get the eventual error
```rust
fn main() {
    match read_number_from_file("number.txt") {
        Ok(v) => println!("Your number is {}", v),
        Err(err) => {
            if let Some(io_err) = err.downcast_ref::<std::io::Error>() {
                eprintln!("Error during IO! {}", io_err)
            } else if let Some(pars_err) = err.downcast_ref::<ParseIntError>() {
                eprintln!("Error during parsing {}", pars_err)
            }
        }
    };
}
```

```yaml
- Logic is so much explicit yet it remains terse
- `if let Some(my_fav_naming) =`
- Above is not a == but a single equal to
- Which implies `if something` then get into its block
```

## References
```yaml
- https://fettblog.eu/rust-error-handling/
```
