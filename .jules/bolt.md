## 2026-08-31 - Use sync.RWMutex for concurrent HTTP read handlers
**Learning:** In Go HTTP handlers serving read-heavy workloads, using a standard `sync.Mutex` causes lock contention when multiple goroutines process read requests simultaneously. Switching to `sync.RWMutex` with `RLock()`/`RUnlock()` and keeping header settings outside the critical section reduced parallel read latency from ~1934 ns/op to ~969 ns/op (approx. 50% speedup).
**Action:** Always check if shared state in Go HTTP endpoints is read-dominated and use `sync.RWMutex` while minimizing the critical section length.
