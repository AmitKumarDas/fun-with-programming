- START: `brew install lima`
- THEN: `limactl start`
```shell
? Creating an instance "default" Proceed with the current configuration
INFO[0021] Attempting to download the image from "https://cloud-images.ubuntu.com/releases/22.04/release-20220902/ubuntu-22.04-server-cloudimg-arm64.img"  digest="sha256:9620f479bd5a6cbf1e805654d41b27f4fc56ef20f916c8331558241734de81ae"
611.06 MiB / 611.06 MiB [-----------------------------------] 100.00% 4.84 MiB/s
INFO[0149] Downloaded the image from "https://cloud-images.ubuntu.com/releases/22.04/release-20220902/ubuntu-22.04-server-cloudimg-arm64.img" 
INFO[0151] Attempting to download the nerdctl archive from "https://github.com/containerd/nerdctl/releases/download/v0.23.0/nerdctl-full-0.23.0-linux-arm64.tar.gz"  digest="sha256:d25171f8b6fe778b77ff0830a8e17bd61c68af69bd734fb9d7f4490e069a7816"
186.88 MiB / 186.88 MiB [-----------------------------------] 100.00% 7.41 MiB/s
INFO[0178] Downloaded the nerdctl archive from "https://github.com/containerd/nerdctl/releases/download/v0.23.0/nerdctl-full-0.23.0-linux-arm64.tar.gz" 
INFO[0183] [hostagent] Starting QEMU (hint: to watch the boot progress, see "/Users/amitd2/.lima/default/serial.log") 
INFO[0183] SSH Local Port: 60022                        
INFO[0184] [hostagent] Waiting for the essential requirement 1 of 5: "ssh" 
INFO[0201] [hostagent] Waiting for the essential requirement 1 of 5: "ssh" 
INFO[0201] [hostagent] The essential requirement 1 of 5 is satisfied 
INFO[0201] [hostagent] Waiting for the essential requirement 2 of 5: "user session is ready for ssh" 
INFO[0201] [hostagent] The essential requirement 2 of 5 is satisfied 
INFO[0201] [hostagent] Waiting for the essential requirement 3 of 5: "sshfs binary to be installed" 
FATA[0778] did not receive an event with the "running" status 
```
- DETOUR: `https://github.com/kakkoyun/lima-config`
- THEN: `git clone git@github.com:kakkoyun/lima-config.git`
- THEN: `limactl start ./ebpf.yml`
```shell
? Creating an instance "ebpf" Proceed with the current configuration
INFO[0014] Attempting to download the image from "https://cloud-images.ubuntu.com/releases/22.04/release/ubuntu-22.04-server-cloudimg-arm64.img"  digest=
613.25 MiB / 613.25 MiB [-----------------------------------] 100.00% 4.66 MiB/s
INFO[0147] Downloaded the image from "https://cloud-images.ubuntu.com/releases/22.04/release/ubuntu-22.04-server-cloudimg-arm64.img" 
INFO[0147] Attempting to download the nerdctl archive from "https://github.com/containerd/nerdctl/releases/download/v0.23.0/nerdctl-full-0.23.0-linux-arm64.tar.gz"  digest="sha256:d25171f8b6fe778b77ff0830a8e17bd61c68af69bd734fb9d7f4490e069a7816"
INFO[0147] Using cache "/Users/amitd2/Library/Caches/lima/download/by-url-sha256/4c138835ee225c3a7a6646414ee1c5918b252ef58815e7e3d47943d297090721/data" 
INFO[0148] [hostagent] Starting QEMU (hint: to watch the boot progress, see "/Users/amitd2/.lima/ebpf/serial.log") 
INFO[0148] SSH Local Port: 53299                        
INFO[0148] [hostagent] Waiting for the essential requirement 1 of 5: "ssh" 
INFO[0169] [hostagent] The essential requirement 1 of 5 is satisfied 
INFO[0169] [hostagent] Waiting for the essential requirement 2 of 5: "user session is ready for ssh" 
INFO[0169] [hostagent] The essential requirement 2 of 5 is satisfied 
INFO[0169] [hostagent] Waiting for the essential requirement 3 of 5: "sshfs binary to be installed" 
INFO[0178] [hostagent] The essential requirement 3 of 5 is satisfied 
INFO[0178] [hostagent] Waiting for the essential requirement 4 of 5: "/etc/fuse.conf (/etc/fuse3.conf) to contain \"user_allow_other\"" 
INFO[0178] [hostagent] The essential requirement 4 of 5 is satisfied 
INFO[0178] [hostagent] Waiting for the essential requirement 5 of 5: "the guest agent to be running" 
INFO[0178] [hostagent] The essential requirement 5 of 5 is satisfied 
INFO[0179] [hostagent] Mounting "/Users/amitd2" on "/Users/amitd2" 
INFO[0179] [hostagent] Mounting "/tmp/lima" on "/tmp/lima" 
INFO[0179] [hostagent] Waiting for the optional requirement 1 of 2: "systemd must be available" 
INFO[0179] [hostagent] Forwarding "/run/lima-guestagent.sock" (guest) to "/Users/amitd2/.lima/ebpf/ga.sock" (host) 
INFO[0179] [hostagent] The optional requirement 1 of 2 is satisfied 
INFO[0179] [hostagent] Not forwarding TCP 0.0.0.0:22    
INFO[0179] [hostagent] Waiting for the optional requirement 2 of 2: "containerd binaries to be installed" 
INFO[0179] [hostagent] Not forwarding TCP 127.0.0.53:53 
INFO[0179] [hostagent] Not forwarding TCP [::]:22       
INFO[0185] [hostagent] The optional requirement 2 of 2 is satisfied 
INFO[0185] [hostagent] Waiting for the final requirement 1 of 1: "boot scripts must have finished" 
INFO[0226] [hostagent] Waiting for the final requirement 1 of 1: "boot scripts must have finished" 
INFO[0266] [hostagent] Waiting for the final requirement 1 of 1: "boot scripts must have finished" 
INFO[0266] [hostagent] The final requirement 1 of 1 is satisfied 
INFO[0266] READY. Run `limactl shell ebpf` to open the shell. 
```
- THEN: `limactl shell ebpf`