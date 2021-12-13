### Quota Service
```yaml
- https://github.com/square/quotaservice
- protobuf - hysterix - rate limiting
```


### Protocols
```yaml
- https://docs.google.com/document/d/1LPhVRSFkGNSuU1fBd81ulhsCPR4hkSZyyBj1SZ8fWOM/edit#heading=h.3p42p5s8n0ui
- Prometheus Remote Write Specs
- api - wire format - protocol

- til
- gRPC issues - internet behind load balancers such as EC2 ELB
- http headers - version
- retries - backoff - http response code

- https://prometheus.io/blog/2021/11/16/agent/
- Prometheus Remote Write Specs - Write Ahead Logging - WAL - TSDB WAL
- idea - can / should mayastor be a Remote Write Compliant Project?

- https://prometheus.io/blog/2016/07/23/pull-does-not-scale-or-does-it/
- push vs pull vs event - at scale
```


```yaml
- https://dev.to/harmonydevelopment/introducing-hrpc-a-simple-rpc-system-for-user-facing-apis-16ge
- hrpc - user facing APIs 
- protobuf

- https://blog.cloudflare.com/moving-k8s-communication-to-grpc/
- pod to pod communciation - http vs grpc - dns pipeline - protobuf
- large DNS zones - protobuf - quick record change propagation for serving API at the edge

- kind: NetworkPolicy - what goes in - what goes out - ingress - egress
- network policy - protect APIs at network level 
- no protection at application level

- https://grpc.io/docs/guides/auth/
- per rpc credentials - protection at app level - idea - mayastor
- HTTP/3 - quic

- https://itnext.io/websocket-1-million-connections-using-appwrite-2d2a2c363a37
- performance - design - testing

- https://github.com/processone/tsung
- scalability testing - multi protocol - MySQL - PostgreSQL - HTTP

- https://github.blog/2021-11-10-make-your-monorepo-feel-small-with-gits-sparse-index/
- monorepo - git - sparse index - contributing guide

- https://kubernetes.io/blog/2021/04/13/kube-state-metrics-v-2-0/
- learn from release notes - testing - troubleshooting - api design
- naming - standards

- https://github.com/Peripli/service-broker-proxy-k8s
- service broker - k8s - proxy

- https://github.com/kubernetes/kube-aggregator
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/aggregated-api-servers.md
- register - discover - proxy 
- multiple api servers - custom api servers
```

```yaml
- https://github.com/kubernetes-sigs/prometheus-adapter/
- https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/deploy/manifests/custom-metrics-apiservice.yaml
- custom api server implementation
- Kubernetes API server arguments 
- authentication - authorization
```
