### State of BPF Code 2021


```yaml
- https://github.com/cloudflare/ebpf_exporter
- metrics - prometheus

https://nakryiko.com/posts/libbpf-bootstrap/
- scaffolding - bootstrap
```

```yaml
// send data from BPF to userspace
// vs PerfBuf
//
// PerfBuf:
// is a collection of per CPU circular buffers
// inefficient use of memory & event re-ordering
// separate buf for each CPU
// trade off for big or small CPU buffer // coupled with app behaviour
//
// RingBuf:
// multi producer single consumer queue // MPSC
// safely shared across multiple CPUs simultaneously
// variable length data records
// read from userspace via memory mapped regions // no extra mem copy or syscalls
// epoll as well busy-loop support
// one big common buffer
// goving from 16 to 32 CPUs does not require twice as big a buffer to accommodate more load
// if BPF app needs to track kernel events then ordering is crucial
// per CPU makes it difficult to ensure ordering // refer Notes
//
// RingBuf:
// uses a lightweight spin-lock which means data reservation may fail
// if lock is contended in NMI (non maskable interrupt) context
//
// Notes:
// kernel events e.g. fork(), exec(), exit() can happen in a very rapid
// succession on different CPUs for short-lived processes due to the kernel
// scheduler migrating them from one CPU to another
https://nakryiko.com/posts/bpf-ringbuf/
```

```
// metadata // function info // line info for source
//
// BPF spec contains two parts:
//
// - BTF Kernel API // contract between user space & Kernel
// - BTF ELF File Format // contract between ELF file & BPF loader
//
https://www.kernel.org/doc/html/latest/bpf/btf.html
```


# ------------------------------------------------
# BPF Example - Node Proxy
# ------------------------------------------------

```c
// node proxy via socket level ebpf
//
// By hooking connections at socket-level, the packet-level NAT would be
// bypassed: for each connection, we just need to perform NAT once! (for TCP)

static int
__sock4_xlate_fwd(struct bpf_sock_addr *ctx)
{
    const __be32 cluster_ip = 0x846F070A; // 10.7.111.132
    const __be32 pod_ip = 0x0529050A;     // 10.5.41.5

    if (ctx->user_ip4 != cluster_ip) {
        return 0;
    }

    ctx->user_ip4 = pod_ip;
    return 0;
}

// connect4 indicates that this piece of code will be triggered when there are
// IPv4 socket connection events (connect() system call). And when it happens,
// the code will modify the socket metadata, replacing destination IP (ClusterIP)
// with PodIP then return (continue connecting process, but with new destination IP)
//
// This hooking operates so early (socket-level, above TCP/IP stack in the kernel)
// that even packets (skb) are not generated at this point. Later, all packets
// (including TCP handshakes) will directly use PodIP as destination IP, so no
// packet-level NAT will be involved.
__section("connect4")
int sock4_connect(struct bpf_sock_addr *ctx)
{
    __sock4_xlate_fwd(ctx);
    return SYS_PROCEED;
}
```


```c
// node proxy via tc-level eBPF
// This is packet level NAT i.e. NAT is performed on every single packet
//
// With our toy proxiers, applications (whether it’s a host app, or an app
// running inside a VM/container) on a non-k8s-node (thus not in K8S cluster)
// can also access K8S services with ClusterIP - note that in Kubernetes’s
// design, ClusterIP is only accessible within K8S cluster nodes.
// (In some sense, our toy proxier turns non-k8s-nodes into K8S nodes.)
//
// Think about the role of the node proxy: it actually acts as a reverse proxy
// in the K8S network model. That is, on each node, it will:
//
// - Hide all backend Pods to all clients
// - Filter all egress traffic (requests to backends)
// - For ingress traffic, it does nothing.
```
```c
__section("egress")
int tc_egress(struct __sk_buff *skb)
{
    const __be32 cluster_ip = 0x846F070A; // 10.7.111.132
    const __be32 pod_ip = 0x0529050A;     // 10.5.41.5

    const int l3_off = ETH_HLEN;    // IP header offset
    const int l4_off = l3_off + 20; // TCP header offset: l3_off + sizeof(struct iphdr)
    __be32 sum;                     // IP checksum

    void *data = (void *)(long)skb->data;
    void *data_end = (void *)(long)skb->data_end;
    if (data_end < data + l4_off) { // not our packet
        return TC_ACT_OK;
    }

    struct iphdr *ip4 = (struct iphdr *)(data + l3_off);
    if (ip4->daddr != cluster_ip || ip4->protocol != IPPROTO_TCP /* || tcp->dport == 80 */) {
        return TC_ACT_OK;
    }

    // DNAT: cluster_ip -> pod_ip, then update L3 and L4 checksum
    sum = csum_diff((void *)&ip4->daddr, 4, (void *)&pod_ip, 4, 0);
    skb_store_bytes(skb, l3_off + offsetof(struct iphdr, daddr), (void *)&pod_ip, 4, 0);
    l3_csum_replace(skb, l3_off + offsetof(struct iphdr, check), 0, sum, 0);
	l4_csum_replace(skb, l4_off + offsetof(struct tcphdr, check), 0, sum, BPF_F_PSEUDO_HDR);

    return TC_ACT_OK;
}
```
```c
__section("ingress")
int tc_ingress(struct __sk_buff *skb)
{
    const __be32 cluster_ip = 0x846F070A; // 10.7.111.132
    const __be32 pod_ip = 0x0529050A;     // 10.5.41.5

    const int l3_off = ETH_HLEN;    // IP header offset
    const int l4_off = l3_off + 20; // TCP header offset: l3_off + sizeof(struct iphdr)
    __be32 sum;                     // IP checksum

    void *data = (void *)(long)skb->data;
    void *data_end = (void *)(long)skb->data_end;
    if (data_end < data + l4_off) { // not our packet
        return TC_ACT_OK;
    }

    struct iphdr *ip4 = (struct iphdr *)(data + l3_off);
    if (ip4->saddr != pod_ip || ip4->protocol != IPPROTO_TCP /* || tcp->dport == 80 */) {
        return TC_ACT_OK;
    }

    // SNAT: pod_ip -> cluster_ip, then update L3 and L4 header
    sum = csum_diff((void *)&ip4->saddr, 4, (void *)&cluster_ip, 4, 0);
    skb_store_bytes(skb, l3_off + offsetof(struct iphdr, saddr), (void *)&cluster_ip, 4, 0);
    l3_csum_replace(skb, l3_off + offsetof(struct iphdr, check), 0, sum, 0);
	  l4_csum_replace(skb, l4_off + offsetof(struct tcphdr, check), 0, sum, BPF_F_PSEUDO_HDR);

    return TC_ACT_OK;
}

char __license[] __section("license") = "GPL";
```

