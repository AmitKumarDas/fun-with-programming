
- BEGIN:
```yaml
- Often, when cross-compiling, it is useful to make a static binary
- In the case of Linux, this will make the resulting binary able to run on any Linux distribution
- Rather than only ones with a hard-coded glibc dynamic linker e.g. /lib/ld-linux-aarch64.so.1
- We can accomplish this by targeting musl rather than glibc
```
- NOTES:
```yaml

```

## References
```yaml
- https://andrewkelley.me/post/zig-cc-powerful-drop-in-replacement-gcc-clang.html
```