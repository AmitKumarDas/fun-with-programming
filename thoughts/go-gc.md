## Learn You Some Go Garbage Collection

```yaml
- The Go garbage collector runs in the background as its own Goroutine
- In fact it's several Goroutines
- The garbage collector can be triggered manually by calling runtime.GC()
- But usually it runs automatically when the heap DOUBLES its size
- This size threshold can be adjusted with the GOGC environment variable
- It is set to a percentage
- The default is 100, meaning the heap has to grow by 100% to trigger the garbage collection
```

```yaml
- Setting it to 200
- Would mean that the collection is only started when the heap has grown to THREE times the previous size
```

```yaml
- On top of the size condition there is also a TIMING condition
- As long as the process is not suspended, the garbage collector will run at least ONCE every TWO minutes
```

### References
```yaml
- https://dev.to/jlauinger/sliceheader-literals-in-go-create-a-gc-race-and-flawed-escape-analysis-exploitation-with-unsafe-pointer-on-real-world-code-4mh7
```