# --------------------------------------------
# BPF - Lifetime of BPF Objects
# --------------------------------------------

https://facebookmicrosites.github.io/bpf/blog/2018/08/31/object-lifetime.html

```c
// File descriptors and Reference counters
//
// BPF objects (progs, maps, and debug info) are accessed by user space via
// file descriptors (FDs),
//
// Each BPF object has a reference counter. For example, when a map is created
// with a call to bpf_create_map(), the kernel allocates a struct bpf_map object
// The kernel then initializes its refcnt to 1, and returns a file descriptor
// to the user space process.
//
// If the process exits or crashes right after, the FD would be closed, and the
// refcnt of the bpf map object will be decremented. At this point the refcnt
// will be zero, which will trigger a memory free after the RCU grace period.
```

```c
// BPF programs that use BPF maps are loaded in two phases.
//
// First, the maps are created and their FDs (file descriptors) are stored as
// part of the BPF program in the 'imm' field of BPF_LD_IMM64 instructions.
// When the kernel verifies the program, it increments refcnt of maps used by
// the program, and initializes program's refcnt to 1.
//
// At this point user space can close FDs associated with maps, but maps will
// not disappear since the program is “using” them (though the program is not
// yet attached anywhere). When prog FD is closed, and refcnt reaches zero, the
// destruction logic will iterate over all maps used by the program and will
// decrement their refcnts. This scheme allows the same map to be used by
// multiple programs at once (even of different program types). For example, a
// tracing program attached to a tracepoint can collect data into a bpf map,
// while a networking program uses that information to make forwarding decisions.
```

```c
// When a program is attached to some hook, the refcnt of the program is
// incremented. The user space process that originally created the BPF
// maps+program, and then loaded and finally attached the program to the hook,
// can now exit.
//
// At this point the maps+program that were created by user space will stay alive,
// since the program has a refcnt > 0. This is the BPF object lifecycle! As long
// as the refcnt of a BPF object (program or map) is > 0, the kernel will keep it
// alive.
```

```c
// Not all attachment points are made equal though.
//
// XDP, tc's clsact, and cgroup-based hooks are global.
//
// Programs will stay attached to global attachment points for as long as those
// objects are alive.
//
// In case of clsact, the program is attached to an ingress or egress qdisc.
// If nothing is doing (for example) tc qdisc del, the program will be processing
// ingress or egress packets even when there is no corresponding user space
// process. On the other hand, you may have programs attached to (for example)
// tracing hooks that will only run for the lifetime of the process that holds
// FD to tracing event.
//
// Say for example your program gets an FD from perf_event_open() and then does
// ioctl(perf_event_fd, IOC_SET_BPF, bpf_prog_fd) (or, alternatively, gets an
// FD from bpf_raw_tracepoint_open(“tracepoint_name”, bpf_prog_fd)). These FDs
// are local to the process, and therefor if the process crashes the perf_event_fd
// or bpf_raw_tracepoint_fd will be closed. In this scenario, the kernel will
// detach the bpf program and decrement its refcnt.
//
// query - why should the process crash?
```

```c
// In summary: xdp, tc, lwt, cgroup hooks are “global”, whereas kprobe, uprobe,
// tracepoint, perf_event, raw_tracepoint, socket filters, so_reuseport hooks
// are “local” to the process, since they're accessed via FD.
//
// Note that people have requested a new type of cgroup object that is FD based,
// so in the future it is possible that cgroup-bpf might be both “global” as well
// as “local”.
```

