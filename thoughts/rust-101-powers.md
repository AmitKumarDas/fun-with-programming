## Rust Quirks/Powers/Syntax Coming from Golang

### {}: Nested {}, Comma & Semicolon while import is possible/needed
```rust
use std::{
    convert::{AsMut, AsRef},
    iter::FromIterator,
    ops::{Deref, DerefMut},
};
```

### Ok(some_logic)
```rust
pub fn run(&mut self) -> Result<crate::TestOutcome, Box<dyn error::Error>> {
    //...
    Ok(if output.status.success() {
        crate::TestOutcome::Pass
    } else {
        crate::TestOutcome::Fail(output.status)
    })
}
```
