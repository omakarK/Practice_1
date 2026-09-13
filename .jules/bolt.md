# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read Concurrency and Lock Granularity in Go HTTP Handlers
**Learning:** Exclusive `sync.Mutex` locks on HTTP GET endpoints cause unnecessary contention under parallel read workloads. Using `sync.RWMutex` with `RLock()` allows concurrent reads. Furthermore, setting HTTP headers before entering the critical section minimizes lock hold duration.
**Action:** When handling read-dominated endpoints in Go, always use `sync.RWMutex` and keep lock scope strictly around shared state access.
