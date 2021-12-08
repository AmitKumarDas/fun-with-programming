### Tilt & Friends - Inner Development Loop
```yaml
- https://github.com/tilt-dev/tilt/tree/master/configs/storybook
- Helm chart, Tiltfile, Dockerfile
```

### Testing via Helm
```yaml
- Test via a Pod that is Helm-ed
- Pod with Helm annotation to test if a deployment installed via Helm Chart works
  - "helm.sh/hook": test-success
  - Use wget or any simple thing
  - This test pod itself is installed via Helm & then used a test tool
```

### Tilt Commands
```yaml
- str
- local
- rstrip
- split

- allow_k8s_contexts
- local_resource
- experimental_analytics_report
- analytics_settings
- docker_build
- sync
- k8s_yaml
- k8s_resource
- helm
```

#### Advanced Commands
```yaml
- https://github.com/kubernetes-sigs/cluster-api/blob/main/Tiltfile
```
