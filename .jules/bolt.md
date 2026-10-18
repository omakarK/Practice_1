## 2026-08-31 - Optimize Concurrent Read Latency in Go HTTP Handlers
**Learning:** Switching `sync.Mutex` to `sync.RWMutex` with `RLock`/`RUnlock` for read-heavy handlers (`getTasksHandler`) and moving header setup outside the critical lock section reduces parallel read contention drastically (~1800 ns/op down to ~840 ns/op).
**Action:** Use `sync.RWMutex` for shared data structures that are read far more frequently than written, and keep non-thread-state operations (like setting HTTP response headers) outside the critical section.
