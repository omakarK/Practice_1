## 2026-08-18 - Read-Heavy Lock Optimization with sync.RWMutex
**Learning:** In Go HTTP handlers serving read-heavy workloads, using an exclusive `sync.Mutex` creates unnecessary thread contention under concurrent requests. Upgrading to `sync.RWMutex` with `RLock` and keeping non-critical work (like HTTP header initialization) outside the lock significantly reduces parallel read latency (~1971 ns/op to ~1111 ns/op).
**Action:** Use `sync.RWMutex` for shared data structures that are predominantly read and minimize work done inside critical lock sections.