```c
// The main advantage of the FD based interface is auto-cleanup, meaning that if
// anything goes wrong with the user space process the kernel will clean up all
// objects automatically. The original BPF API was FD based from the beginning.
// However, while deploying kprobe/uprobe + bpf in production it became clear
// that using the global interface of [ku]probe is too cumbersome to keep working
// around. Hence, a FD based [ku]probe API was introduced in the kernel. A similar
// situation may happen soon with the cgroup API. There could be a cgroup object
// that a process holds with an FD, and attaches BPF programs to that FD (object)
// instead of to a global cgroup entity.
```

```c
// The FD based API is useful for networking as well. Some time ago a
// Facebook Widely Deployed Binary (WDB) had a bug where it 'forgot' to cleanup
// tc's clsact bpf program. As a result, after many daemon restarts there were
// thousands of bpf programs attached to the same clsact egress hook doing the
// same work. All of these programs except one were running without a corresponding
// user space process and were wasting cpu cycles. Eventually this problem was
// noticed when overall system performance gradually degraded. Using a FD based
// networking API would have prevented such situation. Note that there is ongoing
// work to introduce a new tc clsact-like API which adds hooks into ingress and
// egress paths of network devices, but is not actually based on tc and at the
// same time is FD based with auto-cleanup properties.
```

```c
// BPFFS
//
// There is another way to keep BPF programs and maps alive. It's called BPFFS,
// or, “BPF File System”.
//
// A user space process can “pin” a BPF program or map in BPFFS using an arbitrary
// name. Pinning a BPF object in BPFFS will increment the refcnt of the object,
// which will result in the BPF object staying alive, even if the pinned BPF
// program is not attached anywhere or the pinned BPF map is not used by any program.
// An example of where this is useful comes to us from networking. For networking
// you may have BPF programs that can do their packet processing duties without
// a user space daemon, but the admin may want to login to the host and examine
// the map from time to time. The admin could bpf_obj_get(“path”) the object from
// BPFFS, which will return a new FD (and corresponding handle to the object).
// To unpin the object, just delete the file in BPFFS with unlink() and the
// kernel will decrement corresponding refcnt. In summary:
//
// obj create → refcnt=1 attach → refcnt++, detach →refcnt-- obj_pin → refcnt++,
// unlink →refcnt-- close → refcnt--
```

# --------------------------------------------
# BTF - BPF Type Format
# --------------------------------------------

https://nakryiko.com/posts/btf-dedup/

```c
// One way to smooth out the developer experience with BPF is through better
// introspection and debuggability of BPF code.
//
// To allow that BPF needs to know a bunch of metadata information about the
// BPF program. Type information is one of a few critical pieces of metadata
// necessary to bring the BPF experience to the next level. Eventually, this
// metadata will enable not just introspection, but also much higher levels of
// re-usability of BPF code across different versions of the Linux kernel
// (the so-called "compile once, run everywhere" approach, BPF CO-RE).
```

```c
// How is such metadata represented in user-space (non-BPF) programs?
//
// It's typically done through DWARF debug information. The problem with DWARF
// though, is that it's very generic and expressive, which makes it quite a
// complicated and verbose format, and thus unsuitable to include in the Linux
// kernel image due to the size overhead and the complexity of parsing it.
```

```c
// Enter BTF (BPF Type Format). It's a minimalistic, compact format, inspired by Sun's CTF (Compact C Type Format), which is used for representing kernel debug information since Solaris 9. BTF was created for similar purposes with a focus on simplicity and compactness to allow its usage in the Linux kernel and BPF.
```

```c
// BTF represents each type with one of a few possible type descriptors identified
// by a kind: BTF_KIND_INT, BTF_KIND_ENUM, BTF_KIND_STRUCT, BTF_KIND_UNION,
// BTF_KIND_ARRAY, BTF_KIND_FWD, BTF_KIND_PTR, BTF_KIND_CONST, BTF_KIND_VOLATILE,
// BTF_KIND_RESTRICT, BTF_KIND_TYPEDEF — a few more might be added soon for
// functions support, etc.
```

# --------------------------------------------
# BPF - Compile Once Run Everywhere
# --------------------------------------------

https://facebookmicrosites.github.io/bpf/blog/2020/02/19/bpf-portability-and-co-re.html

```c
// Compiler Support
//
// To enable BPF CO-RE and let BPF loader (i.e., libbpf) to adjust BPF program
// to a particular kernel running on target host, Clang was extended with
// few built-ins.
//
// If you were going to access task_struct->pid field, Clang would record that
// it was exactly a field named "pid" of type “pid_t” residing within a struct
// task_struct.
//
// This is done so that even if target kernel has a task_struct layout in which
// “pid” field got moved to a different offset within a task_struct structure
// (e.g., due to extra field added before “pid” field), or even if it was moved
// into some nested anonymous struct or union (and this is completely transparent
// in C code, so no one ever pays attention to details like that), we’ll still
// be able to find it just by its name and type information. This is called a
// **field offset relocation**
//
// It is possible to capture (and subsequently relocate) not just a field offset,
// but other field aspects, like field existence or size. Even for bitfields
// (which are notoriously "uncooperative" kinds of data in the C language,
// resisting efforts to make them relocatable) it is still possible to capture
// enough information to make them relocatable, all transparently to BPF program
// developer.
```

```c
// #include "vmlinux.h" // just this & no more includes
//
// bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
```

