### State of Linux 2021

### PID Namespace
```yaml
- https://www.youtube.com/watch?v=l4I2TVAnBuw
```

### Memory
```yaml
- https://www.bwplotka.dev/2019/golang-memory-monitoring/

- Go runtime in certain cases uses madvise system call
- With madvise Go processes can closely cooperate with the kernel 
- On how to treat certain “pages” of RAM memory in virtual space in a way that helps both sides

- From a high-level perspective the madvise system call consists of 3 arguments
- address and length that define what memory range this call refers to
- advice that says what to advice for those memory pages

- we are interested in two specific values of advice:

- MADV_DONTNEED
- Do not expect access in the near future
- application is finished with the given range
- so the kernel can free resources associated with it

- MADV_FREE (since Linux kernel 4.5)
- application no longer requires the pages in the range specified by addr and len
- kernel can thus free these pages
- but the freeing could be delayed until memory pressure occurs
```

### ptrace
```yaml
- https://man7.org/linux/man-pages/man2/ptrace.2.html
- attach - observe - control execution of another process
- breakpoint - debugging - system call tracing
```

### OS
```yaml
- https://www.cs.bham.ac.uk/~exr/lectures/opsys/10_11/lectures/os-dev.pdf
- Write an OS from scratch
- https://github.com/cfenollosa/os-tutorial
- http://littleosbook.github.io/
```

### Books
```yaml
- Understanding the Linux Kernel, Third Edition, Marco Cesati
- Linux Kernel Development, 3rd Edition, Robert Love
- Linux System Programming, Talking Directly to the Kernel & C Library, 2nd Edition
- The Linux Programming Interface: A Linux & UNIX System Programming Handbook
- Linux Device Drivers, Third Edition, Alessandro Rubini - https://lwn.net/Kernel/LDD3/
- http://littleosbook.github.io/
```
