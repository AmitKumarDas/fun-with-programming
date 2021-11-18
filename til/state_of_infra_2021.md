## State of Infra 2021

### Virtual Cluster
- https://www.vcluster.com/

### Security
- https://docs.aws.amazon.com/eks/latest/userguide/specify-service-account-role.html
  - // iam // service account
- https://kubernetes.io/blog/2021/04/13/kube-state-metrics-v-2-0/
  - capacity planning // monitoring // kube state metrics
  - kube_pod_container_status_restarts_total can be used to alert on a crashing pod
  - kube_deployment_status_replicas & kube_deployment_status_replicas_available
  - alert on whether a deployment is rolled out successfully or stuck
  - kube_pod_container_resource_requests and kube_pod_container_resource_limits for capacity planning

### DNS
- https://github.com/ycd/dstp
  - // dns // domain // testing // network // dig // troubleshoot // e2e

### SLO
- https://github.com/kruize/autotune
  - // tune // performance // recommendation // slo // api

### Prometheus
- https://hackernoon.com/how-to-use-prometheus-adapter-to-autoscale-custom-metrics-deployments-p1p3tl0
  - // prometheus // e2e // curl nginx status // hpa // custom metrics
- Deployment specs
```yaml
spec: # of nginx deployment
  template:
    metadata:
      annotations:
        prometheus.io/path: "/status/format/prometheus"
        prometheus.io/scrape: "true"
        prometheus.io/port: "80"
```

### Orgs
- Container Solutions Github
  - // Secrets Management // platform // consulting // security
  - // Helm monitoring
  - // Rust via rustling // Learn // startup
  - // Trow vs Docker // Just in time images in Kubernetes
  - // Terraform Snippets // fellow // support // startup
  - // From Runbooks to automated E2E
  - // From How To Kubernetes documents to automated E2E
  - // Custom Domains
  - // Creating Docker Registry in Kubernetes
  - // certificate // cert manager // dns // volumes // ebs


```
// yaml // templating // variable declaration
https://sidneyliebrand.io/blog/the-greatnesses-and-gotchas-of-yaml
```

```
// build artifacts from source
// inject into containers
// can version & control your build environments
https://github.com/openshift/source-to-image
```

```
// go // python
// personised advice on IAM roles // platform team // end user
// security // read from error & propose the solution
// startup idea // security policy for all clouds // openebs
// K8s // GCP // Azure
//
// PGP Public Key // fingerprint // builds // penetration testing
https://github.com/common-fate/iamzero
```


```c
// policy // security

https://github.com/cruise-automation/rbacsync
https://github.com/cruise-automation/k-rail
```


```
// yaml to unstruct // unstruct to yaml
//
https://github.com/stefanprodan/kustomizer/blob/main/pkg/objectutil/io.go
```

```
// alerts // prometheus // indentation // rules
https://github.com/kubernetes/kube-state-metrics/blob/master/examples/prometheus-alerting-rules/alerts.yaml
```

```
// prometheus // api // testing // e2e
https://prometheus.io/docs/prometheus/latest/querying/api/
```

```c
// If the URL storage.jimdosite-stage.com reported an SSL error
//
// $ dig storage.jimdosite-stage.com // cmd

// ; <<>> DiG 9.10.6 <<>> storage.jimdosite-stage.com
// ;; global options: +cmd
// ;; Got answer:
// ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 16840
// ;; flags: qr rd ra; QUERY: 1, ANSWER: 4, AUTHORITY: 0, ADDITIONAL: 1
//
// ;; OPT PSEUDOSECTION:
// ; EDNS: version: 0, flags:; udp: 1232
// ;; QUESTION SECTION:
// ;storage.jimdosite-stage.com.	IN	A
//
// ;; ANSWER SECTION:
// storage.jimdosite-stage.com. 300 IN	CNAME	dolphin-content-storage-stage.jimdo-platform.net.
// dolphin-content-storage-stage.jimdo-platform.net. 20 IN	CNAME dolphin-conten-906dd1-1511968020-1586625729.eu-west-1.elb.amazonaws.com.
// dolphin-conten-906dd1-1511968020-1586625729.eu-west-1.elb.amazonaws.com. 60 IN A 54.220.29.254
// dolphin-conten-906dd1-1511968020-1586625729.eu-west-1.elb.amazonaws.com. 60 IN A 52.19.91.241
//
// ;; Query time: 387 msec
// ;; SERVER: 1.1.1.1#53(1.1.1.1)
// ;; WHEN: Wed Nov 10 11:32:26 CET 2021
// ;; MSG SIZE  rcvd: 232
//
// So that points to dolphin-content-storage-stage.jimdo-platform.net
```

