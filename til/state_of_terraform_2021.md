### Terraform
```yaml
- https://www.terraform.io/docs/language/meta-arguments/for_each.html#chaining-for_each-between-resources
- https://www.terraform.io/docs/language/data-sources/index.html
- chain - for each - loop

- https://selleo.com/til/posts/cnfrqv1ipl-foreach-over-tuples-in-terraform
- tuple

- https://www.terraform.io/docs/language/providers/index.html
- providers offer:
- resources 
- data

- naming
- underscores then hyphens
- resource "your_provider_name" "abc-hey" 
- data "your_provider_name" "abc-hey"

- https://stackoverflow.com/questions/69180684/how-do-i-apply-a-crd-from-github-to-a-cluster-with-terraform
- CRDs - yaml - apply

- https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/manifest
- better option 
- tf can parse hcl & hence track the diff
- server side apply
```

```tf
variable "operator-crds" {
  type = list(string)
  default = [
    "monitoring.coreos.com_alertmanagerconfigs.yaml",
    "monitoring.coreos.com_alertmanagers.yaml",
  ]
}

# http get for each operator-crd files
data "http" "operator-crd" {
  count = length(var.operator-crds)
  # url is based on prometheus stack version 17.x
  url = "https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/v0.49.0/example/prometheus-operator-crd/${var.operator-crds[count.index]}"

  request_headers = {
    Accept = "text/plain"
  }
}

# kubectl apply for each operator crd that was http downloaded
resource "kubectl_manifest" "install-operator-crd" {
  # convert the tuple to map where key is the url & value is the object
  for_each = { for resp in data.http.operator-crd : resp.url => resp }
  yaml_body = each.value.body
}

# aliter - better
# server side apply for each operator crd that was http downloaded
resource "kubernetes_manifest" "install-operator-crd" {
  for_each = { for resp in data.http.operator-crd : resp.url => resp }

  # convert to HCL specs - allows tf to handle the diff better
  manifest = yamldecode(each.value.body)
}

resource "helm_release" "prometheus-stack" {
  depends_on = [kubernetes_manifest.install-operator-crd]
}
```

### Parsing the file and send each fragment as a separate manifest
```tf
locals {
  resource_list = yamldecode(file("${path. module}/example.yaml")).items
}

resource "kubectl_manifest" "example" {
  count = length(locals.resource_list)
  yaml_body = yamlencode(locals.resource_list[count.index]) 
}
```
