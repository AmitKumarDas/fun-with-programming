## Ten thousand feet overviews on networking & friends

### no_proxy
- WHEN: Want to use a proxy server for everything but internal.dev.com and internal2.dev.com
- THEN: export no_proxy=internal.dev.com,internal2.dev.com
- TIL: Go tries the uppercase ENV before falling back to the lowercase naming
- QUIRK: If there is a leading . in the no_proxy setting, the behavior varies (curl vs wget)
- TRY: env https_proxy=http://non.existent/ no_proxy=.gitlab.com curl https://gitlab.com
- TRY: env https_proxy=http://non.existent/ no_proxy=.gitlab.com wget https://gitlab.com
- QUIRK: In some cases, setting no_proxy to * effectively disables proxies altogether
- TIL: Do not set IP addresses in no_proxy unless that IPs are explictly used by the client
- TIL: CIDR block e.g. 18.240.0.1/24 only work when request is directly made to an IP address
- TIL: Only Go and Ruby allow CIDR blocks
- TIL: Go automatically disables the use of a proxy if it detects a loopback IP address
- REF: https://about.gitlab.com/blog/2021/01/27/we-need-to-talk-no-proxy/