```
# how to go get any private github repo?
# ans: set below in ~/.profile

machine github.com username AmitKumarDas password <my github token with relevant access>
```

```
// dns // external // nginx // dig // ingress // coredns with etcd backend
- https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/coredns.md
```

```
// headless service // kafka // sts // fqdn
- https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/hostport.md
```

```
// namespace record // CRD // kind: DNSEndpoint // dns // external
// custom resource definition
- https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/ns-record.md
```

```yaml
apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
spec:
  groups:
    - name: ./aws-limits.rules
      rules:
        # We just pick one sample metric for the absent check to verify that the exporter is working
        - alert: AWSServiceLimitMetricsAbsent
          expr: absent(aws_ec2_limit_total) or absent(aws_ec2_used_total)
          for: 10m
          annotations:
            description: Missing aws_limits metrics. Is aws-limits-exporter running?
        - alert: AWSServiceLimitUsageHigh
          expr: '{__name__=~"aws_.*_used_total"} / {__name__=~"aws_.*_limit_total"} * 100 > 80'
          annotations:
            description: AWS resource "{{ $labels.resource }}" in region "{{ $labels.region }}" is using {{ $value }}% of its limit. Check usage and think about requesting a limit increase.
        - alert: AWSServiceLimitUsageCritical
          expr: '{__name__=~"aws_.*_used_total"} / {__name__=~"aws_.*_limit_total"} * 100 == 100'
          annotations:
            description: AWS resource "{{ $labels.resource }}" in region "{{ $labels.region }}" is using {{ $value }}% of its limit. Check usage and think about requesting a limit increase.
```


```
// generates dynamic X.509 certificates
// without going through the usual manual process of generating a private key and CSR
// submitting to a CA, and waiting for a verification and signing process to complete
//
// Vault's built-in authentication and authorization mechanisms provide the
// verification functionality.

https://www.vaultproject.io/docs/secrets/pki
```

```
// vault // root ca // cert manager // mTLS // istio // istio-csr agent

https://medium.com/everything-full-stack/how-to-integrate-vault-as-external-root-ca-with-cert-manager-istio-csr-and-istio-7684baa369db
```

```
// iam scoped to service accounts
// no need for kiam or kube2iam
//
// steps:
//
// 0/ create iam oidc provider -- once per cluster lifecycle
// 1/ create iam role
// 2/ create iam policy with required roles
// 3/ attach iam role with iam policy
// 4/ associate iam role with service account
- https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
```

```
// aws // secrets // mount // API keys // certificates // DB credentials // csi
// iam // service accounts // aws secrets manager // ssm
- https://github.com/aws/secrets-store-csi-driver-provider-aws
```

```
// http 2 vs http 3
// http connection vs tcp connection
- https://medium.com/mutualmobile/http-3-what-was-what-is-what-will-be-bcade7df032b
```

```
// aws // node // interruption // scale // rebalance // drain // cordon
// graceful shutdown
- https://github.com/aws/aws-node-termination-handler
```

```go
// aws cis

- https://blog.gruntwork.io/cis-aws-v1-4-is-out-4131be41d00c
- https://blog.gruntwork.io/its-time-to-update-to-version-1-3-0-of-the-cis-aws-foundations-benchmark-7ebd24080031
```

```go
// virtual cluster
- https://github.com/loft-sh/vcluster
```

```go
// curl cluster ip? pod ip?
// proxy // socket // iptables // github // source code // ebpf
- http://arthurchiao.art/blog/cracking-k8s-node-proxy/
```

```go
// run shell commands from webhook
- https://github.com/adnanh/webhook
```

```go
// troubleshoot: describe svc
// -------------
// -- check for endpoints
// -- implies pods are running // passed health checks
// -- these endpoints must be present in pod ips
//
// troubleshoot: describe ingress
// -------------
// -- check for backends
// -- curl after exec into ingress pod
// -- curl -H "HOST: www.myexample.com" localhost
// -- above is virtual host routing -- this can not detect DNS or LoadBalancer issues
//
// troubleshoot: get svc & grab ELB URL
// -- abcdxxxxx-549.us-east-1.elb.amazonaws.com
// -- curl over ELB with host headers // run from your local machine
// -- curl -H "HOST: www.example.com" abcdxxxxx-549.us-east-1.elb.amazonaws.com
//
// troubleshoot: is DNS resolving
// -- nslookup www.example.com // run from your local machine
// -- above should return the canonical name as the ELB external ip address
//
- https://medium.com/@ManagedKube/kubernetes-troubleshooting-ingress-and-services-traffic-flows-547ea867b120
```

