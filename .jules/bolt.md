# Bolt's Journal - Critical Learnings

## 2026-08-31 - HTTP Handler Read-Write Locks in Go
**Learning:** Using `sync.Mutex` for HTTP read handlers causes serial contention under high parallel read traffic. Switching to `sync.RWMutex` with `RLock` allows concurrent reads without blocking, significantly reducing read latency in Go REST endpoints. Moving header setting outside the critical section further reduces lock hold duration.
**Action:** Use `sync.RWMutex` for shared in-memory data structures with heavy read workloads and keep lock scopes as narrow as possible.
