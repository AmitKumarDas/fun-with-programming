### Network

### dig
```yaml
- https://gist.github.com/mrlesmithjr/c42ebb99a01e8eeeca6a5eb4fa52f852
- dig output in json
- dig k8s.io +noall +answer | awk '{if (NR>3){print}}'| tr '[:blank:]' ';'| jq -R 'split(";") |{Name:.[0],TTL:.[1],Class:.[2],Type:.[3],IpAddress:.[4]}' | jq --slurp '.'

- https://github.com/inguardians/peirates/blob/master/enumerate_dns.go
- https://github.com/inguardians/peirates/blob/master/portscan.go
```

### SSH | Troubleshoot
```yaml
- SSH-Key has not been added to the ssh-agent
- Solve this with ssh-add -K
- You may put that command inside your .bashrc or .zshrc

- SSH does not know the host key of the bastion host yet
- Run below
- ssh -v <githubusername>@bastion.jimdo-platform-eks.net
- and then accept the question with yes
```
