
#### Binary Instruction For Stack Based Virtual Machine
```yaml
- A general technology for running compiled code across multiple platforms
- The Wasm stack machine is designed to be encoded in a size- and load-time-efficient binary format
- WebAssembly aims to execute at native speed
- Take advantage of common hardware capabilities available on a wide range of platforms
```

#### Binary Code / Low Level
```yaml
- Wasm a very interesting prospect for running binary code in the browser
- Wasm assembly format is low-level and closer to the abstraction level of machine code
- This leads to improved startup time and generally a more predictable efficiency
```

#### Open and Debuggable
```yaml
- WebAssembly is designed to be pretty-printed in a textual format
- Debug, Test, Experiment, Optimize, Learn, Teach, and write programs by hand
- The textual format will be used when viewing the source of Wasm modules on the web
```

#### Wasm as a compilation target for Dart - Garbage Collection
```yaml
- Wasm's original form doesn’t work well for languages with garbage collection
- Wasm lacks built-in garbage collection support
- So Dart must include a GC implementation into the compiled Wasm module
- Including a GC implementation would be highly complex:
  - Will inflate the size of the compiled Wasm code
  - Hurt startup time
  - Wouldn’t lend itself well to object-level interop with the rest of the browser system
```
#### Wasm GC
```yaml
- Exploring the possibility of expanding Wasm:
  - with direct and performant support for garbage collected languages
```

#### Interoperability & Cross Platform
```yaml
- Interoperable with existing code written in different languages
- FFI is not the answer
- E.g. C code is platform specific
- Needs distribution of multiple binary modules (one for each platform)
- How about a single Wasm Binary Assembly Format:
  - reusable across all platforms
  - So compile your code to Wasm binary format only
  - Run it everywhere
```
