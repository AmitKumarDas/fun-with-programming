## My rust 101 learnings

### Defining Custom Errors In Rust
```rust
#[derive(Debug)]
pub struct ParseArgumentsError(String);

impl std::error::Error for ParseArgumentsError {}

impl Display for ParseArgumentsError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.0)
    }
}
```

```yaml
- Above implements std::error::Error trait
- It can be implemented by a classic struct or tuple struct or unit struct
- What! You dont have to implement any methods of Error
- However derive the Debug trait
- We implement the Display trait since errors need to be printed somewhere
```

## References
```yaml
- https://fettblog.eu/rust-error-handling/
```
