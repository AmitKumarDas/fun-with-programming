## 2 Minutes Guide To eBPF
- There will be 2 programs - bpf program & userspace program

- libbpf library is imported into both bpf & userspace program
- libbpf is used for 1/ LOADING a bpf program & 2/ INTERACTING with bpf program

- bpf program uses headers defined in libbpf
- bpf program is compiled with clang to produce bpf OBJECT file

- userspace program uses libbpf to LOAD bpf object INTO kernel
- userspace program uses libbpf to interact with running bpf programs

- SEC("kprobe/sys_mmap")
- This is the section macro (defined in bpf_helpers.h)
- All programs must have one
- It tells libbpf what PART of the COMPILED binary to PLACE the program

- BPF_PROG_TYPE_KPROBE is one of the many bpf program types available
- we ATTACH our bpf program to a program TYPE

- Every different type of bpf program has its own 'context' 
- The context that you get access to for use in your bpf program

- struct pt_regs is a CONTEXT
- It gives access to the virtual registers of the calling process
```c
SEC("kprobe/sys_mmap")
int kprobe__sys_mmap(struct pt_regs *ctx)
```

- A custom struct may be defined in common.h file
- This struct is included in bpf & userspace program

- A way to transmit output to userspace everytime the bpf program is called
- You may set up a ring buffer for this

- socket related program types
- SOCKET_FILTER, SK_SKB, SOCK_OPS

- socket creation
- `s = socket(AF_PACKET,SOCK_RAW,htons(ETH_P_ALL));`
- AF_PACKET is the domain
- SOCK_RAW is the type
- ETH_P_ALL is all the protocols

- BPF_PROG_TYPE_SOCKET_FILTER
- Dropping packets if prog returns 0
- Trimming packets if returned length is less than original

- We're NOT trimming or dropping the original packet 
- Which would still reach the intended socket intact
- We're working with a COPY of the packet metadata 
- which RAW SOCKETS can access for observability

- In addition to filtering packet flow to our socket
- We can do things that have side-effects
- e.g. collecting statistics in BPF maps

## References
- https://blog.aquasec.com/libbpf-ebpf-programs
- https://github.com/aquasecurity/libbpfgo/blob/main/Readme.md
- https://blogs.oracle.com/linux/post/bpf-a-tour-of-program-types
