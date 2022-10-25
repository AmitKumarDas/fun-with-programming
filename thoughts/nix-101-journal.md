- WHEN: `curl -L https://nixos.org/nix/install | sh`
- THEN: `nix-shell -p nix-info --run "nix-info -m"`
- WHEN: `df -h`
- THEN:
```shell
Filesystem       Size   Used  Avail Capacity iused      ifree %iused  Mounted on
/dev/disk3s1s1  926Gi   14Gi  781Gi     2%  502068 4293531841    0%   /
devfs           200Ki  200Ki    0Bi   100%     691          0  100%   /dev
/dev/disk3s6    926Gi   20Ki  781Gi     1%       0 8194020720    0%   /System/Volumes/VM
/dev/disk3s2    926Gi  487Mi  781Gi     1%    1639 8194020720    0%   /System/Volumes/Preboot
/dev/disk3s4    926Gi  8.3Mi  781Gi     1%      41 8194020720    0%   /System/Volumes/Update
/dev/disk1s2    500Mi  6.0Mi  480Mi     2%       1    4911040    0%   /System/Volumes/xarts
/dev/disk1s1    500Mi  7.3Mi  480Mi     2%      25    4911040    0%   /System/Volumes/iSCPreboot
/dev/disk1s3    500Mi  2.3Mi  480Mi     1%      47    4911040    0%   /System/Volumes/Hardware
/dev/disk3s5    926Gi  129Gi  781Gi    15% 1051254 8194020720    0%   /System/Volumes/Data
map auto_home     0Bi    0Bi    0Bi   100%       0          0  100%   /System/Volumes/Data/home
/dev/disk3s7    926Gi  307Mi  781Gi     1%   55116 8194020720    0%   /nix
```
- WHEN: nix flake output is a nix function
- THEN: it can be used by other flakes
- WHEN: `nix-env -f '<nixpkgs>' -iA nixUnstable`
- THEN: `building '/nix/store/mp8lxnsn2i9f9kh5wr49557l02h9kdhb-user-environment.drv'...`
- WHEN: `nix show-config`
- THEN: `error: experimental Nix feature 'nix-command' is disabled; use '--extra-experimental-features nix-command' to override`
- THEN: `mkdir -p ~/.config/nix`
- THEN: echo 'experimental-features = nix-command flakes' >> ~/.config/nix/nix.conf
- WHEN: nix show-config
- THEN: ok
- WHEN: enter a shell that has GNU Hello from nixpkgs on branch nixpkgs-unstable
- THEN: `nix shell github:nixos/nixpkgs/nixpkgs-unstable#hello`
- WHERE: `github:nixos/nixpkgs/nixpkgs-unstable` is URL while `hello` is the attribute
- THEN: it gets downloaded and unpacked in what you can consider a cache directory
- WHEN: `hello`
```shell
Hello, world!
```
- HOW: `command -v hello`
```shell
/nix/store/hlng3mr5v7ghdis0as5k9z178k4s493l-hello-2.12.1/bin/hello
```
- WHEN: **build** instead of entering its shell
- THEN: `nix build nixpkgs#hello`
- THEN: 1 - it will either buuld hello or fetch it from binary cache if available
- THEN: 2 - symlink it to **result** in your current directory
```shell
ls -ltr

lrwxr-xr-x    1 amitd2  staff    56 Oct 24 18:28 result -> /nix/store/hlng3mr5v7ghdis0as5k9z178k4s493l-hello-2.12.1
```

### flake.nix
- MUST: One outputs attribute
- TIL: `self` refers to the flake that Nix is currently evaluating
- MINIMAL:
```nix
{
    outputs = { self }: { };
}
```


### References
```yaml
- https://serokell.io/blog/practical-nix-flakes
```
