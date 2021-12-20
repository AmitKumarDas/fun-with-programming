### Kubernetes

#### Fetch Status of Previously Terminated Containers
```yaml
- kubectl get pod with the -o go-template=...
  - kubectl get pod -o go-template='{{range.status.containerStatuses}}{{"Container Name: "}}{{.name}}{{"\r\nLastState: "}}{{.lastState}}{{end}}'  simmemleak-hra99
- O/P:
  - Container Name: simmemleak
  - LastState: map[terminated:map[exitCode:137 reason:OOM Killed startedAt:2015-07-07T20:58:43Z finishedAt:2015-07-07T20:58:43Z containerID:docker://0e4095bba1feccdfe7ef9fb6ebffe972b4b14285d5acdec6f0d3ae8a22fad8b2]]
```

### Network

### Domain Profiler
```yaml
- https://github.com/jpf/domain-profiler
```
```sh
$ ./profile github.com
Fetching data for github.com: DNS Whois SSL ...


==========[ github.com ]==========
Web Hosting:
  (Rackspace)
      207.97.227.239

DNS Hosting:
  (anchor.net.au)
      ns1.anchor.net.au.
      ns2.anchor.net.au.
  (EveryDNS.net)
      ns1.everydns.net.
      ns2.everydns.net.
      ns3.everydns.net.
      ns4.everydns.net.

Email Hosting:
  (Google)
      1 ASPMX.L.GOOGLE.com.
      10 ASPMX2.GOOGLEMAIL.com.
      10 ASPMX3.GOOGLEMAIL.com.
      5 ALT1.ASPMX.L.GOOGLE.com.
      5 ALT2.ASPMX.L.GOOGLE.com.

Domain Registrar:
  (Go Daddy)

SSL Issuer:
  (GoDaddy.com, Inc.)
      Common Name: *.github.com
```

### Dig
```yaml
- https://gist.github.com/mrlesmithjr/c42ebb99a01e8eeeca6a5eb4fa52f852
- dig output in json
- dig k8s.io +noall +answer | awk '{if (NR>3){print}}'| tr '[:blank:]' ';'| jq -R 'split(";") |{Name:.[0],TTL:.[1],Class:.[2],Type:.[3],IpAddress:.[4]}' | jq --slurp '.'

- https://github.com/inguardians/peirates/blob/master/enumerate_dns.go
- https://github.com/inguardians/peirates/blob/master/portscan.go
```

## SSH | Troubleshoot
```yaml
- SSH-Key has not been added to the ssh-agent
- Solve this with ssh-add -K
- You may put that command inside your .bashrc or .zshrc

- SSH does not know the host key of the bastion host yet
- Run below
- ssh -v <githubusername>@bastion.jimdo-platform-eks.net
- and then accept the question with yes
```

## Makefile
### Define A Function
```yaml
- https://coderwall.com/p/cezf6g/define-your-own-function-in-a-makefile
```
