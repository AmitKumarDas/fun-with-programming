## 2 Minutes Guide To eBPF
- There will be 2 programs - bpf program & userspace program

- libbpf library is imported into both bpf & userspace program
- libbpf is used for 1/ loading a bpf program & 2/ interacting with bpf program

- bpf program uses headers defined in libbpf
- bpf program is compiled with clang to produce bpf object file

- userspace program uses libbpf to LOAD bpf object INTO kernel
- userspace program uses libbpf to interact with running bpf programs

- SEC("kprobe/sys_mmap")
- This is the section macro (defined in bpf_helpers.h)
- All programs must have one
- It tells libbpf what part of the compiled binary to place the program

- BPF_PROG_TYPE_KPROBE is one of the many bpf program types available
- we attach our bpf program to a program type

- Every different type of bpf program has its own ‘context’ 
- The context that you get access to for use in your bpf program

- struct pt_regs is a context
- It gives access to the virtual registers of the calling process
```c
SEC("kprobe/sys_mmap")
int kprobe__sys_mmap(struct pt_regs *ctx)
```

- A custom struct may be defined in common.h file
- This struct is included in bpf & userspace program

- A way to transmit output to userspace everytime the bpf program is called
- You may set up a ring buffer for this

## References
- https://blog.aquasec.com/libbpf-ebpf-programs
- https://github.com/aquasecurity/libbpfgo/blob/main/Readme.md
- https://blogs.oracle.com/linux/post/bpf-a-tour-of-program-types
