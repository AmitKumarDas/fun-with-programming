### State of Performance Engineering 2021


### Snippet - Performance - Idiomatic - Fellow - Russ Cox
```yaml
- https://github.com/rsc/benchgraffiti/blob/master/havlak/havlak6.go
```

### Granular Control of Your Array's Growth
```go
// https://go.dev/blog/slices-intro

func AppendByte(slice []byte, data ...byte) []byte {
  m := len(slice)
  n := m + len(data)
  if n > cap(slice) { // if necessary, reallocate
    // allocate double what's needed, for future growth
    newSlice := make([]byte, (n+1)*2)
    copy(newSlice, slice)
    slice = newSlice
  }
  slice = slice[0:n] // reuse existing array with existing capacity & length = n
  copy(slice[m:n], data) // append data into exact positions
  return slice
}
```

### Reuse Arrays
#### AVOID
```go
// https://thanos.io/tip/contributing/coding-style-guide.md/

var messages []string
for _, msg := range recv {
  messages = append(messages, msg)

  if len(messages) > maxMessageLen {
    marshalAndSend(messages)
    // This creates new array
    //
    // Previous array will be garbage collected only after
    // some time (seconds), which can create enormous memory pressure
    messages = []string
  }
}
```
#### BETTER
```go
var messages []string
for _, msg := range recv {
  messages = append(messages, msg)

  if len(messages) > maxMessageLen {
    marshalAndSend(messages)
    // Instead of new array
    // reuse the existing array with same capacity
    // just length equals to zero
    messages = messages[:0]
  }
}
```


### Golang - Stack vs Heap
```go
// std library - go.go

var (
  // max size of the explicit variable that can get allocated to the stack
  // var x T
  // x := ...
  maxStackVarSize = int64(10 * 1024 * 1024)

  // max size of the implicit variable that can get allocated to the stack
  // p := new(T)
  // p := &T{}
  // s := make([]T, n)
  // s := []byte("blah")
  maxImplicitStackVarSize = int64(64 * 1024)
)
```

```go
// maxImplicitStackVarSize/t.Elem().width = 65536/8(int64) = 8192
// so, x will escape to heap
x := make([]int64, 8192)

// maxImplicitStackVarSize/t.Elem().width = 65536/1(byte) = 65536
// so, y will escape to heap
y := make([]byte, 65536)
```

```go
// the difference in memory size of slice is 1 byte
// but smaller slice stays in stack and other escapes to heap

const size = 64 * 1024 // 65536

func Benchmark_LargeSize_Stack_EqualOrLess65535(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // not escape to heap when size <= 65535
    dataLarge := make([]byte, size-1)
    _ = dataLarge
  }
}

func Benchmark_LargeSize_Heap_LargerThan65535(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // escape to heap when size > 65535
    dataLarge := make([]byte, size)
    _ = dataLarge
  }
}
```

### Golang - Stack vs Heap - 2
```yaml
- https://github.com/benhoyt/goawk/commit/af993094e3e8aca2b7ab709ffcda437996c906fe?diff=split
- Two tweaks to evalIndex

- one is to simply speed up the common case
- of a 1-dimensional index and avoid the loop entirely

- other is to allocate an array of fixed size (3) to begin with
- so the compiler can do that first allocation on the heap
```

```go
// heap
func (p *interp) evalIndex(indexExprs []Expr) (string, error) {
  indices := make([]string, len(indexExprs))
  for i, expr := range indexExprs {
  	v, err := p.eval(expr)
  	if err != nil {
  		return "", err
  	}
  	indices[i] = p.toString(v)
  }
  return strings.Join(indices, p.subscriptSep), nil
}
```

```go
// stack
func (p *interp) evalIndex(indexExprs []Expr) (string, error) {
	// Up to 3-dimensional indices won't require heap allocation
	indices := make([]string, 0, 3)
	for _, expr := range indexExprs {
		v, err := p.eval(expr)
		if err != nil {
			return "", err
		}
		indices = append(indices, p.toString(v))
	}
	return strings.Join(indices, p.subscriptSep), nil
}
```

```go
// stack
func (p *interp) evalIndex(indexExprs []Expr) (string, error) {
	// Optimize the common case of a 1-dimensional index
	if len(indexExprs) == 1 {
		v, err := p.eval(indexExprs[0])
		if err != nil {
			return "", err
		}
		return p.toString(v), nil
	}

	// Up to 3-dimensional indices won't require heap allocation
	indices := make([]string, 0, 3)
	for _, expr := range indexExprs {
		v, err := p.eval(expr)
		if err != nil {
			return "", err
		}
		indices = append(indices, p.toString(v))
	}
	return strings.Join(indices, p.subscriptSep), nil
}
```

