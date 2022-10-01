## Learning string in Golang

```yaml
- Strings in go are mostly read only []byte slice
- However, range iterates over runes instead of bytes
- Where rune is a unicode code point
- Since strings in go are encoded in UTF-8
- Hence a unicode code point i.e. rune might need more than one bytes
```

```yaml
- Strings are read-only
- Because the string literal (its data) is placed in a SPECIAL section in the BINARY file produced by the compiler
- When the program is run, this section is probably mapped into a read-only memory page
- Therefore, the Data field in the StringHeader and SliceHeader structures will contain an address inside that read-only page
```

### References
```yaml
- https://dev.to/jlauinger/sliceheader-literals-in-go-create-a-gc-race-and-flawed-escape-analysis-exploitation-with-unsafe-pointer-on-real-world-code-4mh7
```