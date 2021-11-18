### State of API 2021

```yaml
https://kubernetes.io/blog/2021/04/13/kube-state-metrics-v-2-0/

// learn from release notes
// testing // troubleshooting // api design
// naming // standards

// Flag --namespace was renamed to --namespaces
// Flag --collectors was renamed to --resources

// --metric-blacklist to --metric-denylist
// --metric-whitelist to --metric-allowlist

// kube_hpa_* were renamed to kube_horizontalpodautoscaler_*
// Metric labels that relate to Kubernetes were converted to snake_case
```

```yaml
https://github.com/Peripli/service-broker-proxy-k8s

// service broker // k8s // proxy
```

```yaml
https://github.com/kubernetes/kube-aggregator
https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/aggregated-api-servers.md

// register // discover // proxy // multiple api servers // custom api servers
```

```yaml
https://github.com/kubernetes-sigs/prometheus-adapter/
https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/deploy/manifests/custom-metrics-apiservice.yaml

// custom api server implementation
// Kubernetes API server arguments // authentication // authorization
```
