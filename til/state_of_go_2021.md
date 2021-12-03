### State of Golang Code 2021

### Best Practices Guide
#### Thanos
```yaml
- https://thanos.io/tip/contributing/coding-style-guide.md/

- variable shadowing: avoid
- package name shadowing: avoid
- defer error: handle
- exhaust the readers: til
- no globals other than const are allowed: Hence, no init functions
- never use panic: avoid dependencies who use it
- reflect is very slow: avoid
- preallocate slices & maps
- reuse arrarys
- shallow functions: avoid
- inlining improves readability: less cognitive load to readers
- there should be one and preferably only one obvious way to do it
- avoid defining variables that are used only once
```

#### Uber
```yaml
- https://github.com/uber-go/guide/blob/master/style.md
```

#### Comments
```yaml
- https://buttondown.email/hillelwayne/archive/comment-the-why-and-the-what/
```

### Type Cast to an Interface -- Just In Time
```go
func Unwrap(err error) error {
  u, ok := err.(interface {
    Unwrap() error
  })
  if !ok {
    return nil
  }
  return u.Unwrap()
}
```

### New Is Not Same As Make
```yaml
- they do different things
- they apply to different types
```

#### new
```yaml
- a built-in function that ALLOCATES memory
- it does NOT INITIALIZE the memory
- it only zeros it
- i.e. new(T) Allocates Zeroed Storage for a new item of type T
- & returns its address, a value of type *T
- In Go, it returns a Pointer to a Newly Allocated Zero Value of type T

- Since the memory returned by new is Zeroed
- it's helpful to design your data structures
- s.t the zero value of each type can be used without further initialization
- This means a user of the data structure can create one with new and get right to work

- E.g, the documentation for bytes.Buffer states:
- the zero value for Buffer is an empty buffer ready to use
- Similarly, sync.Mutex does not have an explicit constructor or Init method
- Instead, the zero value for a sync.Mutex is defined to be an unlocked mutex
```

#### make
```yaml
- make(T, args) serves a purpose different from new(T)
- it is applicable for slices, maps, and channels only
- it returns an Initialized (not zeroed) value of type T (not *T)

- The reason for the distinction is that these 3 types represent, under the covers:
- references to data structures that Must be Initialized before use

- E.g slice is a three-item descriptor:
- a pointer to the data
- the length &
- the capacity
- until those items are initialized, the slice is nil

- make([]int, 10, 100)
- allocates an array of 100 ints
- then creates a slice structure with length 10 &
- a capacity of 100 Pointing at the First 10 elements of the array
```

#### new([]int)
```yaml
- returns a pointer to a newly allocated
- zeroed slice structure
- i.e. a pointer to a nil slice value
```

#### Conclusion
```yaml
- make vs new:
- Initialized Value vs Zeroed Allocation
```

```go
// To obtain an explicit pointer:
// - allocate with new or
// - take the address of a variable explicitly

var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```


### Zero Value Just Works
```go
type SyncedBuffer struct {
  lock    sync.Mutex
  buffer  bytes.Buffer
}
```
```yaml
- Values of type SyncedBuffer are also ready to use immediately:
- upon allocation or
- just via declaration
```
```go
// both p and v will work correctly without further arrangement

p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

### Null | JSON
```yaml
- https://github.com/kubernetes/kubernetes/pull/104990/
```

### Channel | Signal | Shutdown
```yaml
- https://rudderstack.com/blog/implementing-graceful-shutdown-in-go/
```

### JSON | Get | Performance | E2E | Assert
```yaml
- https://github.com/bhmj/jsonslice
```

### Design | Open API | Kubernetes | Validation
```yaml
- https://danielmangum.com/posts/how-kubernetes-validates-custom-resources/
```

### Learn
```yaml
- https://github.com/toni-moreno/syncflux
- DB - InfluxDB - learn design

- https://github.com/mojura/mojura
- todo - learn design

- https://github.com/stefanprodan/kustomizer/blob/main/pkg/objectutil/io.go
- yaml to unstruct - unstruct to yaml

- https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/deploy/manifests/custom-metrics-apiservice.yaml
- api - custom - kubernetes
- k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1
- k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/typed/apiregistration/v1beta1

- https://github.com/kubernetes/kubernetes/tree/master/pkg/proxy/userspace
- load balancer - proxy - round robin - limit - socket
```

### Fuzz Testing
```yaml
- https://github.com/dvyukov/go-fuzz
- ensures no panic, crash, allocate insane amount of memory, nor hang
```

```yaml
- https://blog.cloudflare.com/dns-parser-meet-go-fuzzer/
```

### Buffers - Zeroing
```go
// zeroBuf is a big buffer of zero bytes, used to zero out the buffers
var zeroBuf = make([]byte, 65535)

