### Learn Helm

```yaml
- refer:
  - https://blog.polymatic.systems/genero-3f788f2167a7
  - https://github.com/polymatic-systems/helm-charts
  - https://github.com/polymatic-systems/tempro
```

#### Basic Instruction to Use Helm
```yaml
helm repo add polymatic https://polymatic-systems.github.io/helm-charts

helm upgrade ${RELEASE} polymatic/genero \
  --install --namespace ${NAMESPACE} \
  --version ${VERSION} --values values.yaml \
  --set default.name="${RELEASE}" \
  --set default.version="${GIT_REF}" \
  --set default.image="${IMAGE}" \
  --set default.slug="${SLUG}" \
  --set default.env="${ENVIRONMENT}" \
  --set default.registries="{private,dockerhub}" \
  --set default.add.metrics.enable."prometheus\.io/scrape"="true" \
  --set default.add.metrics.enable."prometheus\.io/port"="5000" \
  --set default.add.metrics.enable."prometheus\.io/path"="/metrics" \
  --set default.add.metrics.disable."prometheus\.io/scrape"="false"
```
```yaml
- upgrade command with the --install argument
- ensures same command for installation and updating

- Genero is designed so that if even if you don’t make any changes
- upgrade command will cause the pods in the deployment to be cycled
```


#### Preview Environments

```yaml
default:
  slug: '' # used for review deployments, changes some naming
  env: production
  # these should be your ClusterIssuer names
  issuer:
    production: production
    staging: staging
  # these should be your IngressClass names
  ingress:
    external: external
    internal: internal
```

```yaml
- slug field is designed for review environments
- developers launch the same app multiple times for working on different branches
- If slug exists, it will be prefixed to any defined Ingress hosts
- Another side effect, is that it will set the default certificate issuer to your staging issuer
```
