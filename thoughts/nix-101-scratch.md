## Scratching all over the place to use Nix & NixOS

### Nix Conversations Over Coffee
- WHAT: content addressable form
- SAMPLE: editor is stored in `/nix/store/frlxim9yz5qx34ap3iaf55caawgdqkip-neovim-0.5.1/`
- USP: software is never installed globally
- WANT: drop into a shell that has neovim in the $PATH
- THEN: `nix-shell -p neovim`
- USP: multiple versions of same software can coexist at the same time
- NOTE: nix stores all packages in its own unique subdirectory
- SAMPLE: `/nix/store/b6gvzjyb2pg0kjfwrjmg1vfhh54ad73z-firefox-33.1/`
- TIL: `/nix/store/<hash>-<name>-<version>`
- TIL: Hash is a unique identifier calculated from all the package’s dependencies 
- TIL: Hash is a cryptographic hash of the package’s build dependency graph
- TIL: Hash ensures files will not be tampered as the hash verifies the content mathematically
- TIL: Hash can enable caching
- WHAT: Everything in one big directory 
- USP: One big dir store is the one that makes Nix magical
- AVOID: Dependency version conflicts between packages (i.e. “DLL-hell”)
- USP: Every package can reference and use its own desired version
- USP: Atomic upgrades & rollbacks
- TIL: Derivations are independent of programming languages & are deterministic
- TIL: flake.nix is an interface to split & distribute Nix code & makes reuse & composition easy
- NICE: flake's inputs are optional dependencies while outputs are packages & deployment instructions
- WHAT: Exploring nix flakes as Go plugins
- WHERE: `https://flyx.org/nix-flakes-go/part1/`
- WHAT: bash derivation is a directory in the Nix store which contains bin/bash
- WHERE: e.g.: `/nix/store/s4zia7hhqkin1di0f187b79sa2srhv6k-bash-4.2-p45/`
- TRICK: There's no /bin/bash, there's only that self-contained build output in the store
- TRICK: To make them convenient to use from the shell, Nix will arrange for binaries to appear in your PATH as appropriate
- WHEN: ldd  `which bash`
- THEN: libc.so.6 => /nix/store/94n64qy99ja0vgbkf675nyk39g9b978n-glibc-2.19/lib/libc.so.6 (0x00007f0248cce000)


### Nix Refresher
- GIVEN: let ... in binding
```nix
let
    f = { a, b }: a + b;
in f { a = 10; b = 20; }
```
- WHEN: currying
- THEN:
```nix
let
    f = a: b: a + b;
in f 10 20
```
- THEN: f 10 evaluates to b: 10 + b
- THEN: f 10 20 evaluates to 30


### Nix 2.8.0
- CMD: `nix fmt`
- WHEN: read expressions from std input
- THEN: `--file`

### Nix 2.7.0
- GIVEN: A number of “default” flake output attributes have been renamed
```yaml
- defaultPackage.<system> → packages.<system>.default
- defaultApps.<system> → apps.<system>.default
- defaultTemplate → templates.default
- defaultBundler.<system> → bundlers.<system>.default
- overlay → overlays.default
- devShell.<system> → devShells.<system>.default
```
- CMD: `nix flake check`
- WHEN: report version of remote nix daemon
- THEN: `nix store ping`

### Nix 2.6.0
- WHEN: copy build logs from one store to another
- THEN: `nix store copy-log`
- WHEN: inside the docker container from mix master branch
- THEN: `docker run -ti nixos/nix:master`

### - nix2container - faster container image rebuilds
```yaml
- At build time nix uses go program to generate JSON files
  - Go generates these JSON files from the graph of store paths
- At runtime a modified Skopeo version uses the Go library to steam layers
  - Go library generate layer tar streams from these JSON files
  - Go library is coupled to Skopeo data structure
```

## References
```yaml
- https://www.youtube.com/watch?v=-hsxXBabdX0
- https://github.com/nlewo/nix2container
- https://github.com/divnix/std
- https://nixos.org/blog/announcements.html
- https://blog.wesleyac.com/posts/the-curse-of-nixos
- https://foo-dogsquared.github.io/blog/posts/moving-into-nixos/ WIP
```