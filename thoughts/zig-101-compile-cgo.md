## Compiling with zig

### Building SQLite with CGo for (almost) every OS
- START:
```shell
$ CGO_ENABLED=1 CC="zig cc" CXX="zig cc" go test -c
$ ./go-sqlite3.test
PASS
```
- NOTES:
```yaml
- Normally you would just run go test
- Add -c to make Go produce the test executable without running it
- Allowing us to move the executable on the correct machine
```
- THEN: `ldd go-sqlite.test`
```shell
linux-vdso.so.1 (0x0000ffffa4840000)
libpthread.so.0 => /nix/store/34k9b4lsmr7mcmykvbmwjazydwnfkckk-glibc-2.33-50/lib/libpthread.so.0 (0x0000ffffa47e0000)
libc.so.6 => /nix/store/34k9b4lsmr7mcmykvbmwjazydwnfkckk-glibc-2.33-50/lib/libc.so.6 (0x0000ffffa466d000)
libdl.so.2 => /nix/store/34k9b4lsmr7mcmykvbmwjazydwnfkckk-glibc-2.33-50/lib/libdl.so.2 (0x0000ffffa4659000)
/nix/store/34k9b4lsmr7mcmykvbmwjazydwnfkckk-glibc-2.33-50/lib/ld-linux-aarch64.so.1 (0x0000ffffa480e000)
```
- FYI: Above is a dynamic executable
- THEN: Make a static executable
```shell
CGO_ENABLED=1 \
  CC="zig cc -target native-native-musl" \
  CXX="zig cc -target native-native-musl" \
  go test -c

./go-sqlite3.test
PASS

ldd go-sqlite3.test
    statically linked
```
- THEN: Build for Windows x86_64
```shell
CGO_ENABLED=1 \
  GOOS=windows \
  GOARCH=amd64 \
  CC="zig cc -target x86_64-windows" \
  CXX="zig cc -target x86_64-windows" \
  go test -c 
```


## Other Notes
```yaml
- Zig accepts option to specify the target architecture - the libc ABI
- For Windows, you want gnu (e.g. x86_64-windows-gnu) because that will use Zig's bundled MinGW-w64
- For Linux, you probably want musl (e.g. x86_64-linux-musl) because your resulting binary will be statically linked
  - And thus work on all Linux distributions
- However, if you prefer to interact with the system glibc, such as on Ubuntu, you can specify gnu (e.g. x86_64-linux-gnu)
- Zig supports the following targets:
  - https://ziglang.org/documentation/master/#Targets
```

## References
```yaml
- https://dev.to/kristoff/zig-makes-go-cross-compilation-just-work-29ho
- https://zig.news/kristoff/building-sqlite-with-cgo-for-every-os-4cic
```
