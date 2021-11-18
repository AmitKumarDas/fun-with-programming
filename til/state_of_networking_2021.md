### State of Networking & SMI 2021

```
// service broker // k8s // proxy
https://github.com/Peripli/service-broker-proxy-k8s
```

```
// startup idea
// media service // http proxy
// image transformations // resize // crop // overlays // blurring // sharpening
https://github.com/zalando-stups/skrop
```

```c
// tracing // plugins // opentracing // filters // plugins
// lua script extensions
//
// idea - plugins // startup // ecosystem around plugins
// idea - starlark plugins
//
https://github.com/skipper-plugins/
```

```c
// fellow of ideas
//
// http router // k8s // reverse proxy
// service composition // ingress // route identifier // lookup //
// request flow // filter // ingress without reloads // predicates
//
// oauth proxy // basic auth // verify clients // verify requests
// audit logging // rate limiters
// WAF - Web Application Firewall - filter for skipper routes // startup // idea
//
// startup idea - compare vs ingress controllers // compare vs service-mesh // differentiate
//
// startup idea - go module proxy for analytics
// startup idea - image registry proxy for analytics
// startup idea - image registry proxy + trow + mayastor // openebs
// startup idea - image server - https://github.com/zalando-stups/skrop
//
https://github.com/zalando/skipper

```

```c
https://www.gilesthomas.com/2021/03/fun-with-network-namespaces

// Namespaces and cgroups are two of the main kernel technologies
// cgroups are a metering and limiting mechanism, they control how much of a
// system resource (CPU, memory) you can use.
//
// On the other hand, namespaces limit what you can see. Thanks to namespaces
// processes have their own view of the system’s resources.
//
// Linux kernel provides 6 types of namespaces:
// pid, net, mnt, uts, ipc and user.
//
// For instance, a process inside a pid namespace only sees processes in the
// same namespace.
//
// Thanks to the mnt namespace, it’s possible to attach a process to its own
// filesystem (like chroot).
//
// $ ip netns add ns1
// This command will create a new network namespace called ns1.
// the ip command adds a bind mount point for it under /var/run/netns
// This allows the namespace to persist even if there’s no process attached to it
// To list the namespaces available in the system:
//
// $ ls /var/run/netns     // cmd
// ns1
//
// Or via ip
// $ ip netns              // cmd
// ns1
//
- https://blogs.igalia.com/dpino/2016/04/10/network-namespaces/
```