```c
// Reading a field from a kernel struct // via BCC // error prone

pid_t pid = task->pid; // BCC rewrites task->pid into bpf_probe_read()
```

```c
// Reading a field from a kernel struct // via Libbpf + BPF_PROG_TYPE_TRACING

pid_t pid = task->pid;
```

```c
// Reading a field from a kernel struct // via BPF_PROG_TYPE_TRACING + BPF CO-RE

pid_t pid = __builtin_preserve_access_index(({ task->pid; })); // is portable
```

```c
// Reading a field from a kernel struct // via Non-CO-RE libbpf way

pid_t pid;
bpf_probe_read(&pid, sizeof(pid), &task->pid);
```

```c
// Reading a field from a kernel struct // via CO-RE libbpf way

pid_t pid;
bpf_core_read(&pid, sizeof(pid), &task->pid); // bpf_probe_read is replaced by bpf_core_read
```

```c
// bpf_core_read() is a simple macro
// a/ passes all the arguments directly to bpf_probe_read()
// b/ also makes Clang **record** the **field offset relocation** for third argument (&task->pid)
// by passing it through __builtin_preserve_access_index()
//
// So the last example is actually just this, under the hood:
// bpf_probe_read(&pid, sizeof(pid), __builtin_preserve_access_index(&task->pid));
```

```c
// Get inode number for current executable binary

u64 inode = task->mm->exe_file->f_inode->i_ino; // BCC way
u64 inode = BPF_CORE_READ(task, mm, exe_file, f_inode, i_ino); // BPF CO_RE way

// -- aliter --
u64 inode;
BPF_CORE_READ_INTO(&inode, task, mm, exe_file, f_inode, i_ino); // Aliter BPF CO-RE way
```

```c
// Does field exist in target kernel?

pid_t pid = bpf_core_field_exists(task->pid) ? BPF_CORE_READ(task, pid) : -1;
```

```c
// How to capture any field’s size?

u32 comm_sz = bpf_core_field_size(task->comm); /* will set comm_sz to 16 */
```

```c
// How to read bitfield out of a kernel struct?

struct tcp_sock *s = ...;

/* with direct reads */
bool is_cwnd_limited = BPF_CORE_READ_BITFIELD(s, is_cwnd_limited);

/* with bpf_probe_read()-based reads */
u64 is_cwnd_limited;
BPF_CORE_READ_BITFIELD_PROBED(s, is_cwnd_limited, &is_cwnd_limited);
```

```c
// Handle task_struct’s utime field switched from in jiffies to nanoseconds?
// > 4.6 kernel

extern u32 LINUX_KERNEL_VERSION __kconfig; // libbpf provided extern kconfig variable
extern u32 CONFIG_HZ __kconfig;

u64 utime_ns;

if (LINUX_KERNEL_VERSION >= KERNEL_VERSION(4, 11, 0))
  utime_ns = BPF_CORE_READ(task, utime);
else
  /* convert jiffies to nanoseconds */
  utime_ns = BPF_CORE_READ(task, utime) * (1000000000UL / CONFIG_HZ);
```

```c
// Handle fs/fsbase (which got renamed in recent kernels)?

/* up-to-date thread_struct definition matching newer kernels */
struct thread_struct {
    ...
    u64 fsbase;
    ...
};

/* legacy thread_struct definition for <= 4.6 kernels */
/* ___v46 is a "flavor" part */
struct thread_struct___v46 {   
    ...
    u64 fs;
    ...
};

extern int LINUX_KERNEL_VERSION __kconfig;
...

struct thread_struct *thr = ...;
u64 fsbase;
if (LINUX_KERNEL_VERSION > KERNEL_VERSION(4, 6, 0))
    fsbase = BPF_CORE_READ((struct thread_struct___v46 *)thr, fs);
else
    fsbase = BPF_CORE_READ(thr, fsbase);
```

```c
// app-provided read-only configuration and struct flavors are an ultimate big
// hammer to address whatever complicated scenario application has to handle
// From the user-space side, application will be able to easily provide this
// configuration through BPF skeleton.

/* global read-only variables, set up by control app */
const bool use_fancy_helper;
const u32 fallback_value;

...

u32 value;
if (use_fancy_helper)
    value = bpf_fancy_helper(ctx);
else
    value = bpf_default_helper(ctx) * fallback_value;
```


```c
// BPF is using locked memory for BPF maps and various other things.
// By default, this limit is very low, so unless it’s increased,
// even a trivial BPF program won’t load successfully into kernel.
// BCC unconditionally sets this limit to infinity, but libbpf doesn’t
// do this automatically (by design).
//
// Depending on your production environment, there might be better and
// more preferred ways of doing this. But for quick experimentation or
// if there is no better way of doing this, you can do it yourself
// through setrlimit(2) syscall, which should be called at the very
// beginning of your program:

#include <sys/resource.h>

rlimit rlim = {
    .rlim_cur = 512UL << 20, /* 512 MBs */
    .rlim_max = 512UL << 20, /* 512 MBs */
};

err = setrlimit(RLIMIT_MEMLOCK, &rlim);
if (err)
    /* handle error */
```

