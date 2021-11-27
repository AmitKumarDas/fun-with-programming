### State of Performance Engineering 2021

### Golang
```yaml
- input and output is handled using std io.go

- I/O is buffered for efficiency
- use bufio.Scanner to read input
- use bufio.Writer to buffer output
```

### Golang - Fuzz Testing
```yaml
- https://blog.cloudflare.com/dns-parser-meet-go-fuzzer/
- RRDNS in-house DNS server
- uses github.com/miekgs/dns for all its parsing needs
- panicks on malformed packets
- this is Go, not C, and we can afford to recover() panics 
- no worrying about ending up with insane memory states
```

```go
func ParseDNSPacketSafely(buf []byte, msg *old.Msg) (err error) {
  defer func() {
    panicked := recover()
    if panicked != nil {
      err = errors.New("ParseError")
    }
  }()

  err = msg.Unpack(buf)
  return
}
```

### Golang | CPU profile
```yaml
- If Go testing.B is used
- we can use gotest’s standard -cpuprofile and -memprofile flags
- In a standalone program - import runtime/pprof
```
```go
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
  flag.Parse()
  if *cpuprofile != "" {
    f, err := os.Create(*cpuprofile)
    if err != nil {
      log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
  }
```
```yaml
- run program with -cpuprofile flag 
- then run go tool pprof to interpret the profile

- $ make havlak1.prof
- ./havlak1 -cpuprofile=havlak1.prof

- $ go tool pprof havlak1 havlak1.prof
```
```yaml
- go tool pprof is a slight variant of Google’s pprof C++ profiler
- important command is topN
- which shows the top N samples in the profile

- (pprof) top10
- Total: 2525 samples
-  298  11.8%  11.8%      345  13.7% runtime.mapaccess1_fast64
-  268  10.6%  22.4%     2124  84.1% main.FindLoops
-  251   9.9%  32.4%      451  17.9% scanblock
```

### Golang | Memory Profile
```yaml
- If program is spending most of its time allocating memory and garbage collecting
- I.e. runtime.mallocgc - it allocates and runs periodic garbage collections 
- Add memory profiling to the program
- program stops after one iteration of the loop finding, writes a memory profile, and exits
```
```go
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

if *memprofile != "" {
  f, err := os.Create(*memprofile)
  if err != nil {
    log.Fatal(err)
  }
  pprof.WriteHeapProfile(f)
  f.Close()
  return
}
```
```yaml
- $ make havlak3.mprof
go build havlak3.go
./havlak3 -memprofile=havlak3.mprof
```
```yaml
- use go tool pprof exactly the same way
- Now the samples we are examining are memory allocations, not clock ticks

- $ go tool pprof havlak3 havlak3.mprof
- Adjusting heap profiles for 1-in-524288 sampling rate

- (pprof) top5
- Total: 82.4 MB
-    56.3  68.4%  68.4%     56.3  68.4% main.FindLoops
-    17.6  21.3%  89.7%     17.6  21.3% main.(*CFG).CreateNode
-     8.0   9.7%  99.4%     25.6  31.0% main.NewBasicBlockEdge
-     0.5   0.6% 100.0%      0.5   0.6% itab
-     0.0   0.0% 100.0%      0.5   0.6% fmt.init
- (pprof)
```
```yaml
- FindLoops allocated approximately 56.3 of the 82.4 MB in use
- CreateNode accounts for another 17.6 MB

- To reduce overhead
- memory profiler only records information for approximately ONE BLOCK per HALF MEGABYTE allocated 
- i.e. “1-in-524288 sampling rate”
- so these are approximations to the actual counts

- To find the memory allocations - list those functions

- (pprof) list FindLoops
- Total: 82.4 MB
- ROUTINE ====================== main.FindLoops in /home/rsc/g/benchgraffiti/havlak/havlak3.go
-  56.3   56.3 Total MB (flat / cumulative)

   1.9    1.9  268:     nonBackPreds := make([]map[int]bool, size)
   5.8    5.8  269:     backPreds := make([][]int, size)
     .      .  270:
   1.9    1.9  271:     number := make([]int, size)
   1.9    1.9  272:     header := make([]int, size, size)
   1.9    1.9  273:     types := make([]int, size, size)
   1.9    1.9  274:     last := make([]int, size, size)
   1.9    1.9  275:     nodes := make([]*UnionFindNode, size, size)
     .      .  276:
     .      .  277:     for i := 0; i < size; i++ {
   9.5    9.5  278:             nodes[i] = new(UnionFindNode)
     .      .  279:     }
...
     .      .  286:     for i, bb := range cfgraph.Blocks {
     .      .  287:             number[bb.Name] = unvisited
  29.5   29.5  288:             nonBackPreds[i] = make(map[int]bool)
     .      .  289:     }
...
- It looks like the current bottleneck is the same as the last one
- using maps where simpler data structures suffice
- FindLoops is allocating about 29.5 MB of maps
```

