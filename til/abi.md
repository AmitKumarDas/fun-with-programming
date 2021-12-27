#### References
```yaml
- https://kristoff.it/blog/maintain-it-with-zig/
```

### ABI - Application Binary Interface
```yaml
- An ABI defines how data structures or computational routines are accessed in machine code
- Is a low-level, hardware-dependent format

- In contrast, an API defines this access in source code
- Is a relatively high-level, hardware-independent, often human-readable format

- A common aspect of an ABI is the calling convention
- Which determines how data is provided as input to, or read as output from, computational routines
- Examples of this are the x86 calling conventions

- Adhering to an ABI (which may or may not be officially standardized)
- Is usually the job of a compiler, operating system, or library author
```

### ABI Characteristics
```yaml
- Covers details such as
- A processor instruction set:
    - register file structure, stack organization, memory access types
- Sizes, layouts, and alignments of basic data types that the processor can directly access
- Calling convention:
  - which controls how the arguments of functions are passed, and return values retrieved
  - E.g, it controls:
    - whether all parameters are passed on the stack, or some are passed in registers
    - which registers are used for which function parameters
    - and whether the first function parameter passed on the stack is pushed first or last
- How an application should make system calls to the operating system:
  - If the ABI specifies direct system calls rather than procedure calls
- In the case of a complete operating system ABI:
  - The binary format of object files, program libraries, and so on
```
