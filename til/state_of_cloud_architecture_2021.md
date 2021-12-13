
### Architecture Diagram
```yaml
- https://threedots.tech/post/auto-generated-c4-architecture-diagrams-in-go/
```

### Google VPC
```yaml
- https://cloud.google.com/blog/topics/retail/shopify-and-google-cloud-team-up-for-an-epic-bfcm-weekend
- google virtual private cloud - https://cloud.google.com/vpc
- multi region - disaster recovery
```


### Monday.com
```yaml
- https://engineering.monday.com/monday-coms-multi-regional-architecture-a-deep-dive/
- saas platform
- multi region - privacy & compliance - latency - availability

- request flow:
- user - cloudflare - load balancer - ambassador edge stack - routes to services
- ambassador edge stack - Web Application Firewall (WAF)
- ambassador edge stack - auth service

- ambassador edge stack:
- envoy based L7 API gateway

- us-user.monday.com - us load balancer
- eu-user.monday.com - eu load balancer
- api.monday.com - ? load balancer
- anonymous-user.monday.com - ? load balancer
- webhooks - ?.monday.com # how will third party know which region to invoke

- does not scale with growing number of customers:
- mapping each account to unique DNS record
- expensive
- not scalable when customers grow beyond thousands

- how to be vendor neutral:
- cloudflare workers
- aws cloudfront with lambda@edge
- fastly compute@edge

- reproducibility:
- how do you build a new region

- multi region design:
- us:
  vpc:
    - front facing lb - ambassador edge stack - internal lb
    - internal lb - transit gateway
  transit gateway # out of vpc
- eu:
  vpc:
    - front facing lb - ambassador edge stack - internal lb
    - internal lb - transit gateway
  transit gateway # out of vpc
- us-eu:
  - transit gateway peering
```