```go
// cert manager // ingress // external dns // route 53 // network load balancer
// nlb // terraform // tf versions // tf main // tf output // tf iam roles
// github // tf provider aws -> tf module eks -> tf provider kubernetes
// hostname passes from aws to eks to kubernetes // tf provides the kubeconfig
// aws account id to kubernetes service account
- https://www.appvia.io/blog/expose-kubernetes-service-eks-dns-tls
- https://github.com/appvia/How-to-expose-a-Kubernetes-web-application-with-DNS-and-TLS
```

```go
// external dns watches all ingresses
// - picks up hostname from ingress to ingress controller's load LoadBalancer
// - & creates a record in route 53
//
// IAM role with route 53 permissions
// domain filter // hosted zone id // role arn
//
// various ways:
// load balancer service // node port service // ingress
- https://www.padok.fr/en/blog/external-dns-route53-eks
```

```
// domain registrar // A record // CNAME // ingress // load balancer // external ip
- https://hackernoon.com/expose-your-kubernetes-service-from-your-own-custom-domains-cc8a1d965fc
```

```
// centralized authentication // rbac // CIS benchmarking
- https://www.suse.com/c/accelerating-security-innovation/
```

```
// jvm // requestss // limits // cpu // throttling
- https://tech.olx.com/improving-jvm-warm-up-on-kubernetes-1b27dd8ecd58
```

```
// secure ingress
// ingress controller
// -- provides an HTTP proxy service via your cloud provider’s load balancer
// -- few minutes for the LoadBalancer IP to be available
// -- minute or two for the cloud provider to provide and link a public IP address
// -- External IP == public IP == ingress controller svc's External IP  -- all incoming traffic is routed here
// -- add this IP to a DNS zone you control -- i.e. assign DNS entry to an IP address
// app specific - kubectl get ing - Address takes a few minutes
// curl -kivL -H 'Host: www.example.com' 'http://203.0.113.2' // IP here is ing address
// the service will be available with a TLS certificate
// -- but it will be using a self-signed certificate from the ingress-nginx-controller
// -- Browsers will show a warning that this is an invalid certificate
// Issuer vs ClusterIssuer // email to notify
- https://cert-manager.io/docs/tutorials/acme/ingress/
```

```
// openapi // kusk vs helm // kusk vs docker compose
// generate configurations for ingress, ambassador, linkerd
// openapi extension for Kubernetes configurations
// CORS // timeouts // mapping transformations // namespace/service // security
//
// idea
// - kusk then kustomizer
// - yaml to yaml generation is still a great idea - golang + starlark
// - move left to developers + move right to operations
// - left   =~ openapi &/ openapi extensions
// - middle =~ yaml to yaml generator &/ starlark
// - right  =~ kustomizer &/ Work API
- https://kubeshop.io/blog/hello-kusk-openapi-for-kubernetes
```

```
// seccomp by default // release // rootless mode containers
// cgroupsv2 // memory QoS // performance // throttling
// server side apply from curl // priority & fairness
// immutable label selector for namespaces // logarithmic scale down
// pod deletion cost // job index to pods via env // job suspend
// external client-go credential providers // KMS // TPM // HSM
// expanded DNS configuration // endpointslice // multiple IPs per pod
// LoadBalancerClass // leverage multiple LoadBalancer in a cluster
// cpu manager - assign exclusive workloads to specific CPU cores
// pod can now request huge pages of different sizes
// pod hostname can be set to FQDN - legacy applications
// memory manager to guarantee memory & hugepages for pods
// ReadWriteOnce vs ReadWriteOncePod // apiserver tracing via OpenTelemetry
- https://sysdig.com/blog/kubernetes-1-22-whats-new/
```

```
// terraform // kops // aws // dns // route 53
// kops generate terraform // vpc // subnet pairs // ec2
// nat gateways // security groups // load balancer backed by ELB
// s3 bucket // ssh keys
- https://ryaneschinger.com/blog/kubernetes-aws-vpc-kops-terraform/
```

```
// test environments // preview environments // ephemeral environments
// external dns
- https://blog.codemine.be/posts/20190125-devops-eks-externaldns/
```

```
// preview environments // namespace
- https://codefresh.io/kubernetes-tutorial/unlimited-preview-environments/
```

```
// automatic // dns // aws // external dns // nginx ingress controller
// elb frontends ingress // load balancer creates ELB // route
// parent domain // zone for subdomain
// internet -> ELB -> ingress -> k8s service
- https://ryaneschinger.com/blog/automatic-dns-kubernetes-ingresses-externaldns/
```

```
// network // e2e // benchmark // compliance // testing
- https://github.com/kinvolk/egress-filtering-benchmark
```