var bufpool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 2048)
    },
}

// in some function
data := bufpool.Get().([]byte)
defer bufpool.Put(data)
copy(data[0:cap(data)], zeroBuf)
```

### Snippets - API Design
```go
// design thinking w.r.t e2e, testing, assertion

import "k8s.io/apimachinery/pkg/util/sets"

type Strings []string

// some callers might need the resulting bool
func (s Strings) HasAll(expected []string) bool {
	actualSet := sets.NewString(s...)
	return actualSet.HasAll(expected...)
}

// other callers might need the resulting error
//
// Notes: 
// - Don't pass *testing.T as an argument
// - It might be a separate struct that needs *testing.T as first argument
func (s Strings) EnsureHasAll(expected []string) error {
	if !s.HasAll(expected) {
		actualSet := sets.NewString(s...)
		expectedSet := sets.NewString(expected...)
		diff := expectedSet.Difference(actualSet)
		return errors.Errorf("missing : %v: missing count %d:", diff.List(), diff.Len())
	}
	return nil
}
```

### Snippets - Sockets
```go
// Node proxy via userspace socket
//
// Note: Check how socket level eBPF does this in 30 lines of code with best performance
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
//
// For each connection from a local client to a ClusterIP:Port
// intercept the connection and split it into two separate connections:
//
// connection 1: local client <---> toy-proxy
// connection 2: toy-proxy <---> backend pods
// The easiest way to achieve this is to implement it in userspace:
//
// Listen to resources:
// 1/ start a daemon process
// 2/ listen to K8S apiserver
// 3/ watch Service (ClusterIP) and Endpoint (Pod) changes
//
// Proxy traffic:
// 1/ for each connecting request from a local client to a Service (ClusterIP)
// 2/ intercept the request by acting as a middleman
//
// Dynamically apply proxy rules
// 1/ for any Service/Endpoint updates
// 2/ change toy-proxy connection settings accordingly
//
// ClusterIP didn't reside on on any network device of this node
// which means we could not do something like listen(ClusterIP, Port).
//
// Following command will redirect all traffic for ClusterIP:Port to localhost:Port
//
// $ sudo iptables -t nat -A OUTPUT -p tcp -d $CLUSTER_IP --dport $PORT -j REDIRECT --to-port $PORT
//
// $ iptables -t nat -L -n
// ...
// Chain OUTPUT (policy ACCEPT)
// target     prot opt source      destination
// REDIRECT   tcp  --  0.0.0.0/0   10.7.111.132         tcp dpt:80 redir ports 80

func main() {
	clusterIP := "10.7.111.132"
	podIP := "10.5.41.204"
	port := 80
	proto := "tcp"

	addRedirectRules(clusterIP, port, proto)
	createProxy(podIP, port, proto)
}

func addRedirectRules(clusterIP string, port int, proto string) error {
	p := strconv.Itoa(port)
	cmd := exec.Command("iptables", "-t", "nat", "-A", "OUTPUT", "-p", "tcp",
		"-d", clusterIP, "--dport", p, "-j", "REDIRECT", "--to-port", p)
	return cmd.Run()
}

// creates the userspace proxy, and maintains bi-directional forwarding
func createProxy(podIP string, port int, proto string) {
	host := ""
	listener, err := net.Listen(proto, net.JoinHostPort(host, strconv.Itoa(port)))

	for {
		inConn, err := listener.Accept()
		outConn, err := net.Dial(proto, net.JoinHostPort(podIP, strconv.Itoa(port)))

		go func(in, out *net.TCPConn) {
			var wg sync.WaitGroup
			wg.Add(2)
			fmt.Printf("Proxying %v <-> %v <-> %v <-> %v\n",
				in.RemoteAddr(), in.LocalAddr(), out.LocalAddr(), out.RemoteAddr())
			go copyBytes(in, out, &wg)
			go copyBytes(out, in, &wg)
			wg.Wait()
		}(inConn.(*net.TCPConn), outConn.(*net.TCPConn))
	}

	listener.Close()
}

func copyBytes(dst, src *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	if _, err := io.Copy(dst, src); err != nil {
		if !strings.HasSuffix(err.Error(), "use of closed network connection") {
			fmt.Printf("io.Copy error: %v", err)
		}
	}
	dst.Close()
	src.Close()
}
```



### TODO
- https://blog.afoolishmanifesto.com/posts/writing-a-golang-linter/
- https://blog.afoolishmanifesto.com/posts/benefits-using-golang-adhoc-code-leatherman/
