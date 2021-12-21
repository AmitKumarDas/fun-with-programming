### EC2 - AWS
```yaml
- https://gist.github.com/ankyit/d180cdc2843a21204f27473a6c7eeb2c
```

### Modular Programming - Best Practices
```yaml
- https://github.com/cloudposse
- terraform - automation - aws - helm
```

### Syntax
#### Chain - for each - loop
```yaml
- https://www.terraform.io/docs/language/meta-arguments/for_each.html#chaining-for_each-between-resources
- https://www.terraform.io/docs/language/data-sources/index.html
```

#### Tuple
```yaml
- https://selleo.com/til/posts/cnfrqv1ipl-foreach-over-tuples-in-terraform
```

#### Providers
```yaml
- https://www.terraform.io/docs/language/providers/index.html
```

#### Naming
```yaml
- underscores then hyphens
- resource "your_provider_name" "abc-hey" 
- data "your_provider_name" "abc-hey"
```

#### Manifest - HCL - Track Diff - Server Side Apply
```yaml
- https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/manifest
```

#### Kubectl - CRD - Helm
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

# or
resource "kubectl_manifest" "install-operator-crd" {
  for_each = { for idx, resp in data.http.fetch-operator-crd : idx => resp }
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

#### kubectl_path_documents

```hcl
data "kubectl_path_documents" "docs" {
  pattern = "${path.module}/manifests/*.yaml"
}
resource "kubectl_manifest" "test" {
  for_each  = toset(data.kubectl_path_documents.docs.documents)
  yaml_body = each.value
}
```

#### File Parse - Split - Build Each As A Separate Manifest
```tf
locals {
  resource_list = yamldecode(file("${path. module}/example.yaml")).items
}

resource "kubectl_manifest" "example" {
  count = length(locals.resource_list)
  yaml_body = yamlencode(locals.resource_list[count.index]) 
}
```

#### Dynamic
```hcl
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.my_vpc.id

  # multiple route(s)
  dynamic "route" {
    for_each = local.peer_route_mapping
    content {
      cidr_block                = route.value["ip_range"]
      vpc_peering_connection_id = route.value["peering_id"]
    }
  }
  
  # one more
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.my_ig.id
  }
}
```

### Debugging
```sh
terraform workspace select my-work-xyz
terraform init
tarraform state list
```