```
// e2e testing // learn // adopt
- https://github.com/kubeshop/testkube
```

```
// vs wonderland // any helm chart //
- https://github.com/polymatic-systems/helm-charts
```

```
// different IAM roles for K8s Pods
- https://github.com/jtblin/kube2iam
```

```
// dns // external // route 53 // cloudflare // security // rancher dns // crd
// external dns creates dns records from ingress resources
// refer kind: Ingress spec.rules[x].host
// next step - generate TLS certificate for ingress resources with Let's Encrypt
// ExternalDNS currently requires full access to a single managed zone in Route 53
// it will delete any records that are not managed by ExternalDNS
- https://github.com/kubernetes-sigs/external-dns
```

```
// dns // e2e testing // external dns
- https://github.com/zalando-incubator/kubernetes-on-aws/blob/dev/test/e2e/external_dns.go
```

```
// script // DevOps // e2e testing // vs. terraform
- https://github.com/bazelbuild/starlark
```

```
// 12 factor // microservice // boilerplate // instrumentation // tracing
// swagger // signals // health checks // file watcher // mult arch container image
// CVE scanning // linkerd service profile // structured logging // JWT // gRPC
// examples // samples // end to end testing // EKS ingress // Let's Encrypt
// GitOps // Fargate // Auto scaling // custom metrics // istio metrics // canary
- https://github.com/stefanprodan/podinfo
```

```
// controller // runtime // operator // cli // vs. terraform
- https://github.com/stefanprodan/kustomizer
```

# ===========================================
# Kubernetes TILs
# ===========================================

- EndpointSlice producer nothing but its own controller
- Kube-Proxy is one of the EndpointSlice consumer
- Headless service - without cluster ip
- Named service - with cluster ip

# ===========================================
# Proposals
# ===========================================

// network // zone // multicluster // hpa // kube-proxy // design
- https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2433-topology-aware-hints

// pull // push // gitops // work api //
- https://docs.google.com/document/d/1cWcdB40pGg3KS1eSyb9Q6SIRvWVI8dEjFp9RI0Gk0vg/

// multi cluster dns // A/AAAA // PTR // SRV // clusterset ip // service export
- https://docs.google.com/presentation/d/1-BEpCs5JE_l6EDfEmbuXcT2WapCqUk387nJN1LkOypw/



```yaml
kind: Ingress
spec:
  tls:
      - hosts:
          - www.example.com
        secretName: example-tls # enable TLS for ingress
---
kind: Secret
metadata:
  name: example-tls
```

```sh
# shows TLS headers // follows redirects
curl -kivL -H 'Host: www.example.com' 'http://203.0.113.2'
```

```yaml
# cert-manager will have created a secret with the details of the certificate
# based on the secret used in the ingress resource

# kubectl describe certificate quickstart-example-tls // owned by Ingress
# kubectl describe order quickstart-example-tls-889745041 // i.e. ACME order
# ACME Order is a Challenge for a domain
#
# kubectl describe challenge quickstart-example-tls-889745041-0
# cert-manager waits for the challenge record to propagate to the ingress controller
# Once the challenge(s) have been completed, corresponding challenge resources
# will be deleted, and the ‘Order’ will be updated & finally certificate is updated
#
# certificate -> generates new private key <-> order <-> challenge
# kubectl describe secret quickstart-example-tls
Name:         quickstart-example-tls
Namespace:    default
Labels:       cert-manager.io/certificate-name=quickstart-example-tls
Annotations:  cert-manager.io/alt-names=www.example.com
              cert-manager.io/common-name=www.example.com
              cert-manager.io/issuer-kind=Issuer
              cert-manager.io/issuer-name=letsencrypt-staging

Type:  kubernetes.io/tls

Data
====
tls.crt:  3566 bytes
tls.key:  1675 bytes
```

### kubectl tip
- kubectl-convert -f ./my-deployment.yaml --output-version apps/v1
- kubectl get apiservice
- kubectl get apiservice v1beta1.custom.metrics.k8s.io -oyaml
    
### Blogs
- Tilt blogs

### E2E testing library 
- https://github.com/kubeshop/testkube
  - Avoid vendor lock-in for test orchestration and execution in CI/CD pipelines
  - Make it easy to orchestrate and run any kind of tests
    - functional, load/performance, security, compliance, etc.
    - in your clusters, without having to wrap them in docker-images or ??
    - providing network access ??
  - Make it possible to decouple test execution from build processes;
    - engineers should be able to run specific tests whenever needed
  - Centralize all test results in a consistent format for "actionable QA analytics"
  - Provide a modular architecture for adding new types of test scripts and executors
