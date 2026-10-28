## 2026-10-28 - Optimizing Go Concurrent Read Handlers
**Learning:** Using `sync.Mutex` with `Lock()` in HTTP GET handlers blocks all concurrent read requests, causing latency bottleneck (~1911 ns/op). Replacing `sync.Mutex` with `sync.RWMutex` using `RLock()` and moving HTTP header configuration outside critical sections allows concurrent reads without contention, reducing read latency by ~53% (~896 ns/op).
**Action:** Always prefer `sync.RWMutex` with `RLock()` for read-heavy shared slice/map endpoints and move response header configuration outside mutex locks.
