## Motivation
I will journal my journey on understanding rekor via ssh or vice-versa. Rekor is interested in ssh
as one of the may pki implementations to sign files.

### Notes (Sep 2022)
```yaml
- A RECENT feature of SSH keys is the ability to SIGN FILES
- We never say an SSH key but keys i.e. plural
- SSH keys is split into public & private keys
- The public key is named as id_rsa.pub while the private is id_rsa
- These files are ENCODED & FORMATTED
```

```yaml
- Public keys are formatted as in "known hosts" format
- In addition to key material it contains the algorithm (e.g. ssh-rsa) & a comment (e.g. an email-id)
```

```yaml
- Private keys are stored in an PEM format
- They resemble PGP or x509 keys
```

```yaml
- Wire Format
- Bytes are laid out in order
- Fixed length fields are laid out at the proper offset with the specified length
- Strings are stored with the size as a prefix
```

## References
```yaml
- https://github.com/sigstore/rekor/tree/main/pkg/pki/ssh
```