```c
// Libbpf log
// When something doesn’t work as expected, the best way to start investigating
// is to look at libbpf log output. Libbpf outputs a bunch of useful logs at
// various levels of verbosity. By default, libbpf will emit error-level output
// to console. We recommend installing a custom logging callback and set up
// ability to turn on/off verbose debug-level output

int print_libbpf_log(enum libbpf_print_level lvl, const char *fmt, va_list args) {
    if (!FLAGS_bpf_libbpf_debug && lvl >= LIBBPF_DEBUG)
        return 0;
    return vfprintf(stderr, fmt, args);
}

/* ... */

/* set above as the custom log handler */
libbpf_set_print(print_libbpf_log);
```

```c
// support both BCC and libbpf "modes"
// Rely on BCC_SEC macro in BCC

#ifdef BCC_SEC
#define __BCC__
#endif

// After this, throughout your BPF code, you can do:

#ifdef __BCC__
/* BCC-specific code */
#else
/* libbpf-specific code */
#endif
```

```c
#ifdef __BCC__
/* linux headers needed for BCC only */
#else /* __BCC__ */
#include "vmlinux.h"   /* all kernel types */
#include <bpf/bpf_helpers.h>       /* most used helpers: SEC, __always_inline, etc */
#include <bpf/bpf_core_read.h>     /* for BPF CO-RE helpers */
#include <bpf/bpf_tracing.h>       /* for getting kprobe arguments */
#endif /* __BCC__ */
```

```c
// BPF maps
// The way that BCC and libbpf define BPF maps declaratively is different,
// but conversion is very straightforward. Here are some of the examples:

/* Array */
#ifdef __BCC__
BPF_ARRAY(my_array_map, struct my_value, 128);
#else
struct {
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __uint(max_entries, 128);
    __type(key, u32);
    __type(value, struct my_value);
} my_array_map SEC(".maps");
#endif

/* Hashmap */
#ifdef __BCC__
BPF_HASH(my_hash_map, u32, struct my_value);
#else
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 10240);
    __type(key, u32);
    __type(value, struct my_value);
} my_hash_map SEC(".maps")
#endif

/* Per-CPU array */
#ifdef __BCC__
BPF_PERCPU_ARRAY(heap, struct my_value, 1);
#else
struct {
    __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
    __uint(max_entries, 1);
    __type(key, u32);
    __type(value, struct my_value);
} heap SEC(".maps");
#endif
```

```c
// PERF_EVENT_ARRAY, STACK_TRACE and few other specialized maps
// (DEVMAP, CPUMAP, etc) don’t support (yet) BTF types for key/value,
// so specify key_size/value_size directly instead

/* Perf event array (for use with perf_buffer API) */
#ifdef __BCC__
BPF_PERF_OUTPUT(events);
#else
struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
    __uint(key_size, sizeof(u32));
    __uint(value_size, sizeof(u32));
} events SEC(".maps");
#endif
```

```c
// Access BPF map from BPF code
#ifdef __BCC__
    struct event *data = heap.lookup(&zero);
#else
    struct event *data = bpf_map_lookup_elem(&heap, &zero);
#endif

#ifdef __BCC__
    my_hash_map.update(&id, my_val);
#else
    bpf_map_update_elem(&my_hash_map, &id, &my_val, 0 /* flags */);
#endif

#ifdef __BCC__
    events.perf_submit(args, data, data_len);
#else
    bpf_perf_event_output(args, &events, BPF_F_CURRENT_CPU, data, data_len);
#endif
```

```c
// BPF programs
// All functions representing BPF programs need to be marked with custom section
// name through use of SEC() macro, coming from bpf_helpers.h
//
// It’s just a convention, but you’ll get a much better experience overall if
// you follow libbpf’s section naming. Detailed list of expected names can be
// found here. Few most common ones would be:
//
// tp/<category>/<name> for tracepoints;
// kprobe/<func_name> for kprobe and kretprobe/<func_name> for kretprobe;
// raw_tp/<name> for raw tracepoint;
// cgroup_skb/ingress, cgroup_skb/egress, and a whole family of cgroup/<subtype> programs.
```

```c
// tracepoints
#if !defined(__BCC__)
SEC("tracepoint/sched/sched_process_exec")
#endif
int tracepoint__sched__sched_process_exec(
#ifdef __BCC__
    struct tracepoint__sched__sched_process_exec *args
#else
    struct trace_event_raw_sched_process_exec *args
#endif
) {
/* ... */
}
```

```c
// kprobes

#ifdef __BCC__
int kprobe__acct_collect(struct pt_regs *ctx, long exit_code, int group_dead)
#else
SEC("kprobe/acct_collect")
int BPF_KPROBE(kprobe__acct_collect, long exit_code, int group_dead)
#endif
{
    /* BPF code accessing exit_code and group_dead here */
}
```

```c
// Note! Syscall functions got renamed in 4.17 kernels. Starting from 4.17
// version, syscall kprobe that used to be called, say, sys_kill, is called now
// __x64_sys_kill (on x64 systems, other architectures will have different
// prefix, of course). You’ll have to account for that when trying to attach a
// kprobe/kretprobe. If possible, though, try to stick to tracepoints.
//
// N.B. If you are developing a new BPF application with the need for
// tracepoint/kprobe/kretprobe, check out new raw_tp/fentry/fexit probes.
// They provide better performance and usability and are available starting
// from 5.5 kernels.
```

