## README Driven Learning From `libbpf-bootstrap`

- GIVEN: `git clone git@github.com:libbpf/libbpf-bootstrap.git`
- THEN: `cd examples/c`
- THEN: `make minimal`
```shell
make minimal
  MKDIR    .output
  MKDIR    .output/libbpf
  LIB      libbpf.a
make: *** /Users/amitd2/work/libbpf-bootstrap/libbpf/src: No such file or directory.  Stop.
make: *** [/Users/amitd2/work/libbpf-bootstrap/examples/c/.output/libbpf.a] Error 2
```
- THEN: `cd ../..`
- THEN: `git submodule update --init --recursive`
```shell
Submodule 'blazesym' (https://github.com/libbpf/blazesym.git) registered for path 'blazesym'
Submodule 'bpftool' (https://github.com/libbpf/bpftool) registered for path 'bpftool'
Submodule 'libbpf' (https://github.com/libbpf/libbpf.git) registered for path 'libbpf'
Cloning into '/Users/amitd2/work/libbpf-bootstrap/blazesym'...
Cloning into '/Users/amitd2/work/libbpf-bootstrap/bpftool'...
Cloning into '/Users/amitd2/work/libbpf-bootstrap/libbpf'...
Submodule path 'blazesym': checked out '885b4d5297c1633444317d439b312626920b5486'
Submodule path 'bpftool': checked out '2d7bba1e8c17dd0422879c856cda66723b209952'
Submodule 'libbpf' (https://github.com/libbpf/libbpf.git) registered for path 'bpftool/libbpf'
Cloning into '/Users/amitd2/work/libbpf-bootstrap/bpftool/libbpf'...
Submodule path 'bpftool/libbpf': checked out 'b78c75fcb347b06c31996860353f40087ed02f48'
Submodule path 'libbpf': checked out '0e43565ad8b4f7bdfa974916e9d7f800157d06ec'
```
- THEN: `cd examples/c`
- THEN: `make minimal`
```shell
  LIB      libbpf.a
make[1]: pkg-config: Command not found
  MKDIR    /Users/amitd2/work/libbpf-bootstrap/examples/c/.output//libbpf/staticobjs
  CC       /Users/amitd2/work/libbpf-bootstrap/examples/c/.output//libbpf/staticobjs/bpf.o
bpf.c:28:10: fatal error: 'asm/unistd.h' file not found
#include <asm/unistd.h>
         ^~~~~~~~~~~~~~
1 error generated.
make[1]: *** [/Users/amitd2/work/libbpf-bootstrap/examples/c/.output//libbpf/staticobjs/bpf.o] Error 1
make: *** [/Users/amitd2/work/libbpf-bootstrap/examples/c/.output/libbpf.a] Error 2
```
- DETOUR: Understand `journal-libbpfgo.md` first
- 