### State of FileSystem 2021

```c
// Tracing and Visualizing File System Internals with eBPF
//
// How many files open, or being read?
// Which user / processes is operating these files?
// How many files can be created?
// What data is most frequently accessed?
//
// How are these files structured?
// What happens when we move them across?
// What if file is moved from EFS to a pen drive?
```
```c
// App to Storage
// 1/ int fd = open("foo", R)
// 2/ sysopen // Syscall Interface
// 3/ vfs_open // Virtual File System
// 4/ ext4_file_open // Logical File System
// 5/ block devices
// 6/ drivers
// 7/ disk
```
```c
// Cache in between
//
// At VFS - inode cache, directory cache, page cache
//
// fellow - Learn Linux Kernel Map - A diagram that is as complex as CNCF landscape
//
// trace a read() call with ftrace
// 8MB in 2 seconds // 1 read syscall // 5 other syscalls // 5550 function calls
//
// note: IRQ Entry & IRQ Exit - stall in a syscall
```

```c
// /proc filesystem
//
// cat /proc/diskstats
// whats happening with disk, process, block requests?
//
// cons: difficult to understand
```

```c
// iostat
//
// built over /proc & /sys
// gives human readable info e.g. disk pressure, etc.
//
// similarly iotop
```

```c
// Why more? further?
//
// Tracing! Why something has happened? e.g. IRQ entry & IRQ exit
// Targeted Analysis
// Live Analysis! Sometimes post mortem analysis is not enough!
// Live Debugging! Traffic Shaping!
// Live at production setup
//
// via expressions / programs i.e. BPF Programs
```

```c
https://www.youtube.com/watch?v=2SqPdM-YUaw
```
