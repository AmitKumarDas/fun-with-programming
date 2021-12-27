### References
```yaml
- https://hacks.mozilla.org/2019/03/standardizing-wasi-a-webassembly-system-interface/
- https://hacks.mozilla.org/2017/02/creating-and-working-with-webassembly-modules/
- https://hacks.mozilla.org/2017/02/what-makes-webassembly-fast/
```

#### Few Implementations of WASI - Think of JVM flavours
```yaml
- wasmtime, Mozilla’s WebAssembly runtime:
  - https://github.com/CraneStation/wasmtime
- Lucet, Fastly’s WebAssembly runtime:
  - https://www.fastly.com/blog/announcing-lucet-fastly-native-webassembly-compiler-runtime
- a browser polyfill:
  - https://wasi.dev/polyfill/
```

#### WebAssembly Instructions
```yaml
- WebAssembly instructions are sometimes called virtual instructions
- They have a much more direct mapping to machine code than JavaScript source code
- They represent a sort of intersection of what can be done efficiently across common popular hardware
- But they aren’t direct mappings to the particular machine code of one specific hardware
```

#### C/C++/Rust >---> IR ---> WASM ---< x86/ARM
```yaml
- How to go from C to WASM:
  - Use clang frontend to go from C to LLVM intermediate representation
  - LLVM can perform optimisations
  - Some LLVM backend to build a WASM module
```

#### C to .wasm
```c
int add42(int num) {
  return num + 42;
}
```
##### corresponding .wasm file i.e. binary representation
```wasm
00 61 73 6D 0D 00 00 00 01 86 80 80 80 00 01 60
01 7F 01 7F 03 82 80 80 80 00 01 00 04 84 80 80
80 00 01 70 00 00 05 83 80 80 80 00 01 00 01 06
81 80 80 80 00 00 07 96 80 80 80 00 02 06 6D 65
6D 6F 72 79 02 00 09 5F 5A 35 61 64 64 34 32 69
00 00 0A 8D 80 80 80 00 01 87 80 80 80 00 00 20
00 41 2A 6A 0B
```

#### num + 42 looks like below
```yaml
- hexadecimal: 20 00 41 2A 6A
- binary: 00100000 00000000 01000001 00101010 01101010
- text:
  - get_local 0
  - i32.const 42
  - i32.add
```

#### How does above code work?
```yaml
- WebAssembly is an example of something called a stack machine
- All values an operation needs are queued up on the stack before the operation is performed
- Since add needs two, it will take two values from the top of the stack
- This means that the add instruction can be short (a single byte)
- Since the instruction doesn’t need to specify source or destination registers
- This reduces the size of the .wasm file, which means it takes less time to download
```

#### Stack Machine vs. Physical Machine
```yaml
- Even though WebAssembly is specified in terms of a stack machine
- That’s not how it works on the physical machine
- When the browser translates WebAssembly to the machine code for the machine the browser is running on:
  - It will use registers
- Since the WebAssembly code doesn’t specify registers, it gives the browser more flexibility to use the best register allocation for that machine
```

#### Parsing - JavaScript vs. WASM
```yaml
- Once it reaches the browser, JavaScript source gets parsed into an Abstract Syntax Tree
- Browsers often do this lazily:
  - Only parsing what they really need to at first
  - Just creating stubs for functions which haven’t been called yet
- Then AST is converted to an intermediate representation (called bytecode):
  - That is specific to that JS engine
- In contrast, WebAssembly doesn’t need to go through this transformation:
  - Since it is already an intermediate representation
  - It just needs to be decoded and validated to make sure there aren’t any errors in it
```

#### Compiling + Optimizing - JavaScript vs. WASM
```yaml
- JavaScript is compiled during the execution of the code
- Depending on what types are used at runtime:
  - multiple versions of the same code may need to be compiled

- WebAssembly starts off much closer to machine code
- E.g, the types are part of the program:
- This is faster for a few reasons.
- Compiler:
  - Doesn’t observe what types are being used
  - Doesn’t compile different versions of the same code based on the different types
- More optimizations have already been done ahead of time in LLVM
```

#### Garbage Collection
```yaml
- At least for now, WebAssembly does not support garbage collection at all
- Memory is managed manually (as it is in languages like C and C++)
- This makes programming more difficult for the developer
- It does also make performance more consistent
```

#### Shared memory concurrency
```yaml
- Speed up code by running in parallel
- This can sometimes backfire
- Since the overhead of communication between threads

- But if you can share memory between threads, it reduces this overhead
- To do this, WebAssembly will use JavaScript’s new SharedArrayBuffer
```

#### SIMD - Single Instruction, Multiple Data
```yaml
- If you read other posts or watch talks about WebAssembly, you may hear about SIMD support
- It’s another way of running things in parallel.

- SIMD makes it possible to take a large data structure, like a vector of different numbers
- Apply the same instruction to different parts at the same time
- So it can drastically speed up the kinds of complex computations you need for games or VR
```
