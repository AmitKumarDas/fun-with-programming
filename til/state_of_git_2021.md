### Rebase Better than Merge
```yaml
- https://www.linode.com/docs/guides/git-rebase-command/
- basics - learn
```

### Put Your Commits on Top of Latest Upstream Commit
```yaml
- git checkout main
- git pull upstream main
- git checkout my-feature-branch
- git rebase main
```

### Pretty Print Log
```yaml
- git log --pretty=oneline
```

### Unstage File(s) From Being Committed
```yaml
- git reset HEAD <file>
```

### Access To Private Github Repo
```yaml
- set below in ~/.profile
- machine github.com username AmitKumarDas password <my github token with relevant access>
```
