```yaml
https://www.youtube.com/watch?v=BH2jvJ74npM
```

```yaml
- batch
- control plane - 100s of assertion
- zero allocations - allocate in advance in startup
- zero deserialisation
  - fixed sized data structure for everything
  - 64 byte cache line
- zero copy
  - spare the L1-L3 cache
  - reduce memcpy
  - dont thrash the CPU cache
  - reduce cache misses
  - direct IO
- zero syscalls
  - NVME
  - submit IO to kernel without syscalls
- zero threads
  - use kernel threadpool
  - use io_uring
- io_uring
  - used for all disk & network IO
  - minus the cost of userspace thread pools to emulate async IO
  - 0 syscalls (amortized)
  - 0 context switches (ring buffer)
  - 2x throughput (4KiB sector)
```

#### Financial Exchange Architecture
```yaml
- mechanical sympathy
  - martin thompson (LMAX)
- replicated state machine
  - say it backwards & it makes sense
  - A machine (hardware that fails)
  - With state (like a hash map)
  - Replicated (for fault tolerance)
- RSM (in 3 pieces)
  - Append command to disk (like Redis AOF)
  - Replicate command to another disk
  - Apply command to your state
- Keep the same order everywhere s.t. each machine arrives at same state
- After a crash, replay the log
- Strict serializability i.e. high level of consistency
- Keep the log order same by daisy chaining / hash chaining the log entries
  - like ZFS
  - detect kernel or firmware is erroring
- Viewstamped Replication (consensus protocol)
  - Pioneered in 1988
  - a year before PAXOS draft
- comptime consensus
```

```yaml
- TigerBeetle follows the state of the art in financial exchange architectures
- advanced by Martin Thompson and LMAX
- a lockless thread-per-core design with mechanical sympathy
- that models the system of record as a replicated state machine for fault-tolerance

- TigerBeetle incorporates groundbreaking research on storage faults
- led by Remzi and Andrea Arpaci-Dusseau at the University of Wisconsin-Madison

- At the heart of TigerBeetle is the pioneering Viewstamped Replication consensus protocol
- developed by Brian M. Oki with Turing Award-winner Barbara Liskov and later James Cowling at MIT
- for low-latency leader election and optimal strict serializability
```