```c
// kernel // container // network namespaces // samples // how to // examples
//
// if you put a process into a network namespace, it will have its own restricted
// view of what the networking environment looks like -- it won't see the
// machine's main network interface,
//
// two processes inside different namespaces would have different networking
// environments, they could both bind to the same port -- and then could be
// accessed from outside via port forwarding.
//
// two Flask servers on the same machine, both bound to port 8080 inside their
// own namespace. I wanted to be able to access one of them from outside by
// hitting port 6000 on the machine, and the other by hitting port 6001.
//
// # ip netns add netns1                        // cmd
//
// # ip netns exec netns1 /bin/bash             // cmd
// # ip a
// 1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
//     link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
// # exit
//
// So, we have a new namespace, and when we're inside it, there's only one
// interface available, a basic loopback interface. // [fellow]
// We can compare that with what we see with the same command outside:
//
// # ip a                                       // cmd
// 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
//     link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
//     inet 127.0.0.1/8 scope host lo
//        valid_lft forever preferred_lft forever
//     inet6 ::1/128 scope host
//        valid_lft forever preferred_lft forever
// 2: ens5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 9001 qdisc mq state UP group default qlen 1000
//     link/ether 0a:d6:01:7e:06:5b brd ff:ff:ff:ff:ff:ff
//     inet 10.0.0.173/24 brd 10.0.0.255 scope global dynamic ens5
//        valid_lft 2802sec preferred_lft 2802sec
//     inet6 fe80::8d6:1ff:fe7e:65b/64 scope link
//        valid_lft forever preferred_lft forever
//
// actual network card attached to the machine with the name ens5 // [fellow]
//
// details shown for the loopback interface inside the namespace were much shorter
// -- no IPv4 or IPv6 addresses, for example.
// That's because the interface is down by default.  // [fellow]
// Let's see if we can fix that:
//
// # ip netns exec netns1 /bin/bash            // cmd
// # ip a
// 1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
//     link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
// # ping 127.0.0.1
// ping: connect: Network is unreachable
// # ip link set dev lo up                    // cmd
// # ip a
// 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
//     link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
//     inet 127.0.0.1/8 scope host lo
//        valid_lft forever preferred_lft forever
//     inet6 ::1/128 scope host
//        valid_lft forever preferred_lft forever
// # ping 127.0.0.1
// PING 127.0.0.1 (127.0.0.1) 56(84) bytes of data.
// 64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.019 ms
// 64 bytes from 127.0.0.1: icmp_seq=2 ttl=64 time=0.027 ms
// ^C
// --- 127.0.0.1 ping statistics ---
// 2 packets transmitted, 2 received, 0% packet loss, time 1022ms
// rtt min/avg/max/mdev = 0.019/0.023/0.027/0.004 ms
//
// but the external network still is down:
// # ip netns exec netns1 /bin/bash
// # ping 8.8.8.8
// ping: connect: Network is unreachable
//
// Again, that makes sense. There's no non-loopback interface, so there's
// no way to send packets to anywhere but the loopback network.
//
// Virtual network interfaces: connecting the namespace
//
// What we need is some kind of non-loopback network interface inside the
// namespace. However, we can't just put the external interface ens5 inside there;
// an interface can only be in one namespace at a time, so if we put that one
// in there, the external machine would lose networking.        // fellow
//
// What we need to do is create a virtual network interface. These are created
// in pairs, and are essentially connected to each other.      // fellow
//
// # ip link add veth0 type veth peer name veth1        // cmd
// Creates interfaces called veth0 and veth1. Anything sent to veth0 will appear
// on veth1, and vice versa. It's as if they were two separate ethernet cards,
// connected to the same hub (but not to anything else).
//
// run that command (outside the network namespace) we can list all of our
// available interfaces:
//
// # ip a
// 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
//     ...
// 2: ens5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 9001 qdisc mq state UP group default qlen 1000
//     ...
// 5: veth1@veth0: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN group default qlen 1000
//     link/ether ce:d5:74:80:65:08 brd ff:ff:ff:ff:ff:ff
// 6: veth0@veth1: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN group default qlen 1000
//     link/ether 22:55:4e:34:ce:ba brd ff:ff:ff:ff:ff:ff
//
// We can now move one of them -- veth1 -- into the network namespace netns1,
// which means that we have the interface outside connected to the one inside:
//
// # ip link set veth1 netns netns1          // cmd
//
// Now, from outside, we see this:
// # ip a
// 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
//     ...
// 2: ens5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 9001 qdisc mq state UP group default qlen 1000
//     ...
// 6: veth0@if5: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
//     link/ether 22:55:4e:34:ce:ba brd ff:ff:ff:ff:ff:ff link-netns netns1
//
// veth1 has disappeared (and veth0 is now @if5)
// But anyway, inside, we can see our moved interface:
//
// root@giles-devweb1:~# ip netns exec netns1 /bin/bash
// root@giles-devweb1:~# ip a
// 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
//     ...
// 5: veth1@if6: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
//     link/ether ce:d5:74:80:65:08 brd ff:ff:ff:ff:ff:ff link-netnsid 0
//
// At this point we have a network interface outside the namespace, which is
// connected to an interface inside. However, in order to actually use them,
// we'll need to bring the interfaces up and set up routing. The first step
// is to bring the outside one up; we'll give it the IP address 192.168.0.1 on
// the 192.168.0.0/24 subnet (that is, the network covering all addresses from
// 192.168.0.0 to 192.168.0.255)            // fellow // subnet // range
//
// # ip addr add 192.168.0.1/24 dev veth0        // cmd
// # ip link set dev veth0 up                    // cmd
// # ip a
// 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
//     ...
// 2: ens5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 9001 qdisc mq state UP group default qlen 1000
//     ...
// 6: veth0@if5: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state LOWERLAYERDOWN group default qlen 1000
//     link/ether 22:55:4e:34:ce:ba brd ff:ff:ff:ff:ff:ff link-netns netns1
//     inet 192.168.0.1/24 scope global veth0
//        valid_lft forever preferred_lft forever
//
// So that's all looking good; it reports "no carrier" at the moment, of course,
// because there's nothing at the other end yet. Let's go into the namespace and
// sort that out by bringing it up on 192.168.0.2 on the same network:
//
// # ip netns exec netns1 /bin/bash               // cmd
// # ip addr add 192.168.0.2/24 dev veth1         // cmd
// # ip link set dev veth1 up                     // cmd
// # ip a
// 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
//     ...
// 5: veth1@if6: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
//     link/ether ce:d5:74:80:65:08 brd ff:ff:ff:ff:ff:ff link-netnsid 0
//     inet 192.168.0.2/24 scope global veth1
//        valid_lft forever preferred_lft forever
//     inet6 fe80::ccd5:74ff:fe80:6508/64 scope link tentative
//        valid_lft forever preferred_lft forever
//
// Try pinging from inside the namespace to the outside interface    // fellow
//
// # ip netns exec netns1 /bin/bash
// # ping 192.168.0.1
// PING 192.168.0.1 (192.168.0.1) 56(84) bytes of data.
// 64 bytes from 192.168.0.1: icmp_seq=1 ttl=64 time=0.069 ms
// 64 bytes from 192.168.0.1: icmp_seq=2 ttl=64 time=0.042 ms
// ^C
// --- 192.168.0.1 ping statistics ---
// 2 packets transmitted, 2 received, 0% packet loss, time 1024ms
// rtt min/avg/max/mdev = 0.042/0.055/0.069/0.013 ms
//
// And from outside to the inside:              // fellow
//
// # ping 192.168.0.2
// PING 192.168.0.2 (192.168.0.2) 56(84) bytes of data.
// 64 bytes from 192.168.0.2: icmp_seq=1 ttl=64 time=0.039 ms
// 64 bytes from 192.168.0.2: icmp_seq=2 ttl=64 time=0.043 ms
// ^C
// --- 192.168.0.2 ping statistics ---
// 2 packets transmitted, 2 received, 0% packet loss, time 1018ms
// rtt min/avg/max/mdev = 0.039/0.041/0.043/0.002 ms
//
// However, of course, it's still not routed -- from inside the interface,
// we still can't ping Google's DNS server        // fellow
//
// # ip netns exec netns1 /bin/bash
// # ping 8.8.8.8
// ping: connect: Network is unreachable
//
// Connecting the namespace to the outside world with NAT
//
// We need to somehow connect the network defined by our pair of virtual
// interfaces to the one that is accessed via our real hardware network interface,
// either by setting up bridging or NAT. I'm running this experiment on a
// machine on AWS, and I'm not sure how well that would work with bridging
// (my guess is, really badly), so let's go with NAT.     // fellow
//
// First we tell the network stack inside the namespace to route everything
// via the machine at the other end of the connection defined by its internal
// veth1 IP address:
//
// # ip netns exec netns1 /bin/bash
// # ip route add default via 192.168.0.1    // cmd // this other machine is now the router
// # ping 8.8.8.8
// PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
// ^C
// --- 8.8.8.8 ping statistics ---
// 3 packets transmitted, 0 received, 100% packet loss, time 2028ms
//
// but of course there's nothing on the other side to send it onwards,
// so our ping packets are getting dropped on the floor.   // fellow
// We need to use iptables to set up that side of things outside the namespace. // fellow
//
// The first step is to tell the host that it can route stuff:
//
// # cat /proc/sys/net/ipv4/ip_forward
// 0
// # echo 1 > /proc/sys/net/ipv4/ip_forward        // cmd // iptables
// # cat /proc/sys/net/ipv4/ip_forward
// 1
//
// Now that we're forwarding packets, we want to make sure that we're not just
// forwarding them willy-nilly around the network. If we check the current rules
// in the FORWARD chain (in the default "filter" table):
//
// # iptables -L FORWARD                     // cmd
// Chain FORWARD (policy ACCEPT)
// target     prot opt source               destination
//
// We see that the default is ACCEPT, so we'll change that to DROP:
//
// # iptables -P FORWARD DROP                // cmd
// # iptables -L FORWARD
// Chain FORWARD (policy DROP)
// target     prot opt source               destination
//
// OK, now we want to make some changes to the nat iptable so that we have
// routing. Let's see what we have first:
//
// # iptables -t nat -L        // cmd
// Chain PREROUTING (policy ACCEPT)
// target     prot opt source               destination
// DOCKER     all  --  anywhere             anywhere             ADDRTYPE match dst-type LOCAL
//
// Chain INPUT (policy ACCEPT)
// target     prot opt source               destination
//
// Chain OUTPUT (policy ACCEPT)
// target     prot opt source               destination
// DOCKER     all  --  anywhere            !localhost/8          ADDRTYPE match dst-type LOCAL
//
// Chain POSTROUTING (policy ACCEPT)
// target     prot opt source               destination
// MASQUERADE  all  --  ip-172-17-0-0.ec2.internal/16  anywhere
//
// Chain DOCKER (2 references)
// target     prot opt source               destination
// RETURN     all  --  anywhere             anywhere
//
// I have Docker installed on the machine already, and it's got some of its
// own NAT-based routing configured there. I don't think there's any harm in
// leaving that there; it's on a different subnet to the one I chose for my own stuff.
//
// First, enable masquerading
// from the 192.168.0.* network onto main ethernet interface ens5   // fellow
//
// # iptables -t nat -A POSTROUTING -s 192.168.0.0/255.255.255.0 -o ens5 -j MASQUERADE
//
// Now forward stuff that comes in on ens5 can be forwarded to veth0 interface,
// which is the end of the virtual network pair that is outside the namespace:
// # iptables -A FORWARD -i ens5 -o veth0 -j ACCEPT
//
// then the routing in the other direction:
// # iptables -A FORWARD -o ens5 -i veth0 -j ACCEPT
//
// Now, let's see what happens if we try to ping from inside the namespace
// # ip netns exec netns1 /bin/bash
// # ping 8.8.8.8
// PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
// 64 bytes from 8.8.8.8: icmp_seq=1 ttl=112 time=0.604 ms
// 64 bytes from 8.8.8.8: icmp_seq=2 ttl=112 time=0.609 ms
// ^C
// --- 8.8.8.8 ping statistics ---
// 2 packets transmitted, 2 received, 0% packet loss, time 1003ms
// rtt min/avg/max/mdev = 0.604/0.606/0.609/0.002 ms
//
// Running a server with port-forwarding
// Right, now we have a network namespace where we can operate as a
// network client -- processes running inside it can access the external Internet.
//
// However, we don't have things working the other way around;
// we cannot run a server inside the namespace and access it from outside.
// For that, we need to configure port-forwarding.
//
// We use the "Destination NAT" chain in iptables:       // cmd
// # iptables -t nat -A PREROUTING -p tcp -i ens5 --dport 6000 -j DNAT --to-destination 192.168.0.2:8080
//
// in other words, if something comes in for port 6000 then we should sent it
// on to port 8080 on the interface at 192.168.0.2       // fellow
// i.e. the end of the virtual interface pair that is inside the namespace
//
// Next, we say that we're happy to forward stuff back and forth over
// new, established and related  connections to the IP of our namespaced interface
//
// # iptables -A FORWARD -p tcp -d 192.168.0.2 --dport 8080 -m state --state NEW,ESTABLISHED,RELATED -j ACCEPT
//
// we should be able to run a server inside the namespace on port 8080
// # ip netns exec netns1 /bin/bash
// # python3.7 server.py
//  * Serving Flask app "server" (lazy loading)
//  * Running on http://0.0.0.0:8080/ (Press CTRL+C to quit)
//
// from a completely separate machine on the same network as the one where we're
// running the server, we curl it using the machine's external IP address,
// on port 6000:
// $ curl http://10.0.0.233:6000/
// Hello from Flask!
//
- https://www.gilesthomas.com/2021/03/fun-with-network-namespaces
```

