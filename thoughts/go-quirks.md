## Surprising Yet True

### go.mod
```yaml
- If my go.mod says go 1.18 
- But I have a file with a constraint //go:build !go1.18
- That file will today only be used when building with Go 1.17 and lower
- So it would be at least incoherent to use a Go 1.18 feature in that file
```
```yaml
- If my go.mod says go 1.18
- But I have a file with a constraint //go:build go1.19
- That file will be used only when building with Go 1.19 and above
- Even though the module overall uses Go 1.18 language semantics
```