```c
// For dealing with kernel version differences, BPF CO-RE supplies two
// complementary mechanisms: Kconfig externs and struct “flavors”.
// BPF code can know which kernel version it’s dealing with by declaring
// the following extern variable:
// query - is extern variable a macro? function?

#define KERNEL_VERSION(a, b, c) (((a) << 16) + ((b) << 8) + (c))

extern int LINUX_KERNEL_VERSION __kconfig; // extract

if (LINUX_KERNEL_VERSION < KERNEL_VERSION(5, 2, 0)) {
  /* deal with older kernels */
} else {
  /* 5.2 or newer */
}

// Similarly to getting the kernel version, you can extract any CONFIG_xxx value
// from Kconfig:
// query - just like that? you can extract automatically?

extern int CONFIG_HZ __kconfig;

/* now you can use CONFIG_HZ in calculations */
```

```c
// struct flavour to handle changes between kernel version
/* struct kernfs_iattrs will come from vmlinux.h */

struct kernfs_iattrs___old {
    struct iattr ia_iattr;
};

if (bpf_core_field_exists(root_kernfs->iattr->ia_mtime)) {
    data->cgroup_root_mtime = BPF_CORE_READ(root_kernfs, iattr, ia_mtime.tv_nsec);
} else {
    struct kernfs_iattrs___old *root_iattr = (void *)BPF_CORE_READ(root_kernfs, iattr);
    data->cgroup_root_mtime = BPF_CORE_READ(root_iattr, ia_iattr.ia_mtime.tv_nsec);
}
```

```c
// BPF application configuration
// using global variables
//
// global variables can be mutable or constant
//
// mutable is used for bi-directional data exchange between BPF program & its
// user space counterpart
//
// global variables are setup before load & verification phases
```

```c
// On BPF code side
//
// declare read-only global variables using a const volatile global variable
// for mutable ones, just drop const volatile qualifiers
//
// note: const is readonly in c programming

const volatile struct {
    bool feature_enabled;
    int pid_to_filter;
} my_cfg = {}; // query - is this like define & declare & initialise - all together?
```

```c
// const volatile has to be specified to prevent clever compiler optimizations
// compiler might and will erroneously assume zero values and inline them in code
//
// if you are defining a mutable (non-const) variable
// make sure they are not marked as static
// non-static globals interoperate with compiler the best
// volatile is usually not necessary in such case
//
// your variables have to be initialized
// otherwise libbpf will decline to load BPF application
// Initialization can be to zeroes or any other value you need
// Such value will be a default value of variable, unless overridden from a control app
```

```c
// Global variables provide much nicer user experience and avoid BPF map lookup overhead
// query - map lookup is an overhead! how much?
//
// for constant variables, their values are well-known to BPF verifier and
// treated as constants during program verification
// which allows BPF verifier to verify code more precisely and
// eliminate dead code branches effectively.
```

```c
// The way control app provides values for global variables is simple and natural
// with the usage of BPF skeleton:

struct <name> *skel = <name>__open();
if (!skel)
    /* handle errors */

skel->rodata->my_cfg.feature_enabled = true;
skel->rodata->my_cfg.pid_to_filter = 123;

if (<name>__load(skel))
    /* handle errors */
```

```c
// Read-only variables, can be set and modified from user-space only before
// a BPF skeleton is loaded. Once a BPF program is loaded, neither BPF nor
// user-space code will be able to modify it.
//
// Non-const variables, on the other hand, can be modified after BPF skeleton
// is loaded throughout the entire lifetime of BPF program, both from BPF and
// user-space sides. They can be used for exchanging mutable configuration, stats, etc.
```

```c
// global variables usage
//
// use of global variables in BPF code is trivial
//
// BPF global variables look and behave exactly like a user-space variables
// they can be used in expressions, updated (the non-const ones), you can even
// take their address and pass around into helper functions. But that is only
// true for the BPF code side.
//
// From user-space, they can be read and updated only through BPF skeleton:
//
// skel->rodata for read-only variables;
// skel->bss for mutable zero-initialized variables;
// skel->data for non-zero-initialized mutable variables.
//
// You can still read/update them from user-space and those updates will be
// immediately reflected on BPF side. But they are not global variables on the
// user-space side, they are just members of BPF skeleton’s rodata, bss, or data
// members, which are initialized during the skeleton load phase. This,
// subsequently, means that declaring exactly the same global variable in BPF
// code and user-space code will declare completely independent variables,
// which won’t be connected in any way.
```

```c
// loop unrolling
//
// Unless you are targeting 5.3+ kernel, all the loops in your BPF code have
// to be marked with #pragma unroll to force Clang to unroll them and eliminate
// any possible control flow loops:

#pragma unroll
for (i = 0; i < 10; i++) { ... }

// Without loop unrolling or if the loop doesn’t terminate within fixed amount
// of iterations, you’ll get a verifier error about “back-edge from insn X to Y”,
// meaning that BPF verifier detected an infinite loop (or can't prove that
// loop will finish in a limited amount of iterations).
```