```c
// running a server in linux namespace & accessing it from another machine
// in the same network // just commands
//
// # ip netns add netns2
// # ip link add veth2 type veth peer name veth3
// # ip link set veth3 netns netns2
// # ip addr add 192.168.1.1/24 dev veth2
// # ip link set dev veth2 up
// # iptables -t nat -A POSTROUTING -s 192.168.1.0/255.255.255.0 -o ens5 -j MASQUERADE
// # iptables -A FORWARD -i ens5 -o veth2 -j ACCEPT
// # iptables -A FORWARD -o ens5 -i veth2 -j ACCEPT
// # iptables -t nat -A PREROUTING -p tcp -i ens5 --dport 6001 -j DNAT --to-destination 192.168.1.2:8080
// # iptables -A FORWARD -p tcp -d 192.168.1.2 --dport 8080 -m state --state NEW,ESTABLISHED,RELATED -j ACCEPT
//
//
// # ip netns exec netns2 /bin/bash
// # ip link set dev lo up
// # ip addr add 192.168.1.2/24 dev veth3
// # ip link set dev veth3 up
// # ip route add default via 192.168.1.1
// # python3.7 server2.py
//  * Serving Flask app "server2" (lazy loading)
//  * Running on http://0.0.0.0:8080/ (Press CTRL+C to quit)

```

```c
// how networks work

https://dzone.com/articles/how-networks-work-what-is-a-switch-router-dns-dhcp
```

```c
// how networks work - part ii

https://dzone.com/articles/how-networks-work-part-ii
```
