## Motivation
I will journal my journey on understanding rekor via ssh or vice-versa. Rekor is interested in ssh
as one of the may pki implementations to sign files.

### Day 1 (Sep 2022)
- A recent feature of SSH keys is the ability to sign files
- We never say an SSH key but keys i.e. plural
- SSH keys is split into public & private keys
- The public key is named as id_rsa.pub while the private is id_rsa
- These files are encoded & formatted

- Public keys are formatted as in "known hosts" format
- In addition to key material it contains the algorithm (e.g. ssh-rsa) & a comment (e.g. an email-id)

- Private keys are stored in an PEM format
- They resemble PGP or x509 keys

- Wire Format
- Bytes are laid out in order
- Fixed length fields are laid out at the proper offset with the specified length
- Strings are stored with the size as a prefix

## References
- https://github.com/sigstore/rekor/tree/main/pkg/pki/ssh
