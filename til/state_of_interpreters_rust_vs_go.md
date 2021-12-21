### References
```yaml
- https://chr4.org/posts/2016-12-09-writing-an-interpreter-in-rust/
- https://chr4.org/posts/2016-12-17-writing-an-interpreter-in-rust-part-2/
```

### Const In Go
```go
const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  IDENT = "IDENT"
  INT   = "INT"
)
```

### Const In Rust
```rust
pub const ILLEGAL: TokenType = "ILLEGAL";
pub const EOF: TokenType = "EOF";

pub const IDENT: TokenType = "IDENT";
pub const INT: TokenType = "INT";
```

### Const - Idiomatic Rust via ENUM
```rust
#[derive(Debug, PartialEq)]
pub enum TokenType {
  Illegal,
  EndOfFile,

  Ident,
  Integer,
}
```

### Detecting An End Of File In Go
```go
// uses integer 0 (as a byte) to indicate when the end of file is reached
func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}
```