```yaml
- If we run go tool pprof with the --inuse_objects flag
- it will report ALLOCATION COUNTS instead of SIZES

- $ go tool pprof --inuse_objects havlak3 havlak3.mprof
- Adjusting heap profiles for 1-in-524288 sampling rate

- (pprof) list FindLoops
- Total: 1763108 objects
- ROUTINE ====================== main.FindLoops in /home/rsc/g/benchgraffiti/havlak/havlak3.go
- 720903 720903 Total objects (flat / cumulative)

-      .      .  277:     for i := 0; i < size; i++ {
- 311296 311296  278:             nodes[i] = new(UnionFindNode)
-      .      .  279:     }
-      .      .  280:
-      .      .  281:     // Step a:
-      .      .  282:     //   - initialize all nodes as unvisited.
-      .      .  283:     //   - depth-first traversal and numbering.
-      .      .  284:     //   - unreached BB's are marked as dead.
-      .      .  285:     //
-      .      .  286:     for i, bb := range cfgraph.Blocks {
-      .      .  287:             number[bb.Name] = unvisited
- 409600 409600  288:             nonBackPreds[i] = make(map[int]bool)
-      .      .  289:     }

- Since the ~200,000 maps account for 29.5 MB
- it looks like the initial map allocation takes about 150 bytes
- That’s reasonable when a map is being used to hold key-value pairs
- but not when a map is being used as a stand-in for a simple set, as it is here

- Instead of using a map, we can use a simple slice to list the elements
- In all but one of the cases where maps are being used
- it is impossible for the algorithm to insert a duplicate element
- In the one remaining case, write a simple variant of the append
```
```go
func appendUnique(a []int, x int) []int {
  for _, y := range a {
    if x == y {
      return a
    }
  }
  return append(a, x)
}
```

### Golang - Slice Better Than Map If You Can
```yaml
- Instead of map[*BasicBlock]int 
- use a []int
- a slice indexed by the block number
- no reason to use a map when an array or slice will do
```
```go
type BasicBlock struct {
  Name int // use this as the index of the slice
}
```

### Assembly | Compiler | Instructions | Interpreter
```yaml
- see the CPU instructions generated by the Go compiler
- need a disassembler 
- e.g. objdump which comes with the GNU binutils - may already be installed if Linux

- push instructions to add values onto the stack vs. placement onto the stack using mov
- performance - mov requires fewer CPU cycles than a push

- https://benhoyt.com/writings/goawk/
- awk interpreter in go

- https://blog.cloudflare.com/automatically-generated-types/
- workers - rust - typescript - intermediate representation
- abstract syntax tree
```

### Blogs
```yaml
- https://prometheus.io/blog/2019/10/10/remote-read-meets-streaming/
- https://github.com/google/cadvisor
- https://github.com/cloudflare/ebpf_exporter


- https://www.brendangregg.com/blog/2021-09-26/the-speed-of-time.html
- service restarting in loop // can be invisible in top(8)
```

### Tools
```yaml
- https://strace.io/
- attach to an already running process
```
