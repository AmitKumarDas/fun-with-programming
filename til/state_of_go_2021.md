### State of Golang Code 2021
```yaml
- https://rudderstack.com/blog/implementing-graceful-shutdown-in-go/
- context - channel - goroutine - signal - shutdown

- https://github.com/toni-moreno/syncflux
- DB - InfluxDB - learn design

- https://github.com/mojura/mojura
- todo - learn design

- https://github.com/bhmj/jsonslice
- performance - query - raw - []byte - e2e - kubernetes - assert

- how to go get any private github repo?
- set below in ~/.profile
- machine github.com username AmitKumarDas password <my github token with relevant access>

- https://danielmangum.com/posts/how-kubernetes-validates-custom-resources/
- design - history - kubernetes - open api - validation - api

- https://github.blog/2021-11-10-make-your-monorepo-feel-small-with-gits-sparse-index/
- monorepo - git - sparse index - contributing guide

- https://github.com/stefanprodan/kustomizer/blob/main/pkg/objectutil/io.go
- yaml to unstruct - unstruct to yaml

- https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/deploy/manifests/custom-metrics-apiservice.yaml
- api - custom - kubernetes
- k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1
- k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/typed/apiregistration/v1beta1

- https://github.com/kubernetes/kubernetes/tree/master/pkg/proxy/userspace
- load balancer - proxy - round robin - limit - socket
```

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
