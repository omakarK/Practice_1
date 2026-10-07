# Bolt's Journal

## 2026-08-31 - HTTP Handler Concurrency with RWMutex
**Learning:** In Go HTTP read handlers with high concurrency, using `sync.Mutex` causes lock contention across goroutines. Switching to `sync.RWMutex` with `RLock()` and hoisting HTTP response header configuration outside the critical lock section reduces parallel read latency by ~70% (from ~3300 ns/op to ~915 ns/op).
**Action:** Always prefer `sync.RWMutex` with `RLock()` for read-heavy handlers and minimize the critical section by hoisting non-state-dependent setup logic (like setting headers) outside the lock.
