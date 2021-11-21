## State of Infra 2021

### Terraform
```yaml
- https://itnext.io/the-definitive-guide-for-setting-up-a-clean-and-working-aws-eks-kubernetes-infrastructure-for-2d4153d47e20
- https://github.com/yandok/terraform-aws-eks-template
- EKS - terraform - solution - workflow - basics
```

### kubectl one liners
```yaml
- kubectl-convert -f ./my-deployment.yaml --output-version apps/v1
- kubectl get apiservice
- kubectl get apiservice v1beta1.custom.metrics.k8s.io -oyaml
- kubectl auth can-i create deployments --namespace dev
- kubectl auth can-i list secrets --namespace dev --as dave # user impersonation
- kubectl auth can-i list pods --namespace dev --as system:serviceaccount:dev:dev-sa
```

### Virtual Cluster
```yaml
- https://www.vcluster.com/
```

### Security
```yaml
- https://www.cncf.io/blog/2021/08/20/how-to-secure-your-kubernetes-control-plane-and-node-components/
- best practices - certificates - all core components

- https://blog.aquasec.com/runtime-security-tracee-rules - bpf 
- alerting via golang or rego
- socket - 

- https://docs.aws.amazon.com/eks/latest/userguide/specify-service-account-role.html
- iam - service account

- https://kubernetes.io/blog/2021/04/13/kube-state-metrics-v-2-0/
- capacity planning - monitoring - kube state metrics
- kube_pod_container_status_restarts_total can be used to alert on a crashing pod
- kube_deployment_status_replicas & kube_deployment_status_replicas_available
- alert on whether a deployment is rolled out successfully or stuck
- kube_pod_container_resource_requests and kube_pod_container_resource_limits for capacity planning

- https://github.com/common-fate/iamzero
- startup idea
- personised advice on IAM roles // platform team // end user
- security - read from error & propose the solution
- security policy for all clouds
- PGP Public Key - fingerprint - builds - penetration testing

- https://github.com/cruise-automation/rbacsync
- https://github.com/cruise-automation/k-rail

- https://kubernetes.io/docs/concepts/security/controlling-access/
- authentication 
- then authorization 
- then admission control

- https://kubernetes.io/docs/reference/access-authn-authz/authorization/
- kubectl - rbac - user impersonation - can i 
- kind: SelfSubjectAccessReview
- kind: SubjectAccessReview
- kind: LocalSubjectAccessReview
- ABAC - RABC - Node - Webhook - AlwaysAllow - AlwaysDeny
```

