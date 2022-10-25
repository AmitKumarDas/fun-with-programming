### C primitive types
```yaml
- Zig provides special c_ prefixed types for conforming to the C ABI
- These do not have fixed sizes, but rather change in size depending on the ABI being used
- NOTE: C ABI used is dependant on the target which you are compiling for 
- NOTE: e.g. CPU architecture, operating system
```

### void type
```yaml
- Note: C’s void (and Zig’s c_void) has an unknown non-zero size
- Zig’s void is a true zero-sized type
```

### calling conventions
```yaml
- How functions are called
- How arguments are supplied (in registries or on the stack & how)
- How return value is received
```

### References
```yaml
- https://ziglearn.org/chapter-4/
```
