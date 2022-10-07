## Learn containerd from its commits
This is one of my ideas to learn a project. In other words, read and perhaps try interesting
commits from the project. This should help me in understanding parts of the project by focusing
on some particular fix or feature. Alternative ways e.g. getting involved with the community
or spending weeks &/ months may not be feasible unless it is part of my day job. Needless to say
this works better when the project follows atomic commits.

## Commits
### Add support of CAP_BPF and CAP_PERFMON
- Prior to kernel 5.8 bpf and perf_event_open requires CAP_SYS_ADMIN
- This change enables finer control of the privilege setting
- Thus allowing us to run certain system tracing tools with minimal privileges
- PR: https://github.com/containerd/containerd/pull/7301
- FILE: contrib/seccomp/seccomp_default.go

### Remove tun/tap from the default devices
- A container should not have access to tun/tap device
- Unless it is explicitly specified in configuration
- This device was already removed from docker's default, and runc's default
- PR: https://github.com/containerd/containerd/pull/6923
- FILE: oci/spec_opts.go

### TIL: Vagrantfile for testing
- PR: https://github.com/containerd/containerd/pull/7265

### Rollback Ubuntu to 18.04 (except for riscv64)
- Rollback the build environment from Ubuntu 22.04 to 18.04
- Except for riscv64 that isn't supported by Ubuntu 18.04
- Fix issue 7255 (1.6.7 can't be run on Ubuntu LTS 20.04 (GLIBC_2.34 not found))
- PR: https://github.com/containerd/containerd/pull/7260
- FILE: .github/workflows/release.yml
- FILE: .github/workflows/release/Dockerfile

### allow ptrace(2) by default for kernel >= 4.8
- FILE: contrib/seccomp/kernelversion/kernel_linux.go
- FILE: contrib/seccomp/seccomp_default.go
- PR: https://github.com/containerd/containerd/pull/7171

### Experimental CRI Sandbox server - Enable / Disable
- PR: https://github.com/containerd/containerd/pull/7169
- ENABLE_CRI_SANDBOXES
