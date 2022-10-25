## Why define traits for external types

### A DNS resolver that BLOCKs HTTP calls to localhost

#### Attempt 1
```rust
let addr = req.as_str();
let addr = (addr, 0).to_socket_addrs();

if let Ok(addresses) = addr {
    for a in addresses {
        if a.ip().eq(&Ipv4Addr::new(127, 0, 0, 1)) {
            return Box::pin(async { Err(io::Error::from(ErrorKind::Other)) });
        }
    }
}
```

```yaml
- GIVEN: Above is a middleware on top of Tower / hyper
- HOPE: Is to_socket_addrs a parser
- HOPE: Is to_socket_addrs a resolver
- HOPE: Filter only the valid addresses
- MAD: Last line of logic is just "What is this?"
- WOW: Err within an OK condition => What is OK after all?
- AUTHOR: Above is not a foolproof solution
```

#### NOTE: This is available in Rust
```rust
pub(crate) trait IsLocalhost {
    fn is_localhost(&self) -> bool;
}
```

```yaml
- QUIRK: create / crate
- QUIRK: Name of trait is Boolean
- SIGN: A function that takes a reference of itself & returns bool
- KIND: Treat the trait like an interface
```

```rust
impl IsLocalhost for Ipv4Addr {
    fn is_localhost(&self) -> bool {
        Ipv4Addr::new(127, 0, 0, 1).eq(self) || Ipv4Addr::new(0, 0, 0, 0).eq(self)
    }
}

impl IsLocalhost for Ipv6Addr {
    fn is_localhost(&self) -> bool {
        Ipv6Addr::new(0, 0, 0, 0, 0, 0, 0, 1).eq(self)
    }
}
```

```yaml
- NOTE: Above are available in rust std::net
- NOTE: std::net has an enum IpAddr enum for IPV4 & IPV6
```

#### Attempt 2 - Implement IsLocalhost for IpAddr
```rust
impl IsLocalhost for IpAddr {
    fn is_localhost(&self) -> bool {
        match self {
            IpAddr::V4(ref a) => a.is_localhost(),
            IpAddr::V6(ref a) => a.is_localhost(),
        }
    }
}
```

```yaml
- QUIRK: `a` just like that?
```

#### Attempt 2 contd. - Implement IsLocalhost for SockerAddr (of std::net)
```rust
impl IsLocalhost for SocketAddr {
    fn is_localhost(&self) -> bool {
        self.ip().is_localhost()
    }
}
```

#### IMP: Parse the entire route of IP address until we get the actual server
```yaml
- 
```

## References
```yaml
- https://fettblog.eu/rust-tiny-little-traits/
```
