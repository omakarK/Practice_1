# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read Lock Bottlenecks in HTTP Handlers
**Learning:** In Go HTTP handlers reading shared state, using `sync.Mutex` and setting response headers inside critical sections severely throttles parallel read throughput (~1900 ns/op). Switching to `sync.RWMutex` with `RLock` and keeping header mutation outside the lock reduces lock contention significantly.
**Action:** Use `sync.RWMutex` with `RLock` for read-only endpoint handlers and set headers outside lock boundaries.