```c
// Helper sub-programs
//
// If you are using static helper functions, they have to be marked as
// static __always_inline, due to current limitations in libbpf’s handling of them

static __always_inline unsigned long
probe_read_lim(void *dst, void *src, unsigned long len, unsigned long max)
{
    ...
}

// Non-inlined global functions are also supported starting from 5.5 kernels,
// but they have different semantics and verification constraints than static
// functions. Make sure to check them out as well!
```

```c
// bpf_printk debugging
//
// There is no conventional debugger available for BPF programs
//
// For such cases, logging extra debug information is your best bet
// Use bpf_printk(fmt, args...) to emit extra pieces of data to help understand
// what's going on. It accepts printf-like format string and can handle only up
// to 3 arguments.
//
// It's simple and easy to use, but it's quite expensive, making it unsuitable
// to be used in production. So it's mostly appropriate only for ad-hoc debugging

char comm[16];
u64 ts = bpf_ktime_get_ns(); // query - ns vs ts?
u32 pid = bpf_get_current_pid_tgid(); // query - tgid?

bpf_get_current_comm(&comm, sizeof(comm)); // query - current command?
bpf_printk("ts: %lu, comm: %s, pid: %d\n", ts, comm, pid);

// Logged messages can be read from a special /sys/kernel/debug/tracing/trace_pipe file:
//
// $ sudo cat /sys/kernel/debug/tracing/trace_pipe
// ...
//       [...] ts: 342697952554659, comm: runqslower, pid: 378
//       [...] ts: 342697952587289, comm: kworker/3:0, pid: 320
// ...
```


# ==============================================
# XDP
# ==============================================

```c
// What is eXpress Data Path XDP ?
// -----
// process packet out of the driver code i.e. bare metal speed
// no need of queues // hence no lock
// no need to allocate CPU to a queue
//
// BPF program does the following:
// ------
// packet parsing
// table look ups
// manage stateful filters
// encap/decap packets, etc
//
// XDP delivers the following:
// -------
// removes the need to allocate large pages
// removes the need for dedicated CPUs
// removes the need to inject packets into kernel from a user space app
// removes the need to define a new security model for accessing networking hardware
//
- https://www.iovisor.org/technology/xdp
```

# ==============================================
# Blogs / Videos -- BPF / Generic
# ==============================================

```c
// BPF 101 talk by Liz Rice - May 2019
//
// clang converts the limited c code to BPF bytecode
//
// kernel contains a just-in-time (JIT) compiler that translates this BPF
// bytecode into native machine code for better performance
//
// (limited) C -> BPF bytecode -> machine code
```

```c
// BCC makes BPF programs easier to write
// 1/ has kernel instrumentation in C (runs in kernel)
// 2/ includes a C wrapper around LLVM
// 3/ exposes frontends in Python & Lua (user space code that uses BPF syscalls)
//
// run the userspace python code as ./hello_world.py
// see what's happening by running it with strace
// fellow - use strace as a kind of debugging BPF programs
//
// strace -e bpf, read, write, openat, perf_event_open, ioctl ./hello_world.py
```

```c
// eBPF maps
//
// generic data structure // store different kinds of data
//
// sharing between eBPF kernel programs &
// sharing between kernel & user-space applications
```

```c
https://www.youtube.com/watch?v=4SiWL5tULnQ
```

```c
// When it comes to filtering packets on a Linux server, many options are available.
// Firewalls (iptables/nftables) and Traffic Control (TC) filtering are different
// hooks that can be used in the networking stack.
//
// Higher up on the system, filtering on sockets is done for programs like
// tcpdump or Wireshark. In the opposite direction, directly on the hardware,
// some NICs support hardware “n-tuple” filters set up with ethtool.
- https://dzone.com/articles/libkefir-all-your-rules-in-one-bottle
```

```c
// handle DDOS attacks

- https://blog.cloudflare.com/introducing-the-bpf-tools/
```

```c
// Xtables to match packets using a BPF filter
// BPF filter match

- https://github.com/torvalds/linux/blob/master/net/netfilter/xt_bpf.c
```

```c
// general purpose execution engine & not a Virtual Machine
// 2 design goals - low overhead - esp. for x86-64, arm64
// verifiable for safety by kernel at program load time
//
// its an instruction set with C calling convention in mind
// approx 150 BPF kernel helpers, 30 maps
// share data with userspace with BPF maps
//
// How far is BPF C from generic C?
// BPF to BPF function calls // bounded loops // global variables
// static linking // BTF // upto 1 Million instructions / program
//
// w.r.t K8s
// migrate load balancing from kube proxy to BPF at XDP layer
// north south load balancing case
// reduces CPU cost // achieves DPDK speeds
//
// LoadBalancer - client to remote node 2 - via - cilium
// packet arrives at node 1 LB -> DNAT-ed to remote node 2
// src ip address of client is preserved -> reply from node 2 to client directly
- https://youtu.be/PJY-rN1EsVw // WIP
```

```c
// BPF CO-RE - Compile Once Run Everywhere // Fellow
// ----------------------------------------
// BCC (BPF Compiler Collection) that uses embedded Clang/LLVM is bad
// BTF (BPF Type Format) is better
// no compile time #ifdef/ #else required
//
- https://facebookmicrosites.github.io/bpf/blog/2020/02/19/bpf-portability-and-co-re.html
```

