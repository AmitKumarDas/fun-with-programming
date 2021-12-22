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

### Setting EOF In Go
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

### Idiomatic Go to Find EOF
```go
switch l.ch {
case '=':
  // [...]
case 0: // char can be an integer i.e. a byte
  tok.Literal = ""
  tok.Type = token.EOF
default:
  // [...]
}
```

#### Idiomatic Rust to Find EOF

```yaml
- Use an Option<char> and return None in case the end of the file is reached
- May be more idiomatic to return a Result with an EOF error
```

```rust
match self.ch {
  Some(ch @ '=') => {
  [...]

  // Handle EOF
  None => {
    tok.literal = String::new();
    tok.token_type = token::TokenType::EndOfFile;
  }
}
```

#### Main
```rust
use std::io::{self, BufRead, Write}; // How Is BufRead used?

pub mod lexer;
pub mod token;

fn main() {
  let stdin = io::stdin();

  loop { // For Ever Loop
    // Stdout needs to be flushed, due to missing newline
    print!(">> ");
    io::stdout().flush().expect("Error flushing stdout");

    let mut line = String::new(); // String CapitalCased?
    stdin.lock().read_line(&mut line).expect("Error reading from stdin");
    let mut lexer = lexer::Lexer::new(&mut line); // CapitalCased Lexer ?

    loop { // For Ever Loop
      let tok = lexer.next_token();
      println!("{:?}", tok);
      if tok.token_type == token::TokenType::EndOfFile { // CapitalCased TokenType?
          break;
      }
    }
  }
}
```

#### Source Code => Lexer => Tokens
```sh
$ cargo run
    Finished debug [unoptimized + debuginfo] target(s) in 0.0 secs
     Running `target/debug/writing_an_interpreter_in_rust`

>> let add = fn(x, y) { x + y; };
Token { token_type: Let, literal: "let" }
Token { token_type: Ident, literal: "add" }
Token { token_type: Assign, literal: "=" }
Token { token_type: Function, literal: "fn" }
Token { token_type: LeftParenthesis, literal: "(" }
Token { token_type: Ident, literal: "x" }
Token { token_type: Comma, literal: "," }
Token { token_type: Ident, literal: "y" }
Token { token_type: RightParenthesis, literal: ")" }
Token { token_type: LeftBrace, literal: "{" }
Token { token_type: Ident, literal: "x" }
Token { token_type: Plus, literal: "+" }
Token { token_type: Ident, literal: "y" }
Token { token_type: Semicolon, literal: ";" }
Token { token_type: RightBrace, literal: "}" }
Token { token_type: Semicolon, literal: ";" }
Token { token_type: EndOfFile, literal: "" }
>>
```