### DNS / Domain / Ingress / Proxy / Load Balancer
```yaml
- https://blog.cloudflare.com/secondary-dns-deep-dive/
- fellow - my next version - learn 

- https://github.com/gardener/dnslb-controller-manager
- load balancer over dns entries (til)
- depends on https://github.com/gardener/external-dns-management
- https://www.civo.com/learn/getting-started-with-load-balancers - basics - til

- https://github.com/gardener/external-dns-management
- manage external DNS entries for Kubernetes - custom resources - CRDs
- https://github.com/gardener/external-dns-management#the-model - architecture - fellow - learn CRD design

- https://www.digitalocean.com/community/tutorials/how-to-automatically-manage-dns-records-from-digitalocean-kubernetes-using-externaldns
- record - external dns - either ingress or external dns (til) - A or TXT record - external dns provider i.e. DO

- https://www.civo.com/learn/categories/dns - basics - testing - CNAME TXT record
- https://www.civo.com/learn/using-civo-dns-as-a-dynamic-dns-for-your-home-server - dig - testing - e2e

- https://github.com/ycd/dstp
- dns - domain - testing - network - dig - troubleshoot - e2e

- https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/coredns.md
- dns - external - nginx - dig - ingress - coredns with etcd backend

- https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/hostport.md
- headless service - kafka - sts - fqdn

- https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/ns-record.md
- namespace record - CRD - kind: DNSEndpoint - dns - external

- https://kubernetes.github.io/ingress-nginx/user-guide/ingress-path-matching/
- RE2 engine syntax - PCRE syntax - expressions
- https://github.com/google/re2/wiki/Syntax - regex - match - filter - assert
- host + path + service name + port ~ unique
- warning - unwanted path matching behaviour

- https://www.digitalocean.com/community/tutorials/understanding-nginx-server-and-location-block-selection-algorithms
- nginx can be - webserver - mailserver - proxyserver
- server name - host - ip address - port - listen - match
- request host's header - algorithm

- https://kubernetes.github.io/ingress-nginx/user-guide/custom-errors/
- ingress passes several http headers to backend in case of error

- https://kubernetes.github.io/ingress-nginx/user-guide/default-backend/
- /healthz returns 200
- / returns 404

- https://kubernetes.github.io/ingress-nginx/user-guide/multiple-ingress/
- https://kubernetes.github.io/ingress-nginx/user-guide/third-party-addons/opentracing/
- design - annotation - kind: IngressClass - tracing

- https://www.sanderverbruggen.com/add-an-http-echo-service-to-your-kubernetes-cluster/
- echo image

- https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ingress-guide-nginx-example.html
- ingress vs nodeport vs loadbalancer
- sample - testing - e2e - old samples

- til - ingress refers to oauth2 proxy service
- oauth proxy deployment for prometheus service exposed via ingress
- i.e. prometheus ingress refers to oauth proxy's kubernetes service
spec:
  automountServiceAccountToken: true
  containers:
  - args:
    - --provider=google
    - --upstream=http://prometheus-stack-kube-prom-prometheus:9090
    - --http-address=0.0.0.0:4180
    - --email-domain=abc.com
    - --htpasswd-file=/etc/htpasswd/auth
    env:
    - name: OAUTH2_PROXY_CLIENT_ID
      valueFrom:
        secretKeyRef:
    - name: OAUTH2_PROXY_CLIENT_SECRET
      valueFrom:
        secretKeyRef:
    - name: OAUTH2_PROXY_COOKIE_SECRET
      valueFrom:
        secretKeyRef:
    image: quay.io/oauth2-proxy/oauth2-proxy:v7.1.3

- https://stackoverflow.com/questions/68690185/how-to-create-dns-entries-for-localhost-kubernetes-ingress-hosts
- kind - localhost - access from desktop browser & from inside the pod

- minikube logs
- Kubernetes master is running at https://192.168.99.100:8443
- KubeDNS is running at https://192.168.99.100:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
```

### SLO
```yaml
- https://github.com/kruize/autotune
- tune - performance - recommendation - slo - api
```

### Prometheus | Grafana | Metrics | Logs | Observability
```yaml
- https://github.com/wangjia184/pod-inspector
- files - volumes - processes - cpu - memory - age - container - !kubectl

- https://github.com/vmware/kube-fluentd-operator
- custom resource - config reloader - fluentd

- https://github.com/K-Phoen/grabana
- https://github.com/K-Phoen/dark
- grafana in golang - custom resource - version the dashboard as code / yaml

- https://hackernoon.com/how-to-use-prometheus-adapter-to-autoscale-custom-metrics-deployments-p1p3tl0
- prometheus - e2e - curl nginx status - hpa - custom metrics

- https://github.com/kubernetes/kube-state-metrics/blob/master/examples/prometheus-alerting-rules/alerts.yaml
- snippets
``` 

### Orgs
```yaml
- Container Solutions Github:
- Secrets Management - platform - consulting - security
- Helm monitoring
- Rust via rustling - learn
- Trow vs Docker - Just in time images in Kubernetes
- Terraform Snippets - learn
- From Runbooks to automated E2E
- From How To Kubernetes documents to automated E2E
- Creating Docker Registry in Kubernetes - custom domain
- certificate - cert manager - dns - volumes - ebs

- Gardener:
- https://github.com/gardener/remedy-controller - ip address - cloud detect - cloud remedy
- https://github.com/gardener/cert-management/blob/master/test/functional/config/utils.go - e2e - testing
- https://github.com/gardener/cert-management/blob/master/test/functional/basics.go - cert - crd - controller - e2e
- https://github.com/gardener/cert-management - TLS certificate - vs cert manager - CRDs - DNS challenge - revoke
```

### YAML
```yaml
- https://sidneyliebrand.io/blog/the-greatnesses-and-gotchas-of-yaml
```

### Registry / Image
```yaml
- https://github.com/openshift/source-to-image
```

### Kubernetes code snippets
```yaml
- https://github.com/stefanprodan/kustomizer/blob/main/pkg/objectutil/io.go
- unstructured
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
