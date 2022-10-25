## My Rust 101 from cargo-all-features 

### import in golang vs use in rust
```yaml
- Full import paths in golang are shortened via :: in rust
- src/types.rs file is declared as a mod in src/lib.rs
- types mod is used ( i.e. imported) as crate::types
```

### os/exec in golang is process::Command in rust
```rust
    let mut command = process::Command::new(&crate::cargo_cmd());
    command.arg("--no-default-features");
```

### Note the builder pattern's naming convention
```rust
    command.arg("--no-default-features");
```

### Match expressions Against Enum
```rust
command.arg(match cargo_command {
    CargoCommand::Build => "build",
    CargoCommand::Check => "check",
    CargoCommand::Test => "test",
});
```

### Iterate Vec<String> into a single comma separated string
```rust
let mut features = feature_set
    .iter()
    .fold(String::new(), |s, feature| s + feature + ",");
```

### String Parsing Goodies
```rust
if !features.is_empty() {
    features.remove(features.len() - 1);
    command.arg("--features");
    command.arg(&features);
}
```

### env with command args
```rust
for arg in env::args().skip(2) {
    command.arg(&arg);
}
```

### Thinking A Bit Even if its println!
```rust
println!("crate={} features=[{}]", self.crate_name, self.features); // [{}] is nothing special but improves readability
```

## References
```yaml
- https://github.com/frewsxcv/cargo-all-features
```