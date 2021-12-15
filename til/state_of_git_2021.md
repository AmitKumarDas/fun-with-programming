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

### Change Older Commit Messages

```yaml
- Display last n commits in your current branch:
 - git rebase -i HEAD~n
 - `reword` instead of `pick` for all the commits you want to reword
- Save and close the commit list file
- In each resulting commit file, type the new commit message, save the file, and close it
- git push --force
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

### Monorepo | Sparse Index
```yaml
- https://github.blog/2021-11-10-make-your-monorepo-feel-small-with-gits-sparse-index/
```
