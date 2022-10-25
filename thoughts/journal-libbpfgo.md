- GIVEN: `git clone git@github.com:aquasecurity/libbpfgo.git`
- GIVEN: uname
```shell
Darwin
```
- THEN: git submodule init
- THEN: git submodule update
- THEN: make libbpfgo-static
```shell
mkdir -p ./output
/bin/bash: /bin/false: No such file or directory
/bin/bash: /bin/false: No such file or directory
ERROR: kernel does not seem to support BTF
make: *** [output/vmlinux.h] Error 1
```
- DETOUR: Read `makefile-101-libbpfgo.md` first
- THEN: `make vagrant-up`
```shell
VAGRANT_VAGRANTFILE=/Users/amitd2/work/libbpfgo/builder/Vagrantfile-ubuntu \
		ARCH=arm64 \
		HOSTOS=Darwin \
		vagrant up
/bin/bash: vagrant: command not found
make: *** [.vagrant-up] Error 127
```
- THEN: `brew install vagrant`
- THEN: `make vagrant-up`
```shell
VAGRANT_VAGRANTFILE=/Users/amitd2/work/libbpfgo/builder/Vagrantfile-ubuntu \
		ARCH=arm64 \
		HOSTOS=Darwin \
		vagrant up
No usable default provider could be found for your system.

Vagrant relies on interactions with 3rd party systems, known as
"providers", to provide Vagrant with resources to run development
environments. Examples are VirtualBox, VMware, Hyper-V.

The easiest solution to this message is to install VirtualBox, which
is available for free on all major platforms.

If you believe you already have a provider available, make sure it
is properly installed and configured. You can see more details about
why a particular provider isn't working by forcing usage with
`vagrant up --provider=PROVIDER`, which should give you a more specific
error message for that particular provider.
make: *** [.vagrant-up] Error 1
```
- THEN: `brew install VirtualBox`
```shell
Running `brew update --auto-update`...
==> Auto-updated Homebrew!
Updated 1 tap (homebrew/core).

==> Downloading https://download.virtualbox.org/virtualbox/7.0.0/VirtualBox-7.0.0-153978-OSX.dmg
######################################################################## 100.0%
Error: Cask virtualbox depends on hardware architecture being one of [{:type=>:intel, :bits=>64}], but you are running {:type=>:arm, :bits=>64}.
```
- THEN: `brew install --cask parallels`
- THEN: `vagrant plugin install vagrant-parallels`
- THEN: `make vagrant-up`
```shell
VAGRANT_VAGRANTFILE=/Users/amitd2/work/libbpfgo/builder/Vagrantfile-ubuntu \
		ARCH=arm64 \
		HOSTOS=Darwin \
		vagrant up
No usable default provider could be found for your system.

Vagrant relies on interactions with 3rd party systems, known as
"providers", to provide Vagrant with resources to run development
environments. Examples are VirtualBox, VMware, Hyper-V.

The easiest solution to this message is to install VirtualBox, which
is available for free on all major platforms.

If you believe you already have a provider available, make sure it
is properly installed and configured. You can see more details about
why a particular provider isn't working by forcing usage with
`vagrant up --provider=PROVIDER`, which should give you a more specific
error message for that particular provider.
make: *** [.vagrant-up] Error 1
```
- DETOUR: `journal-lima.md`
- THEN: `limactl shell ebpf`
- THEN: `make libbpfgo-static`
```shell
INFO: generating ./output/vmlinux.h from /sys/kernel/btf/vmlinux
WARNING: bpftool not found for kernel 5.15.0-52

  You may need to install the following packages for this specific kernel:
    linux-tools-5.15.0-52-generic
    linux-cloud-tools-5.15.0-52-generic

  You may also want to install one of the following packages to keep up to date:
    linux-tools-generic
    linux-cloud-tools-generic
ERROR: could not create ./output/vmlinux.h
make: *** [Makefile:101: output/vmlinux.h] Error 1
```
- THEN: `sh builder/prepare-ubuntu.sh`
```shell
INFO: coreutils installed
INFO: bsdutils installed
INFO: findutils installed
INFO: build-essential installed
INFO: pkgconf installed
INFO: golang-1.18-go installed
INFO: llvm-12 installed
INFO: clang-12 installed
INFO: clang-format-12 installed
INFO: linux-headers-generic installed
INFO: linux-tools-generic installed
INFO: linux-tools-5.15.0-52-generic installed
INFO: zlib1g-dev installed
INFO: libelf-dev installed
INFO: libbpf-dev installed
INFO: Go 1.18 set as default
INFO: Clang 12 set as default
```
- THEN: `make libbpfgo-static`
```shell
INFO: generating ./output/vmlinux.h from /sys/kernel/btf/vmlinux
mkdir -p ./output/libbpf
CC="gcc" CFLAGS="-g -O2 -Wall -fpie" LD_FLAGS="" \
   make -C /Users/amitd2/work/libbpfgo/libbpf/src \
	BUILD_STATIC_ONLY=1 \
	OBJDIR=/Users/amitd2/work/libbpfgo/output/libbpf \
	DESTDIR=/Users/amitd2/work/libbpfgo/output \
	INCLUDEDIR= LIBDIR= UAPIDIR= install
make[1]: Entering directory '/Users/amitd2/work/libbpfgo/libbpf/src'
  MKDIR    /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/bpf.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/btf.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/libbpf.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/libbpf_errno.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/netlink.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/nlattr.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/str_error.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/libbpf_probes.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/bpf_prog_linfo.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/btf_dump.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/hashmap.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/ringbuf.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/strset.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/linker.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/gen_loader.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/relo_core.o
  CC       /Users/amitd2/work/libbpfgo/output/libbpf/staticobjs/usdt.o
  AR       /Users/amitd2/work/libbpfgo/output/libbpf/libbpf.a
  INSTALL  bpf.h libbpf.h btf.h libbpf_common.h libbpf_legacy.h bpf_helpers.h bpf_helper_defs.h bpf_tracing.h bpf_endian.h bpf_core_read.h skel_internal.h libbpf_version.h usdt.bpf.h
  INSTALL  /Users/amitd2/work/libbpfgo/output/libbpf/libbpf.pc
  INSTALL  /Users/amitd2/work/libbpfgo/output/libbpf/libbpf.a 
make[1]: Leaving directory '/Users/amitd2/work/libbpfgo/libbpf/src'
CC=clang \
	CGO_CFLAGS="-I/Users/amitd2/work/libbpfgo/output" \
	CGO_LDFLAGS="-lelf -lz /Users/amitd2/work/libbpfgo/output/libbpf.a" \
	GOOS=linux GOARCH=arm64 \
	go build \
	-tags netgo -ldflags '-w -extldflags "-static"' \
```
- CHECK: `ls -lrt output/`
```shell
total 6084
-rw-r--r-- 1 amitd2 dialout 3404492 Oct 19 11:01 vmlinux.h
drwxr-xr-x 1 amitd2 dialout      96 Oct 19 11:01 pkgconfig
-rw-r--r-- 1 amitd2 dialout 2808164 Oct 19 11:01 libbpf.a
drwxr-xr-x 1 amitd2 dialout     160 Oct 19 11:01 libbpf
drwxr-xr-x 1 amitd2 dialout     480 Oct 19 11:01 bpf
```
