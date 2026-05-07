## 2026-05-07 - RWMutex for concurrent read performance
**Learning:** In a Go web server with global state, sync.Mutex can become a bottleneck for read-heavy endpoints under high concurrency. Replacing it with sync.RWMutex allows parallel reads, significantly improving throughput.
**Action:** Always prefer sync.RWMutex when a shared resource is read significantly more often than it is written, and use RLock() for read handlers.
