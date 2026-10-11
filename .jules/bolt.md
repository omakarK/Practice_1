# Bolt's Journal - Critical Learnings

## 2026-08-31 - Read Lock Concurrency for Go HTTP Handlers
**Learning:** Switching from `sync.Mutex` to `sync.RWMutex` with `RLock` for read-only HTTP endpoints (`getTasksHandler`), along with setting response headers outside the critical section, dramatically reduces lock contention under parallel request workloads (latency dropped from ~2067 ns/op to ~995 ns/op, >50% improvement).
**Action:** Always prefer `sync.RWMutex` for data structures with high read-to-write ratios, and minimize work done inside critical sections by setting headers or preparing un-synchronized state prior to acquiring locks.
