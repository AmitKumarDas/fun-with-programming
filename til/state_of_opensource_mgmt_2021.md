### State of OpenSource Management 2021

### Software Truths
```yaml
- https://buttondown.email/hillelwayne/archive/uncomfortable-truths-in-software-engineering/
```

### Vision
```yaml
- Ideas for projects, products, solutions, APIs, etc
- Ideas to sponsor the maintainers & hence the open source community
```

### Development Cycle - Before code

#### Watch These
```yaml
- https://www.youtube.com/watch?v=0Zbh_vmAKvk
```

#### Few Points To Remember
```yaml
- setup automation
- release automation
- doc automation
- linter automation
- local dev loop automation
- unit &/ integration automation
- e2e automation
- issue template
- website with proper domain(s)

- effective README
- https://github.com/brancz/kube-rbac-proxy
```

### Development Cycle - During code
```yaml
- Initial Days / Minimal Viable Product With Velocity
- Do you trust the person to submit a pull request
- Did we do the right hire
- Can this individual be given the full responsibility to bring up the project without a team

- Can we have one individual mapped to specific repo instead of a team mapped to a single repo
- Let the team brainstorm over the designs only
- In other words code reviews will be done by bots & linters
- Human code reviews will be optional & only on-demand
- Let the team adopt common template that has necessary automation (read bots & linters)

- small commits
- small PRs
- every feature is feature gated

- code must have code level comments
- code must have unit &/ integration tests 
- unit & integration tests must have proper comments
- let automation determine the approval or rejection of code

- do not force your style of coding & in turn delay the process
- stick to idiomatic & best practices
- its a marathon & futile to argue over the sprints

- Plugin based, composable, extensions, library, tracing, metrics comes before API
- API - both with & without
- DB - both with & without

- Declarative 
- YAML / Starlark / HCL 
- on top of API as well as library

- when API generate the SDK
- when API generate the website for API schema
- when API then proxy, load balancer, security, analytics

- APIs are great to test the product
```

### Development Cycle - Post Release
```yaml
- bugs to compliance kit
- bugs to post mortem analysis kit
```

### Process
```yaml
- `til` repo
- consist of relevant learnt items / links in md file(s)
- similar to awesome list

- `proposals` repo
- issues will have problem statement, solutions & comments
- issue graduates to a proposal md file

- use `MIRO` to state the team's work
- how to commercialise
- effective open source
- confirm the vision
- what's happening now
- what has changed
- change or no change in strategy when customers come in
- change or no change in strategy when community grows
- change or no change in strategy when we have multiple releases
- change or no change in strategy at scale

- standups should have following (every day can have a different theme)
- code reviews
- happiness reviews
- retrospectives
- vision reviews

- rules
- no more than 2 meetings per day
- some days can be no meeting day
- team events that force team to travel together, explore things beyond code
- should run, manage, upgrade, automate, scale, etc. the solution in EKS

- tools
- [+] GitHub
- [+] MIRO
- [+] EKS
- [-] JIRA
- [-] Confluence
- [~] PPTs
```

### Team should
```yaml
- have its own compliance kit
- i.e. team knows what to test
- i.e. team knows how to implement the scenarios

- have its own Post Mortem Analysis Kit 
- i.e. team is faced with some production issue e.g. high latency
- i.e. source code or logs wont help
- PMA kit should still be able to identify the issue

- have enough time to learn
- avoid mistakes made by others
- avoid reinvention of wheel
- avoid not understood here
- utilise existing projects as much as possible
- learn from bugs, wars, troubleshooting articles
- learn from other projects'
- release notes
- LWKD
- webinars
- slacks
- engineering blogs
```
