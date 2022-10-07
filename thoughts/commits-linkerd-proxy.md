## Learn LINKERD PROXY from its commits
This is one of my ideas to learn a project. In other words, read and perhaps try interesting
commits from the project. This should help me in understanding parts of the project by focusing
on some particular fix or feature. Alternative ways e.g. getting involved with the community
or spending weeks &/ months may not be feasible unless it is part of my day job. Needless to say
this works better when the project follows atomic commits.

### Default inbound connection idle timeout
- Should be less than or equal to the server's idle timeout
- So that we don't try to reuse a connection as it is being timed out of the server
- TAG: ReadHeaderTimeout, IdleTimeout,
- ISSUE: https://github.com/golang/go/issues/54784
- PR: https://github.com/golang/go/pull/54785
- ISSUE: https://github.com/linkerd/linkerd2/issues/9273
- PR: https://github.com/linkerd/linkerd2/pull/9272
- PR: https://github.com/linkerd/linkerd2-proxy/pull/1931

### DNS: Prefer SRV records over A/AAAA
- PR: https://github.com/linkerd/linkerd2-proxy/pull/1915

### ENV: Shutdown grace period timeout for graceful shutdowns
- TAG: tokio, await
- ISSUEL https://github.com/linkerd/linkerd2/issues/8033
- PR: https://github.com/linkerd/linkerd2-proxy/pull