```c
// Developing a BPF program with the following steps:
// 1/ Write the BPF code in C
// 2/ Compile the code for the BPF VM
// 3/ Write a user space component that loads the output of step 2. into the BPF VM
// 4/ Use the BPF API to exchange data between the user space component and the BPF code
//
- https://blog.redsift.com/labs/writing-bpf-code-in-rust/
```

```c
// control traffic leaving the cluster to internet
// avoid exfiltration of data by malicious workloads
//
// Kubernetes NetworkPolicy resource
// A crd // but implementation is via CNI // calico // Cilium
// NetworkPolicy is Pod <-> pod, svc, ingress, egress
// -- is not for node level or cluster level
//
- https://kinvolk.io/blog/2020/12/egress-filtering-benchmark-part-2-calico-and-cilium/
```

```c
// intelligent load balancing // bpf // ddos mitigation
// user mode app          -> syscalls         -> kernel -> hardware
// kernel mode apps (BPF) -> BPF helper calls -> kernel -> Hardware
// wait // block // sleep // idle
//
// no context switches // no scheduler // bpf runs to completion // spin locks
// similar to nodejs non-blocking programming
//
// at facebook, netflix, cloudflare
// modern linux is becoming smaller i.e. micro kernel
//
// BPF // 1992 // Berkeley Packet Filter // a limited virtual machine
// a limited VM for packet filters // a byte code // a VM inside the kernel
// limited & minimum VM is BPF // eBPF is extended
//
// SDN configuration // DDoS Mitigation // Intrusion Detection
// Container Security // Observability // Firewalls // Device Drivers
//
// BPF internals
// BPF Instructions // Events
// Verifier // Interpreter // JIT Compiler // BPF Context
// 11 Registers // Machine Code Execution // BPF Helpers
// Map Storage (Mbytes)
//
// Spectre Meltdown
// all layers changed // cpu // hypervisor// runtime // compilers // app code
// however BPF byte code was not affected
// if there is any security issue?
// -- then patch the jit compiler
// -- all s/w immediately recompiles
//
// Is BPF turing complete?
// Can BPF program run BPF programs?
// No. Verifier rejects unbounded loops!
- https://youtu.be/7pmXdG8-7WU
```

```c
// which processes are connecting to which port?
// a BPF tool // better than tcpdump
// https://github.com/iovisor/bcc/blob/master/tools/tcplife.py
//
// bpftrace is a high level language
// #!/usr/local/bin/bpftrace
// bpftrace is great to create single page programs
- tcplife
```

```c
// bpf
// kernel api
// - define & really define the problem statement
// -- do not let others to guess
// - consider existing interfaces
// - extensible & future needs
// - holes in the design // can it be exploited
// - will it be obsolete soon
// by David S Miller
- https://youtu.be/mFxs3VXABPU
```

```c
// run sandboxed programs in OS kernel // no need to change kernel code
// no need to load kernel modules
- https://ebpf.io/
```

```c
// eBPF vs Service Mesh // custom hooks in kernel
// user space -> syscalls -> kernel + eBPF
// internet // eth0 // ethernet // tcp/ip // socket // app // image // diagram
- https://thenewstack.io/how-ebpf-streamlines-the-service-mesh/
```

### TODO
- https://qmonnet.github.io/whirl-offload/2016/09/01/dive-into-bpf/ // awesome // fellow
- https://github.com/xdp-project/xdp-tutorial // awesome // fellow
- https://qmonnet.github.io/whirl-offload/ // awesome
- https://github.com/iovisor/bpftrace/blob/master/docs/tutorial_one_liners.md // samples
- https://github.com/iovisor/bpftrace/blob/master/docs/reference_guide.md // guide
- https://github.com/iovisor/bpftrace/blob/master/man/adoc/bpftrace.adoc // manual
- https://www.kernel.org/doc/Documentation/kprobes.txt
- https://www.iovisor.org/technology/xdp
- https://docs.cilium.io/en/latest/bpf/ // XDP
- https://pingcap.com/blog/tips-and-tricks-for-writing-linux-bpf-applications-with-libbpf
- https://facebookmicrosites.github.io/bpf/blog/2020/02/19/bpf-portability-and-co-re.html
- https://nakryiko.com/posts/btf-dedup/
- https://nakryiko.com/posts/bpf-portability-and-co-re/
- https://nakryiko.com/posts/libbpf-bootstrap/
- https://dzone.com/articles/libkefir-all-your-rules-in-one-bottle
- https://github.com/weaveworks/tcptracer-bpf
- https://blog.cloudflare.com/introducing-the-bpf-tools/
- https://github.com/libbpf/libbpf // 1st
- https://github.com/libbpf/libbpf-bootstrap
- https://github.com/cilium // 2nd
- https://github.com/iovisor/gobpf
- https://github.com/iovisor/kubectl-trace
- https://github.com/iovisor/bcc // legacy
- https://github.com/cloudflare/bpftools
- https://git.kernel.org/pub/scm/linux/kernel/git/bpf/bpf-next.git/tree/tools/testing/selftests/bpf // testing
