## Why types are so powerful in Rust

### Your own String
```rust
/// A transparent wrapper around [`String`]
#[derive(Clone, Ord, PartialOrd, Eq, PartialEq, Hash, Debug)]
pub struct Feature(pub(crate) String);
```

### Think of your own List
```rust
/// A transparent wrapper around [`Vec<String>`]
#[derive(Default, Clone, Debug)]
pub struct FeatureList(pub(crate) Vec<Feature>);
```

## References
```yaml
- https://github.com/frewsxcv/cargo-all-features/blob/master/src/types.rs
```
