### State of CLI | Tools | Commands

### Network
```yaml
- dig
- https://gist.github.com/mrlesmithjr/c42ebb99a01e8eeeca6a5eb4fa52f852
- dig output in json
- dig k8s.io +noall +answer | awk '{if (NR>3){print}}'| tr '[:blank:]' ';'| jq -R 'split(";") |{Name:.[0],TTL:.[1],Class:.[2],Type:.[3],IpAddress:.[4]}' | jq --slurp '.'
```

### Git
```yaml
- https://www.linode.com/docs/guides/git-rebase-command/
- git - rebase vs merge

- steps you need to put your commits on top of latest commit in upstream
- git checkout main
- git pull upstream main
- git checkout my-feature-branch
- git rebase main

- git log --pretty=oneline
```