### Golang Snippets - Error Handling
```yaml
- https://github.com/benhoyt/goawk/commit/aa6aa75368afeb40897b180c5a36501012e94907
- Initially code returned just the value and panic with a special error on runtime error
- But that was a significant slow-down
- So switched to using more verbose but more Go-like error return values
- This change gave a 2-3x improvement on a lot of benchmarks
```
```diff
- func (p *interp) evalSafe(expr Expr) (v value, err error) {
-   defer func() {
-     if r := recover(); r != nil {
-       // Convert to interpreter Error or re-panic
-       err = r.(*Error)
-     }
-   }()
-   return p.eval(expr), nil
- }
```

### Golang Snippets - Buffers
```yaml
- input and output is handled using std io.go
- I/O is buffered for efficiency
```

```diff
- output = os.Stdout
+ output = bufio.NewWriterSize(os.Stdout, 64*1024)
```

```diff
- errorOutput = os.Stderr
+ errorOutput = bufio.NewWriterSize(os.Stderr, 64*1024)
```

```go
// https://github.com/benhoyt/goawk/commit/6ba004f5fbf9b84bc6196d50c2a0dd496ed1771b

// Implement a buffered version of WriteCloser 
// so output is buffered when redirecting to a file 
// eg: print >"out"
type bufferedWriteCloser struct {
	*bufio.Writer
	io.Closer
}

func newBufferedWriteClose(w io.WriteCloser) *bufferedWriteCloser {
	writer := bufio.NewWriterSize(w, outputBufSize)
	return &bufferedWriteCloser{writer, w}
}

func (wc *bufferedWriteCloser) Close() error {
	err := wc.Writer.Flush()
	if err != nil {
		return err
	}
	return wc.Closer.Close()
}
```

### Golang - Switch Seems Better Than Map
```yaml
- https://github.com/benhoyt/goawk/commit/ad8ff0e5f6cc89fdd480099614187ee23b20a8c9
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

### Golang - CPU Profile
```yaml
- https://go.dev/blog/pprof

- If Go testing.B is used
- we can use gotest’s standard -cpuprofile and -memprofile flags
- In a standalone program - import runtime/pprof
```

### Golang - Memory Profile
```yaml
- https://go.dev/blog/pprof

- If program is spending most of its time allocating memory and garbage collecting
- I.e. runtime.mallocgc - it allocates and runs periodic garbage collections 
- Add memory profiling to the program
```

```yaml
- use go tool pprof exactly the same way
- Now the samples we are examining are memory allocations, not clock ticks

- To reduce overhead
- Memory profiler only records information for approximately ONE BLOCK per HALF MEGABYTE allocated 
- i.e. “1-in-524288 sampling rate”
- So these are approximations to the actual counts
- To find the memory allocations - list those functions
```

```yaml
- If we run go tool pprof with the --inuse_objects flag
- it will report ALLOCATION COUNTS instead of SIZES

- Instead of using a map, we can use a simple slice to list the elements
- In all but one of the cases where maps are being used
- It is impossible for the algorithm to insert a duplicate element
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

```yaml
- Writing idiomatic Go style, using data structures and methods 
- Does not make your program slow
```

```yaml
- https://github.com/rsc/benchgraffiti/blob/master/havlak/havlak6.go
- Idiomatic - Performant - Fellow
```

### Golang - Slice Better Than Map
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

### Golang - Slice Better Than Map - 2
```yaml
- https://github.com/benhoyt/goawk/commit/e0d7287ac1580bd0f144c763b222b9db8a858c54
- Resolve variable names to indexes at parse time
```

```diff
type VarExpr struct {
-   Name string
+   Index int
+   Name  string
}
```
```diff
type GetlineExpr struct {
-   Command Expr
-   Var     string
-   File    Expr
+   Command  Expr
+   VarIndex int
+   VarName  string
+   File     Expr
}
```

```go
const (
  V_ILLEGAL = iota
  V_ARGC
  V_SUBSEP

  V_LAST = V_SUBSEP
)

var Vars = map[string]int{
  "ARGC":     V_ARGC,
  "SUBSEP":   V_SUBSEP,
}

func VarIndex(name string) int {
  return Vars[name]
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

### Blogs | References
```yaml
- https://prometheus.io/blog/2019/10/10/remote-read-meets-streaming/
- https://github.com/google/cadvisor
- https://github.com/cloudflare/ebpf_exporter

- https://www.brendangregg.com/blog/2021-09-26/the-speed-of-time.html
- service restarting in loop // can be invisible in top(8)

- https://benhoyt.com/writings/goawk/
- golang
```

### Tools
```yaml
- https://strace.io/
- attach to an already running process
```

### Repo - QOI - 1
```yaml
- https://github.com/xfmoulet/qoi
```

### Repo - QOI - 2
```yaml
- https://github.com/MasterQ32/zig-qoi
```
